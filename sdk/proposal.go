package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ProposalInfo 提案信息
type ProposalInfo struct {
	Proposer    common.Address
	ContentHash string
	Status      uint8
	VoteCount   *big.Int
}

// SubmitProposal 提交新提案
func (c *TruthValidatorClient) SubmitProposal(
	ctx context.Context,
	wallet *Wallet,
	chainID *big.Int,
	contentHash string,
) (common.Hash, error) {
	// 创建交易签名者
	opts, err := wallet.NewTransactor(chainID)
	if err != nil {
		return common.Hash{}, err
	}
	opts.Context = ctx

	// 调用合约提交提案
	tx, err := c.contract.SubmitProposal(opts, contentHash)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// GetProposal 获取提案信息
func (c *TruthValidatorClient) GetProposal(
	ctx context.Context,
	proposalID *big.Int,
) (*ProposalInfo, error) {
	// 查询提案基础信息
	proposal, err := c.contract.Proposals(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, err
	}

	// 查询投票统计
	voteCounts, err := c.contract.GetVoteCounts(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, err
	}

	totalVotes := new(big.Int).Add(voteCounts.YesVotes, voteCounts.NoVotes)

	return &ProposalInfo{
		Proposer:    common.HexToAddress("0x0"), // 合约中未存储proposer，使用默认值
		ContentHash: proposal.Content,
		Status:      uint8(0), // 0=进行中,1=已通过,2=已拒绝
		VoteCount:   totalVotes,
	}, nil
}
