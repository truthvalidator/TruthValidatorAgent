package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// VoteResult represents a vote result
type VoteResult struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}

// GetVoteResult gets a single vote result
func (c *TruthValidatorClient) GetVoteResult(
	ctx context.Context,
	proposalID *big.Int,
	voter common.Address,
) (*VoteResult, error) {
	vote, err := c.contract.GetVote(&bind.CallOpts{
		Context: ctx,
	}, proposalID, voter)
	if err != nil {
		return nil, WrapError(err, "failed to get vote result")
	}

	return &VoteResult{
		Voter:      vote.Voter,
		IsApproved: vote.IsApproved,
		Reason:     vote.Reason,
	}, nil
}

// GetVoteResults gets all vote results
func (c *TruthValidatorClient) GetVoteResults(
	ctx context.Context,
	proposalID *big.Int,
) ([]VoteResult, error) {
	// Get voter list
	voters, err := c.contract.GetVoters(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, WrapError(err, "failed to get voters")
	}

	// Get each voter's vote result
	results := make([]VoteResult, 0, len(voters))
	for _, voter := range voters {
		vote, err := c.contract.GetVote(&bind.CallOpts{
			Context: ctx,
		}, proposalID, voter)
		if err != nil {
			return nil, WrapError(err, "failed to get vote")
		}

		results = append(results, VoteResult{
			Voter:      vote.Voter,
			IsApproved: vote.IsApproved,
			Reason:     vote.Reason,
		})
	}

	return results, nil
}
