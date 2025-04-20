package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	searchStorage "github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the TruthValidatorSentientNet contract to Ethereum or Filecoin FEVM",
	Long: `Deploys the TruthValidatorSentientNet contract to either:
- Ethereum network (default)
- Filecoin FEVM (when --fevm flag is used)`,
	Args: cobra.NoArgs,
	Run:  Deploy,
}

var fevmFlag bool

func init() {
	deployCmd.Flags().BoolVar(&fevmFlag, "fevm", false, "Deploy to Filecoin FEVM network")
	RootCmd.AddCommand(deployCmd)
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

func Deploy(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	var networkAddr string
	if fevmFlag {
		networkAddr = "https://api.calibration.node.glif.io/rpc/v1" // FEVM Calibration testnet RPC
		fmt.Fprintf(os.Stderr, "Deploying to Filecoin FEVM network...\n")
	} else {
		networkAddr = GetNetworkAddress()
	}

	// Connect to the network
	conn, err := connection.NewConnection(ctx, networkAddr)
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	// Deploy TruthValidatorSentientNet contract.
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare next tx", err)
	}

	if fevmFlag {
		fmt.Fprintf(os.Stderr, "Deploying TruthValidatorSentientNet contract to Filecoin FEVM...\n")
	} else {
		fmt.Fprintf(os.Stderr, "Deploying TruthValidatorSentientNet contract...\n")
	}

	searchAddr, deployTx, _, err := searchStorage.DeployTruthValidatorSentientNet(auth, conn.Client)
	if err != nil {
		ExitWithError("Failed to create deploy tx", err)
	}

	_, err = bind.WaitDeployed(ctx, conn.Client, deployTx)
	if err != nil {
		ExitWithError("Failed to deploy contract", err)
	}

	// Output deployed contract's address to stdout.
	fmt.Printf("%s\n", searchAddr.Hex())
}
