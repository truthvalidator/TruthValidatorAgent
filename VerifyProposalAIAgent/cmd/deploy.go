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
	Short: "Deploy the TruthValidatorSentientNet contract",
	Args:  cobra.NoArgs,
	Run:   Deploy,
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

func Deploy(cmd *cobra.Command, args []string) {
	// Set up a context for calls with a timeout of 1 minute.
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	// Connect to the network.
	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	// Deploy TruthValidatorSentientNet contract.
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare next tx", err)
	}

	fmt.Fprintf(os.Stderr, "Deploying TruthValidatorSentientNet contract...\n")

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
