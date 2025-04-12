package sdk

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

type TruthValidatorClient struct {
	contract *TruthValidatorSentientNet.TruthValidatorSentientNet
	client   *ethclient.Client
}

// NewClient 创建新的TruthValidator客户端
func NewClient(cfg *Config) (*TruthValidatorClient, error) {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// 连接以太坊节点
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	// 实例化合约
	contract, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(cfg.ContractAddress, client)
	if err != nil {
		return nil, err
	}

	return &TruthValidatorClient{
		contract: contract,
		client:   client,
	}, nil
}

// Close 关闭客户端连接
func (c *TruthValidatorClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
