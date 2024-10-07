package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	bindings2 "github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/mdehoog/op-enclave/bindings"
)

func main() {
	chainID := 660380098
	deployedJSON := fmt.Sprintf("deployments/84532-%d-deployed.json", chainID)
	ecKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key from PRIVATE_KEY envvar: %v", err)
	}

	var deployed bindings.DeployChainDeploy
	deployedBytes, err := os.ReadFile(deployedJSON)
	if err != nil {
		log.Fatalf("Failed to read deployed JSON: %v", err)
	}
	err = json.Unmarshal(deployedBytes, &deployed)
	if err != nil {
		log.Fatalf("Failed to parse deployed JSON: %v", err)
	}

	l1Bridge := deployed.Addresses.L1StandardBridge
	portalAddr := deployed.Addresses.OptimismPortal
	l2OO := deployed.Addresses.L2OutputOracle
	l2Bridge := predeploys.L2StandardBridgeAddr
	l2Messenger := predeploys.L2CrossDomainMessengerAddr

	deposit := true
	withdraw := false
	withdrawDeposit := true
	proveWithdrawal := true

	withdrawalTxHash := common.HexToHash("0x")
	withdrawalTxBlock := uint64(0)

	ctx, cancel := context.WithCancel(context.Background())
	l1, err := ethclient.DialContext(ctx, "https://sepolia.base.org")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	l2, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	l2g := gethclient.New(l2.Client())

	outputOracle, err := bindings.NewOutputOracle(l2OO, l1)
	if err != nil {
		log.Fatalf("Failed to bind to output oracle: %v", err)
	}
	portal, err := bindings.NewPortal(portalAddr, l1)
	if err != nil {
		log.Fatalf("Failed to bind to portal: %v", err)
	}

	l1ChainID, err := l1.ChainID(ctx)
	if err != nil {
		log.Fatalf("Error getting L1 chain ID: %v", err)
	}
	l2ChainID, err := l2.ChainID(ctx)
	if err != nil {
		log.Fatalf("Error getting L2 chain ID: %v", err)
	}
	transactOpsFactory := func(l2 bool) *bind.TransactOpts {
		var signer types.Signer
		if l2 {
			signer = types.LatestSignerForChainID(l2ChainID)
		} else {
			signer = types.LatestSignerForChainID(l1ChainID)
		}
		return &bind.TransactOpts{
			From: crypto.PubkeyToAddress(ecKey.PublicKey),
			Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(tx, signer, ecKey)
			},
		}
	}

	if deposit {
		receipt := Deposit(ctx, l1, transactOpsFactory(false), l1Bridge)
		AwaitDeposit(ctx, l2, receipt, portalAddr)
	}

	balance, _ := l2.BalanceAt(ctx, crypto.PubkeyToAddress(ecKey.PublicKey), nil)
	fmt.Printf("L2 balance: %s\n", balance)
	if withdrawDeposit {
		withdrawalTxHash, withdrawalTxBlock = WithdrawDeposit(ctx, l1, l2, transactOpsFactory, l1Bridge, l2Messenger)
	} else if withdraw {
		withdrawalTxHash, withdrawalTxBlock = Withdraw(ctx, l2, transactOpsFactory(true), l2Bridge)
	}

	if proveWithdrawal {
		receipt := ProveWithdrawal(ctx, l1, l2, l2g, transactOpsFactory(false), outputOracle, portal, withdrawalTxHash, withdrawalTxBlock)
		AwaitDeposit(ctx, l2, receipt, portalAddr)
	}
	fmt.Printf("L2 balance: %s\n", balance)

	cancel()
}

func Deposit(ctx context.Context, l1 *ethclient.Client, opts *bind.TransactOpts, l1Bridge common.Address) *types.Receipt {
	amountInEth := 0.01
	receipt, err := send(ctx, l1, l1Bridge, opts, nil, amountInEth, false)
	if err != nil {
		log.Fatalf("Error sending deposit: %v", err)
	}
	fmt.Printf("Deposit sent: %s\n", receipt.TxHash)
	return receipt
}

