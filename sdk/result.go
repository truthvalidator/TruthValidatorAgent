package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// VoteResult 投票结果
type VoteResult struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}

// GetVoteResult 获取单个投票结果
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

// GetVoteResults 获取所有投票结果
func (c *TruthValidatorClient) GetVoteResults(
	ctx context.Context,
	proposalID *big.Int,
) ([]VoteResult, error) {
	// 获取投票人列表
	voters, err := c.contract.GetVoters(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, WrapError(err, "failed to get voters")
	}

	// 获取每个投票人的投票结果
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
