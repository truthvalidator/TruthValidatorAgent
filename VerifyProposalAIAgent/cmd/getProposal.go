package cmd

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var getProposalCmd = &cobra.Command{
	Use:   "getProposal [contract_address] [proposalId]",
	Short: "Get details of a specific proposal and its votes",
	Args:  cobra.ExactArgs(2),
	Run:   GetProposal,
}

func init() {
	RootCmd.AddCommand(getProposalCmd)
}

// FetchAndPrintVotes retrieves and prints the vote details for a given proposal
func FetchAndPrintVotes(sas *TruthValidatorSentientNet.TruthValidatorSentientNet, from common.Address, proposalId *big.Int) error {
	// Fetch the list of voters
	voters, err := sas.GetVoters(&bind.CallOpts{From: from}, proposalId)
	if err != nil {
		return err
	}

	fmt.Println("Total voters:", voters)

	// Fetch and print individual votes
	for _, voterAddr := range voters {
		vote, err := sas.GetVote(&bind.CallOpts{From: from}, proposalId, voterAddr)
		if err != nil {
			return err
		}

		fmt.Printf("Voter: %s, Approved: %t, Reason: %s\n", vote.Voter.Hex(), vote.IsApproved, vote.Reason)
	}

	return nil
}

func GetProposal(cmd *cobra.Command, args []string) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	proposalId, err := ParseUint256(args[1])
	if err != nil {
		ExitWithError("Invalid proposal ID", err)
	}

	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
	if err != nil {
		ExitWithError("Unable to get instance of contract", err)
	}

	// Fetch the proposal details
	proposal, err := sas.Proposals(&bind.CallOpts{From: conn.Address}, proposalId)
	if err != nil {
		ExitWithError("Failed to retrieve proposal", err)
	}

	fmt.Printf("Proposal ID: %s\n", proposal.Id.String())
	fmt.Printf("Content: %s\n", proposal.Content)
	fmt.Printf("Yes Votes: %s\n", proposal.YesVotes.String())
	fmt.Printf("No Votes: %s\n", proposal.NoVotes.String())
	fmt.Printf("Is Finalized: %t\n", proposal.IsFinalized)

	// Fetch and print vote details
	fmt.Println("\nVote Details:")
	err = FetchAndPrintVotes(sas, conn.Address, proposalId)
	if err != nil {
		ExitWithError("Failed to retrieve vote details", err)
	}
}
