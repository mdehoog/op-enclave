package enclave

import (
	"archive/zip"
	"bytes"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/hf/nitrite"
	"github.com/hf/nsm"
	"github.com/hf/nsm/request"
)

const (
	// DefaultCARoots contains the PEM encoded roots for verifying Nitro
	// Enclave attestation signatures. You can download them from
	// https://docs.aws.amazon.com/enclaves/latest/user/verify-root.html
	DefaultCARoots       = "UEsDBBQAAAAIALkYV1GVtvolRwIAAAkDAAAIABwAcm9vdC5wZW1VVAkAA10ekl9dHpJfdXgLAAEESHEtDwQUAAAAZZJLk6JQDIX3/IrZW10Igo2LWdwXiBoE5HXZCSq0iNgKfYVfP9guJ8tTqS85Ofn4GAszy3b+EOYHtmkTFLCX+CGBbRMWEILSfYGEjVFh+8itnoe4yKq1XC7DDNptcJ2YXJCC2+smtYfzlCEBYhewjQSospASMlwCiSJ40gE5uHAijBrAldny5PaTnRkAan77iBDUiw4B+A9heZxKkedRilflYQZdVl+meW20aayfM8tU0wTEsswdCKonUFuDAPotRUo8ag59axIE3ls84xV4D0FG6gi1mFhF4cBcQNP35GIcGCvlsV504ImXnVffRqLjxpECT2tA6Xt1AFabs7zXu33i91mvXLLaefAkveQDVgEjC/ff1g60BSqYJeFdhzFCX0i1EXYFibZdTWA57Jf0q26/vZ+Ka3BbDVlz2chy2qv8wnYK9vVgVz1OWSZpBjFi3PTtp6li8Xlk7X7vTprSUrNr+FgspofpKlGNIHe9hDA3nWGE7WPgcsEaEqdMKo2LzhtPBHkoL9YOgTEgKkZ//jRA3lLGKBRIMCwP6PCyuPQ0ZhZeWJFYoYfKlPzJMRZ6Ns9vM7feX087nQta/ALcN8CjqLCsV4yEvL2Pd6JIrRBYnEjgkfOpn/hNXi+S7qjxq4hrZxUhTTuhqavH6vbGG7HYchL5e3b82RjdVkn4vdOfLbixdD8BGSFfhv6IcbYS63Vy2M3xrfXMLs2Cz1kjF7hUvsPnRb46d0UNtwY/iftcuJtsMnckW2yGmcz/Sr+fzRz637f/A1BLAQIeAxQAAAAIALkYV1GVtvolRwIAAAkDAAAIABgAAAAAAAEAAACkgQAAAAByb290LnBlbVVUBQADXR6SX3V4CwABBEhxLQ8EFAAAAFBLBQYAAAAAAQABAE4AAACJAgAAAAA="
	DefaultCARootsSHA256 = "8cf60e2b2efca96c6a9e71e851d00c1b6991cc09eadbe64a6a1d1b1eb9faff7c"
)

var (
	defaultRoot                = createAWSNitroRoot()
	l2ToL1MessagePasserAddress = common.HexToAddress("0x4200000000000000000000000000000000000016")
)
var uint256Type abi.Type
var uint64Type abi.Type
var boolType abi.Type
var addressType abi.Type
var bytes32Type abi.Type

func init() {
	uint256Type, _ = abi.NewType("uint256", "", nil)
	uint64Type, _ = abi.NewType("uint64", "", nil)
	boolType, _ = abi.NewType("bool", "", nil)
	addressType, _ = abi.NewType("address", "", nil)
	bytes32Type, _ = abi.NewType("bytes32", "", nil)
}

func createAWSNitroRoot() *x509.CertPool {
	roots, err := base64.StdEncoding.DecodeString(DefaultCARoots)
	if err != nil {
		panic("error decoding AWS root cert")
	}
	sha := sha256.Sum256(roots)
	expected := common.HexToHash(DefaultCARootsSHA256)
	if !bytes.Equal(sha[:], expected[:]) {
		panic("DefaultCARoots checksum failed")
	}
	reader, err := zip.NewReader(bytes.NewReader(roots), int64(len(roots)))
	ca, err := reader.File[0].Open()
	if err != nil {
		panic("error reading AWS root cert zip")
	}
	pem, err := io.ReadAll(ca)
	if err != nil {
		panic("error reading AWS root cert")
	}
	pool := x509.NewCertPool()
	ok := pool.AppendCertsFromPEM(pem)
	if !ok {
		panic("error parsing AWS root cert")
	}
	return pool
}

