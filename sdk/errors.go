package sdk

import "errors"

var (
	// ErrInvalidPrivateKey 无效的私钥
	ErrInvalidPrivateKey = errors.New("invalid private key")

	// ErrContractCall 合约调用失败
	ErrContractCall = errors.New("contract call failed")

	// ErrRPCConnection RPC连接失败
	ErrRPCConnection = errors.New("rpc connection failed")

	// ErrInvalidProposalID 无效的提案ID
	ErrInvalidProposalID = errors.New("invalid proposal ID")

	// ErrInvalidContentHash 无效的内容哈希
	ErrInvalidContentHash = errors.New("invalid content hash")

	// ErrInvalidChainID 无效的链ID
	ErrInvalidChainID = errors.New("invalid chain ID")
)

// WrapError 包装错误信息
func WrapError(err error, msg string) error {
	return errors.New(msg + ": " + err.Error())
}