func AwaitDeposit(ctx context.Context, l2 *ethclient.Client, receipt *types.Receipt, portalAddr common.Address) {
	deposits, err := derive.UserDeposits([]*types.Receipt{receipt}, portalAddr)
	if err != nil {
		log.Fatalf("Error deriving deposits: %v", err)
	}
	if len(deposits) != 1 {
		log.Fatalf("Expected 1 deposit, got %d", len(deposits))
	}

	receipt, err = waitForConfirmation(ctx, l2, types.NewTx(deposits[0]).Hash())
	if err != nil {
		log.Fatalf("Error waiting for confirmation: %v", err)
	}
	fmt.Printf("Deposit confirmed: %s\n", receipt.TxHash)
}

func Withdraw(ctx context.Context, l2 *ethclient.Client, opts *bind.TransactOpts, l2Bridge common.Address) (withdrawalTxHash common.Hash, withdrawalTxBlock uint64) {
	receipt, err := send(ctx, l2, l2Bridge, opts, nil, 0, true)
	if err != nil {
		log.Fatalf("Error sending withdrawal: %v", err)
	}
	withdrawalTxHash = receipt.TxHash
	withdrawalTxBlock = receipt.BlockNumber.Uint64()
	fmt.Printf("Withdrawal sent: %s (block %d)\n", withdrawalTxHash, withdrawalTxBlock)
	return
}

func WithdrawDeposit(ctx context.Context, l1, l2 *ethclient.Client, optsFactory func(l2 bool) *bind.TransactOpts, l1Bridge, l2Messenger common.Address) (withdrawalTxHash common.Hash, withdrawalTxBlock uint64) {
	bridge, err := bindings2.NewL1StandardBridge(l1Bridge, l1)
	if err != nil {
		log.Fatalf("Error binding to L1 bridge: %v", err)
	}
	opts := optsFactory(false)
	opts.NoSend = true
	tx, err := bridge.DepositETHTo(opts, opts.From, 200_000, []byte{})
	if err != nil {
		log.Fatalf("Error creating deposit tx: %v", err)
	}

	messenger, err := bindings2.NewL2CrossDomainMessenger(l2Messenger, l2)
	if err != nil {
		log.Fatalf("Error binding to messenger: %v", err)
	}
	opts = optsFactory(true)
	opts.NoSend = true
	tx, err = messenger.SendMessage(opts, l1Bridge, tx.Data(), 400_000)
	if err != nil {
		log.Fatalf("Error creating message tx: %v", err)
	}

	opts = optsFactory(true)
	receipt, err := send(ctx, l2, l2Messenger, opts, tx.Data(), 0, true)
	if err != nil {
		log.Fatalf("Error sending withdrawal: %v", err)
	}
	withdrawalTxHash = receipt.TxHash
	withdrawalTxBlock = receipt.BlockNumber.Uint64()
	fmt.Printf("Deposit withdrawal sent: %s (block %d)\n", withdrawalTxHash, withdrawalTxBlock)
	return
}