type Server struct {
	pcr0          []byte
	signerKey     *ecdsa.PrivateKey
	decryptionKey *rsa.PrivateKey
}

func NewServer() (*Server, error) {
	session, err := nsm.OpenDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open session: %w", err)
	}
	defer func() {
		_ = session.Close()
	}()
	decryptionKey, err := rsa.GenerateKey(session, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate decryption key: %w", err)
	}
	signerKey, err := ecdsa.GenerateKey(crypto.S256(), session)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signer key: %w", err)
	}
	pcr, err := session.Send(&request.DescribePCR{
		Index: 0,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to describe PCR: %w", err)
	}
	if pcr.Error != "" {
		return nil, fmt.Errorf("NSM device returned an error: %s", pcr.Error)
	}
	if pcr.DescribePCR == nil || pcr.DescribePCR.Data == nil || len(pcr.DescribePCR.Data) == 0 {
		return nil, errors.New("NSM device did not return PCR data")
	}
	return &Server{
		pcr0:          pcr.DescribePCR.Data,
		signerKey:     signerKey,
		decryptionKey: decryptionKey,
	}, nil
}

func (s *Server) SignerPublicKey() (hexutil.Bytes, error) {
	return crypto.FromECDSAPub(&s.signerKey.PublicKey), nil
}

func (s *Server) SignerAttestation() (hexutil.Bytes, error) {
	return s.publicKeyAttestation(s.SignerPublicKey)
}

func (s *Server) DecryptionPublicKey() (hexutil.Bytes, error) {
	return x509.MarshalPKIXPublicKey(s.decryptionKey.Public())
}

func (s *Server) DecryptionAttestation() (hexutil.Bytes, error) {
	return s.publicKeyAttestation(s.DecryptionPublicKey)
}

func (s *Server) publicKeyAttestation(publicKey func() (hexutil.Bytes, error)) (hexutil.Bytes, error) {
	session, err := nsm.OpenDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open session: %w", err)
	}
	defer func() {
		_ = session.Close()
	}()
	public, err := publicKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}
	res, err := session.Send(&request.Attestation{
		PublicKey: public,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get attestation: %w", err)
	}
	if res.Error != "" {
		return nil, fmt.Errorf("NSM device returned an error: %s", res.Error)
	}
	if res.Attestation == nil || res.Attestation.Document == nil {
		return nil, errors.New("NSM device did not return an attestation")
	}
	return res.Attestation.Document, nil
}

func (s *Server) EncryptedSignerKey(attestation hexutil.Bytes) (hexutil.Bytes, error) {
	verification, err := nitrite.Verify(
		attestation,
		nitrite.VerifyOptions{
			Roots:       defaultRoot,
			CurrentTime: time.Now(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to verify attestation: %w", err)
	}
	if !bytes.Equal(verification.Document.PCRs[0], s.pcr0) {
		return nil, errors.New("attestation does not match PCR0")
	}
	publicKey, err := x509.ParsePKIXPublicKey(verification.Document.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}
	public, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not RSA")
	}
	session, err := nsm.OpenDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("failed to open session: %w", err)
	}
	defer func() {
		_ = session.Close()
	}()
	ciphertext, err := rsa.EncryptPKCS1v15(session, public, crypto.FromECDSA(s.signerKey))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt key: %w", err)
	}
	return ciphertext, nil
}

func (s *Server) SetSignerKey(encrypted hexutil.Bytes) error {
	session, err := nsm.OpenDefaultSession()
	if err != nil {
		return fmt.Errorf("failed to open session: %w", err)
	}
	defer func() {
		_ = session.Close()
	}()
	decrypted, err := rsa.DecryptPKCS1v15(session, s.decryptionKey, encrypted)
	if err != nil {
		return fmt.Errorf("failed to decrypt key: %w", err)
	}
	key, err := crypto.ToECDSA(decrypted)
	if err != nil {
		return fmt.Errorf("failed to convert key: %w", err)
	}
	s.signerKey = key
	return nil
}

type Proposal struct {
	OutputRoot common.Hash
	Signature  hexutil.Bytes
}

