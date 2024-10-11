package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	bindings2 "github.com/ethereum-optimism/optimism/op-e2e/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/predeploys"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/mdehoog/op-enclave/bindings"
	"github.com/mdehoog/op-enclave/op-withdrawer/withdrawals"
)

const pollInterval = 250 * time.Millisecond

func main() {
	chainID := 1709200504
	deployedJSON := fmt.Sprintf("deployments/84532-%d-deployed.json", chainID)
	ecKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key from PRIVATE_KEY envvar: %v", err)
	}
	from := crypto.PubkeyToAddress(ecKey.PublicKey)

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
	withdraw := true
	withdrawDeposit := false
	proveWithdrawal := true
	redeposit := false

	withdrawalTxHash := common.HexToHash("0x")
	withdrawalTxBlock := big.NewInt(0)

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

	ReportBalances(ctx, l1, l2, from)

	amountInEth := 0.01
	amountInGwei := int64(amountInEth * 1000000000)
	value := big.NewInt(0).Mul(big.NewInt(amountInGwei), big.NewInt(1000000000))
	start := time.Now()
	if deposit {
		receipt := Deposit(ctx, l1, transactOpsFactory(false), l1Bridge, value)
		AwaitDeposit(ctx, l2, receipt, portalAddr)
		total := time.Since(start)
		fmt.Printf("Deposit E2E took: %s\n\n", total)
		ReportBalances(ctx, l1, l2, from)
	}

	start = time.Now()
	if withdrawDeposit {
		withdrawalTxHash, withdrawalTxBlock = WithdrawDeposit(ctx, l1, l2, transactOpsFactory, l1Bridge, l2Messenger)
	} else if withdraw {
		value, withdrawalTxHash, withdrawalTxBlock = Withdraw(ctx, l2, transactOpsFactory(true), l2Bridge)
	}

	if proveWithdrawal {
		receipt := ProveWithdrawal(ctx, l1, l2, l2g, transactOpsFactory(false), outputOracle, portal, withdrawalTxHash, withdrawalTxBlock)
		if withdrawDeposit {
			AwaitDeposit(ctx, l2, receipt, portalAddr)
		}
		total := time.Since(start)
		fmt.Printf("Withdrawal E2E took: %s\n\n", total)
		ReportBalances(ctx, l1, l2, from)
	}

	if redeposit {
		receipt := Deposit(ctx, l1, transactOpsFactory(false), l1Bridge, value)
		AwaitDeposit(ctx, l2, receipt, portalAddr)
		ReportBalances(ctx, l1, l2, from)
	}

	cancel()
}

var l1Balance, l2Balance *big.Int

func ReportBalances(ctx context.Context, l1, l2 *ethclient.Client, addr common.Address) {
	l1b, _ := l1.BalanceAt(ctx, addr, nil)
	l2b, _ := l2.BalanceAt(ctx, addr, nil)
	if l1Balance != nil {
		fmt.Printf("Balance change of %s on L2: %s, L3: %s\n\n", addr, new(big.Int).Sub(l1b, l1Balance), new(big.Int).Sub(l2b, l2Balance))
	}
	l1Balance = l1b
	l2Balance = l2b
}

func Deposit(ctx context.Context, l1 *ethclient.Client, opts *bind.TransactOpts, l1Bridge common.Address, value *big.Int) *types.Receipt {
	fmt.Printf("Depositing %s wei to the L3\n", value)
	_, receipt, err := send(ctx, l1, l1Bridge, opts, nil, value, false)
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

	receipt, err = withdrawals.WaitForReceipt(ctx, l2, types.NewTx(deposits[0]).Hash(), pollInterval)
	if err != nil {
		log.Fatalf("Error waiting for confirmation: %v", err)
	}
	fmt.Printf("Deposit confirmed: %s\n", receipt.TxHash)
}

