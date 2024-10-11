package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum-optimism/optimism/op-service/ctxinterrupt"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/mdehoog/op-enclave/bindings"
	"github.com/mdehoog/op-enclave/op-withdrawer/withdrawals"
	"github.com/urfave/cli/v2"
)

var (
	L1URLFlag = &cli.StringFlag{
		Name:     "l1-url",
		Usage:    "URL of an L1 RPC host",
		EnvVars:  []string{"L1_URL"},
		Required: true,
	}
	L2URLFlag = &cli.StringFlag{
		Name:     "l2-url",
		Usage:    "URL of an L2 RPC host",
		EnvVars:  []string{"L2_URL"},
		Required: true,
	}
	PortalAddressFlag = &cli.StringFlag{
		Name:     "portal-address",
		Usage:    "Path to the config file",
		EnvVars:  []string{"PORTAL_ADDRESS"},
		Required: true,
	}
	WithdrawalTxHashFlag = &cli.StringFlag{
		Name:     "withdrawal-tx-hash",
		Usage:    "Hash of the withdrawal transaction",
		EnvVars:  []string{"WITHDRAWAL_TX_HASH"},
		Required: true,
	}
	PrivateKeyFlag = &cli.StringFlag{
		Name:     "private-key",
		Usage:    "Private key to sign the transaction",
		EnvVars:  []string{"PRIVATE_KEY"},
		Required: true,
	}
)

var Flags = []cli.Flag{
	L1URLFlag,
	L2URLFlag,
	PortalAddressFlag,
	WithdrawalTxHashFlag,
	PrivateKeyFlag,
}

func main() {
	oplog.SetupDefaults()

	app := cli.NewApp()
	app.Flags = cliapp.ProtectFlags(Flags)
	app.Name = "withdrawer"
	app.Usage = "Withdraws funds from L2 to L1"
	app.Action = Main

	ctx := ctxinterrupt.WithSignalWaiterMain(context.Background())
	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Crit("Application failed", "message", err)
	}
}

func Main(cliCtx *cli.Context) error {
	l1URL := cliCtx.String(L1URLFlag.Name)
	l2URL := cliCtx.String(L2URLFlag.Name)
	portalAddress := common.HexToAddress(cliCtx.String(PortalAddressFlag.Name))
	withdrawalTxHash := common.HexToHash(cliCtx.String(WithdrawalTxHashFlag.Name))
	privateKey, err := crypto.HexToECDSA(cliCtx.String(PrivateKeyFlag.Name))
	if err != nil {
		return err
	}

	ctx := context.Background()
	l1, err := ethclient.DialContext(ctx, l1URL)
	if err != nil {
		return err
	}
	l2, err := ethclient.DialContext(ctx, l2URL)
	if err != nil {
		return err
	}
	l2g := gethclient.New(l2.Client())

	chainID, err := l1.ChainID(ctx)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	portal, err := bindings.NewPortal(portalAddress, l1)
	if err != nil {
		return err
	}

	receipt, err := withdrawals.WaitForReceipt(ctx, l2, withdrawalTxHash, 1*time.Second)
	if err != nil {
		return err
	}

	receipt, err = ProveWithdrawal(ctx, l1, l2, l2g, opts, portal, withdrawalTxHash, receipt.BlockNumber)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrawal proved: %s\n", receipt.TxHash)

	return nil
}

func ProveWithdrawal(ctx context.Context, l1, l2 *ethclient.Client, l2g *gethclient.Client, opts *bind.TransactOpts, portal *bindings.Portal, withdrawalTxHash common.Hash, withdrawalTxBlock *big.Int) (*types.Receipt, error) {
	pollInterval := 1 * time.Second

	outputOracleAddress, err := portal.L2Oracle(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	outputOracle, err := bindings.NewOutputOracle(outputOracleAddress, l1)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Waiting for TEE proof of block %d... ", withdrawalTxBlock)
	l2OutputBlock, err := withdrawals.WaitForOutputBlock(ctx, outputOracle, withdrawalTxBlock, pollInterval)
	fmt.Println("done")

	tx, err := withdrawals.ProveAndFinalizeWithdrawal(ctx, l2g, l2, opts, outputOracle, portal, withdrawalTxHash, l2OutputBlock)
	if err != nil {
		return nil, err
	}
	receipt, err := withdrawals.WaitForReceipt(ctx, l1, tx.Hash(), pollInterval)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Message proved: %s\n", receipt.TxHash)
	return receipt, nil
}