func ProveWithdrawal(ctx context.Context, l1, l2 *ethclient.Client, l2g *gethclient.Client, opts *bind.TransactOpts, outputOracle *bindings.OutputOracle, portal *bindings.Portal, withdrawalTxHash common.Hash, withdrawalTxBlock uint64) *types.Receipt {
	fmt.Printf("Waiting for TEE proof of block %d...\n", withdrawalTxBlock)
	var l2OutputBlock *big.Int
	for {
		var err error
		l2OutputBlock, err = outputOracle.LatestBlockNumber(&bind.CallOpts{})
		if err != nil {
			log.Fatalf("Error getting latest L2 output block: %v", err)
		}
		if l2OutputBlock.Uint64() >= withdrawalTxBlock {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	header, err := l2.HeaderByNumber(ctx, l2OutputBlock)
	if err != nil {
		log.Fatalf("Error getting L2 header: %v", err)
	}
	l2OutputIndex, err := outputOracle.GetL2OutputIndexAfter(&bind.CallOpts{}, header.Number)
	if err != nil {
		log.Fatalf("Error getting L2 output index: %v", err)
	}
	l2BlockNumber := header.Number

	withdrawal, err := withdrawals.ProveWithdrawalParametersForBlock(ctx, l2g, l2, l2, withdrawalTxHash, l2BlockNumber, l2OutputIndex)
	if err != nil {
		log.Fatalf("Error proving withdrawal parameters: %v", err)
	}

	outputRootProof := bindings.TypesOutputRootProof{
		Version:                  withdrawal.OutputRootProof.Version,
		StateRoot:                withdrawal.OutputRootProof.StateRoot,
		MessagePasserStorageRoot: withdrawal.OutputRootProof.MessagePasserStorageRoot,
		LatestBlockhash:          withdrawal.OutputRootProof.LatestBlockhash,
	}

	tx, err := portal.ProveAndFinalizeWithdrawalTransaction(
		opts,
		bindings.TypesWithdrawalTransaction{
			Nonce:    withdrawal.Nonce,
			Sender:   withdrawal.Sender,
			Target:   withdrawal.Target,
			Value:    withdrawal.Value,
			GasLimit: withdrawal.GasLimit,
			Data:     withdrawal.Data,
		},
		withdrawal.L2OutputIndex,
		outputRootProof,
		withdrawal.WithdrawalProof,
	)
	if err != nil {
		log.Fatalf("Error proving and finalizing withdrawal: %v", err)
	}
	receipt, err := waitForConfirmation(ctx, l1, tx.Hash())
	if err != nil {
		log.Fatalf("Error waiting for confirmation: %v", err)
	}
	fmt.Printf("Message proved: %s\n", receipt.TxHash)
	return receipt
}

func send(ctx context.Context, client *ethclient.Client, to common.Address, opts *bind.TransactOpts, data []byte, amountInEth float64, sendFullBalance bool) (*types.Receipt, error) {
	amountInGwei := int64(amountInEth * 1000000000)
	opts.Value = big.NewInt(0).Mul(big.NewInt(amountInGwei), big.NewInt(1000000000))

	fmt.Printf("From address: %s\n", opts.From)

	binding := bind.NewBoundContract(to, abi.ABI{}, client, client, client)

	if sendFullBalance {
		balance, err := client.BalanceAt(ctx, opts.From, nil)
		if err != nil {
			return nil, err
		}

		opts.NoSend = true
		opts.Value = big.NewInt(1)
		tx, err := binding.RawTransact(opts, data)

		totalGas := new(big.Int).Sub(tx.Cost(), opts.Value)
		txBin, err := tx.MarshalBinary()
		if err != nil {
			return nil, err
		}
		gasPriceOracle, _ := bindings2.NewGasPriceOracleCaller(predeploys.GasPriceOracleAddr, client)
		if gasPriceOracle != nil {
			l1Fee, _ := gasPriceOracle.GetL1Fee(&bind.CallOpts{}, txBin)
			if l1Fee != nil {
				totalGas.Add(totalGas, l1Fee)
			}
		}

		opts.Value.Sub(balance, totalGas)
		opts.GasFeeCap = tx.GasFeeCap()
		opts.GasTipCap = tx.GasTipCap()
		opts.NoSend = false
		opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	}

	tx, err := binding.RawTransact(opts, data)

	receipt, err := waitForConfirmation(ctx, client, tx.Hash())
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func waitForConfirmation(ctx context.Context, client *ethclient.Client, tx common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, tx)
		if errors.Is(err, ethereum.NotFound) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(1 * time.Second):
			}
		} else if err != nil {
			return nil, err
		} else if receipt.Status != types.ReceiptStatusSuccessful {
			return nil, errors.New("unsuccessful receipt status")
		} else {
			return receipt, nil
		}
	}
}

func calculateBatchInbox(chainID *big.Int) common.Address {
	a1 := common.HexToAddress(chainID.Text(10))
	a1[0] = a1[0] | 0xff
	return a1
}