func (s *Server) ExecuteStateless(
	chainConfig *params.ChainConfig,
	block *Block,
	witness hexutil.Bytes,
	messageAccount *AccountResult,
	prevMessageAccountHash common.Hash,
	beaconHash common.Hash,
	previousDepositHash common.Hash,
	firstDepositIndex uint64,
) (*Proposal, error) {
	// TODO prove all DepositTxs in block (both the L1Info tx as well as Portal deposits)
	// L1Info tx calldata:
	//    ///   1. _baseFeeScalar      L1 base fee scalar
	//    ///   2. _blobBaseFeeScalar  L1 blob base fee scalar
	//    ///   3. _sequenceNumber     Number of L2 blocks since epoch start.
	//    ///   4. _timestamp          L1 timestamp.
	//    ///   5. _number             L1 blocknumber.
	//    ///   6. _basefee            L1 base fee.
	//    ///   7. _blobBaseFee        L1 blob base fee.
	//    ///   8. _hash               L1 blockhash.
	//    ///   9. _batcherHash        Versioned hash to authenticate batcher by.

	signer := types.LatestSignerForChainID(chainConfig.ChainID)
	depositCount := 0
	depositHash := common.Hash{}
	for i, tx := range block.Transactions() {
		if tx.IsDepositTx() {
			if i == 0 {
				// TODO verify L1Info tx
			} else {
				depositCount++
				args := abi.Arguments{
					{Name: "lastDeposit", Type: bytes32Type},
					{Name: "from", Type: addressType},
					{Name: "to", Type: addressType},
					{Name: "mint", Type: uint256Type},
					{Name: "value", Type: uint256Type},
					{Name: "gasLimit", Type: uint64Type},
					{Name: "isCreation", Type: boolType},
				}
				from, err := signer.Sender(tx)
				if err != nil {
					return nil, fmt.Errorf("failed to get deposit tx sender: %w", err)
				}
				isCreation := tx.To() == nil
				to := tx.To()
				if isCreation {
					to = &common.Address{}
				}
				data, err := args.Pack(previousDepositHash, from, *to, tx.Mint(), tx.Value(), tx.Gas(), isCreation)
				if err != nil {
					return nil, fmt.Errorf("failed to pack deposit tx data: %w", err)
				}
				data = append(data, tx.Data()...)
				depositHash = crypto.Keccak256Hash(data)
				previousDepositHash = depositHash
				// TODO verify against beaconRoot that depositHash exists at firstDepositIndex
				firstDepositIndex++
			}
		}
	}

	w := &stateless.Witness{}
	err := rlp.DecodeBytes(witness, w)
	if err != nil {
		return nil, fmt.Errorf("failed to decode witness: %w", err)
	}

	stateRoot, _, err := core.ExecuteStateless(chainConfig, block.Block, w)

	if err = messageAccount.Verify(l2ToL1MessagePasserAddress, stateRoot); err != nil {
		return nil, fmt.Errorf("failed to verify message account: %w", err)
	}

	prevOutputRoot := outputRootV0(w.Headers[0], prevMessageAccountHash)
	outputRoot := outputRootV0(block.Header(), messageAccount.StorageHash)

	chainConfigJson, err := json.Marshal(chainConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chain config: %w", err)
	}
	chainConfigJsonHash := crypto.Keccak256(chainConfigJson)

	data := append(chainConfigJsonHash, prevOutputRoot[:]...)
	data = append(data, outputRoot[:]...)
	sig, err := crypto.Sign(crypto.Keccak256(data), s.signerKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}

	if err != nil {
		return nil, err
	}
	return &Proposal{
		OutputRoot: outputRoot,
		Signature:  sig,
	}, nil
}

func (s *Server) Aggregate(chainConfig *params.ChainConfig, prevOutputRoot common.Hash, proposals []*Proposal) (*Proposal, error) {
	chainConfigJson, err := json.Marshal(chainConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chain config: %w", err)
	}
	chainConfigJsonHash := crypto.Keccak256(chainConfigJson)

	outputRoot := prevOutputRoot
	for _, p := range proposals {
		data := append(chainConfigJsonHash, outputRoot[:]...)
		data = append(data, p.OutputRoot[:]...)
		if !crypto.VerifySignature(crypto.FromECDSAPub(&s.signerKey.PublicKey), crypto.Keccak256(data), p.Signature[:64]) {
			return nil, errors.New("invalid signature")
		}
		outputRoot = p.OutputRoot
	}

	data := append(chainConfigJsonHash, prevOutputRoot[:]...)
	data = append(data, outputRoot[:]...)
	sig, err := crypto.Sign(crypto.Keccak256(data), s.signerKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign: %w", err)
	}

	return &Proposal{
		OutputRoot: outputRoot,
		Signature:  sig,
	}, nil
}

func outputRootV0(header *types.Header, storageRoot common.Hash) common.Hash {
	hash := header.Hash()
	var buf [128]byte
	copy(buf[32:], header.Root[:])
	copy(buf[64:], storageRoot[:])
	copy(buf[96:], hash[:])
	return crypto.Keccak256Hash(buf[:])
}