func Withdraw(ctx context.Context, l2 *ethclient.Client, opts *bind.TransactOpts, l2Bridge common.Address) (value *big.Int, withdrawalTxHash common.Hash, withdrawalTxBlock *big.Int) {
	fmt.Printf("Withdrawing entire balance from L3\n")
	tx, receipt, err := send(ctx, l2, l2Bridge, opts, nil, nil, true)
	if err != nil {
		log.Fatalf("Error sending withdrawal: %v", err)
	}
	value = tx.Value()
	withdrawalTxHash = receipt.TxHash
	withdrawalTxBlock = receipt.BlockNumber
	fmt.Printf("Withdrew %s wei: %s (block %d)\n", tx.Value(), withdrawalTxHash, withdrawalTxBlock)
	return
}

func WithdrawDeposit(ctx context.Context, l1, l2 *ethclient.Client, optsFactory func(l2 bool) *bind.TransactOpts, l1Bridge, l2Messenger common.Address) (withdrawalTxHash common.Hash, withdrawalTxBlock *big.Int) {
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
	_, receipt, err := send(ctx, l2, l2Messenger, opts, tx.Data(), nil, true)
	if err != nil {
		log.Fatalf("Error sending withdrawal: %v", err)
	}
	withdrawalTxHash = receipt.TxHash
	withdrawalTxBlock = receipt.BlockNumber
	fmt.Printf("Deposit withdrawal sent: %s (block %d)\n", withdrawalTxHash, withdrawalTxBlock)
	return
}

func ProveWithdrawal(ctx context.Context, l1, l2 *ethclient.Client, l2g *gethclient.Client, opts *bind.TransactOpts, outputOracle *bindings.OutputOracle, portal *bindings.Portal, withdrawalTxHash common.Hash, withdrawalTxBlock *big.Int) *types.Receipt {
	fmt.Printf("Waiting for TEE proof of block %d... ", withdrawalTxBlock)
	l2OutputBlock, err := withdrawals.WaitForOutputBlock(ctx, outputOracle, withdrawalTxBlock, pollInterval)
	fmt.Println("done")

	tx, err := withdrawals.ProveAndFinalizeWithdrawal(ctx, l2g, l2, opts, outputOracle, portal, withdrawalTxHash, l2OutputBlock)
	if err != nil {
		log.Fatalf("Error proving and finalizing withdrawal: %v", err)
	}
	receipt, err := withdrawals.WaitForReceipt(ctx, l1, tx.Hash(), pollInterval)
	if err != nil {
		log.Fatalf("Error waiting for confirmation: %v", err)
	}
	fmt.Printf("Message proved: %s\n", receipt.TxHash)
	return receipt
}

func send(ctx context.Context, client *ethclient.Client, to common.Address, opts *bind.TransactOpts, data []byte, value *big.Int, sendFullBalance bool) (*types.Transaction, *types.Receipt, error) {
	if value == nil {
		value = big.NewInt(0)
	}
	opts.Value = value

	binding := bind.NewBoundContract(to, abi.ABI{}, client, client, client)

	if sendFullBalance {
		balance, err := client.BalanceAt(ctx, opts.From, nil)
		if err != nil {
			return nil, nil, err
		}

		opts.NoSend = true
		opts.Value = new(big.Int).Div(balance, big.NewInt(2))
		tx, err := binding.RawTransact(opts, data)

		totalGas := new(big.Int).Sub(tx.Cost(), opts.Value)
		txBin, err := tx.MarshalBinary()
		if err != nil {
			return nil, nil, err
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
		opts.GasLimit = tx.Gas()
		opts.NoSend = false
		opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	}

	tx, err := binding.RawTransact(opts, data)
	if err != nil {
		return nil, nil, err
	}

	receipt, err := withdrawals.WaitForReceipt(ctx, client, tx.Hash(), pollInterval)
	if err != nil {
		return nil, nil, err
	}
	return tx, receipt, nil
}

func calculateBatchInbox(chainID *big.Int) common.Address {
	a1 := common.HexToAddress(chainID.Text(10))
	a1[0] = a1[0] | 0xff
	return a1
}
