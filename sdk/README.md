# TruthValidator Go SDK

用于与TruthValidator智能合约交互的Go SDK

## 功能

- 连接以太坊节点
- 钱包管理
- 提案提交
- 结果查询
- 投票操作

## 安装

```bash
go get github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/sdk
```

## 使用示例

### 初始化客户端

```go
package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/sdk"
)

func main() {
	// 配置
	cfg := &sdk.Config{
		RPCURL:         "https://mainnet.infura.io/v3/YOUR_PROJECT_ID",
		ContractAddress: common.HexToAddress("0x..."), // 合约地址
		ChainID:        big.NewInt(1), // 主网
	}

	// 创建客户端
	client, err := sdk.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 创建钱包
	wallet, err := sdk.NewWalletFromPrivateKey("YOUR_PRIVATE_KEY")
	if err != nil {
		panic(err)
	}

	// 提交提案
	contentHash := "Qm...IPFS哈希..."
	txHash, err := client.SubmitProposal(
		context.Background(),
		wallet,
		cfg.ChainID,
		contentHash,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("提案已提交，交易哈希: %s\n", txHash.Hex())

	// 查询提案
	proposalID := big.NewInt(1)
	proposal, err := client.GetProposal(context.Background(), proposalID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("提案内容哈希: %s\n", proposal.ContentHash)
	fmt.Printf("投票总数: %s\n", proposal.VoteCount.String())

	// 查询投票结果
	results, err := client.GetVoteResults(context.Background(), proposalID)
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("投票人: %s, 结果: %v, 原因: %s\n", 
			result.Voter.Hex(), result.IsApproved, result.Reason)
	}
}
```

## API文档

### TruthValidatorClient

- `NewClient(cfg *Config)`: 创建新客户端
- `Close()`: 关闭连接
- `SubmitProposal()`: 提交新提案
- `GetProposal()`: 获取提案信息
- `GetVoteResult()`: 获取单个投票结果
- `GetVoteResults()`: 获取所有投票结果

### Wallet

- `NewWalletFromPrivateKey()`: 从私钥创建钱包
- `GetAddress()`: 获取钱包地址
- `NewTransactor()`: 创建交易签名者
