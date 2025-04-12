package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// DefaultRPCURL 默认RPC节点URL
	DefaultRPCURL = "http://localhost:8545"
)

// Config SDK配置
type Config struct {
	// RPCURL 以太坊节点RPC URL
	RPCURL string

	// ContractAddress 合约地址
	ContractAddress common.Address

	// ChainID 链ID
	ChainID *big.Int
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig(contractAddress common.Address, chainID *big.Int) *Config {
	return &Config{
		RPCURL:         DefaultRPCURL,
		ContractAddress: contractAddress,
		ChainID:        chainID,
	}
}

// Validate 验证配置有效性
func (c *Config) Validate() error {
	if c.RPCURL == "" {
		return ErrRPCConnection
	}
	if c.ContractAddress == (common.Address{}) {
		return ErrContractCall
	}
	if c.ChainID == nil {
		return ErrInvalidChainID
	}
	return nil
}
