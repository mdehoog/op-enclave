package server

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/hf/nsm"
	"github.com/hf/nsm/request"
)

const (
	tpmDevice = "/dev/tpm0" // TPM device on Nitro Enclaves
)

type Server struct {
	keys map[string]*ecdsa.PrivateKey
}

func NewServer() *Server {
	return &Server{
		keys: make(map[string]*ecdsa.PrivateKey),
	}
}

type KeyInfo struct {
	Attestation hexutil.Bytes
}

func (s *Server) NewPrivateKey() (*KeyInfo, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate key: %w", err)
	}

	tpm, err := tpm2.OpenTPM(tpmDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to open TPM: %w", err)
	}
	defer func() {
		_ = tpm.Close()
	}()

	pcrSelection := tpm2.PCRSelection{Hash: tpm2.AlgSHA256, PCRs: []int{0}}
	sealedBlob, _, _, _, _, err := tpm2.CreateKeyWithSensitive(
		tpm,
		tpm2.HandleOwner,
		pcrSelection,
		"",
		"",
		tpm2.Public{
			Type:    tpm2.AlgKeyedHash,
			NameAlg: tpm2.AlgSHA256,
		},
		crypto.FromECDSA(key),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to seal key: %w", err)
	}

	session, err := nsm.OpenDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open session: %w", err)
	}
	defer func() {
		_ = session.Close()
	}()

	res, err := session.Send(&request.Attestation{
		UserData:  sealedBlob,
		PublicKey: crypto.FromECDSAPub(&key.PublicKey),
	})
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, fmt.Errorf("NSM device returned an error: %s", res.Error)
	}
	if res.Attestation == nil || res.Attestation.Document == nil {
		return nil, errors.New("NSM device did not return an attestation")
	}

	return &KeyInfo{
		Attestation: res.Attestation.Document,
	}, nil
}

type StatelessResponse struct {
	StateRoot   common.Hash
	ReceiptRoot common.Hash
	Signature   hexutil.Bytes
}

func (s *Server) ExecuteStateless(chainConfig *params.ChainConfig, block *Block, witness, sealedKey hexutil.Bytes) (*StatelessResponse, error) {
	key, ok := s.keys[sealedKey.String()]
	if !ok {
		tpm, err := tpm2.OpenTPM(tpmDevice)
		if err != nil {
			return nil, fmt.Errorf("failed to open TPM: %w", err)
		}
		defer func() {
			_ = tpm.Close()
		}()

		loadedHandle, _, err := tpm2.Load(tpm, tpm2.HandleOwner, "", sealedKey, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to load sealed key: %w", err)
		}
		defer func() {
			_ = tpm2.FlushContext(tpm, loadedHandle)
		}()

		unsealedData, err := tpm2.Unseal(tpm, loadedHandle, "")
		if err != nil {
			return nil, fmt.Errorf("failed to unseal key: %w", err)
		}

		key, err = crypto.ToECDSA(unsealedData)
		if err != nil {
			return nil, fmt.Errorf("failed to convert key: %w", err)
		}
	}

	w := &stateless.Witness{}
	err := rlp.DecodeBytes(witness, w)
	if err != nil {
		return nil, err
	}

	stateRoot, receiptRoot, err := core.ExecuteStateless(chainConfig, block.Block, w)

	data, err := json.Marshal(chainConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chain config: %w", err)
	}
	data = crypto.Keccak256(data)
	data = append(data, w.Root().Bytes()...)
	data = append(data, stateRoot.Bytes()...)
	sig, err := crypto.Sign(crypto.Keccak256(data), key)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}

	if err != nil {
		return nil, err
	}
	return &StatelessResponse{
		StateRoot:   stateRoot,
		ReceiptRoot: receiptRoot,
		Signature:   sig,
	}, nil
}
