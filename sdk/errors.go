package sdk

import "errors"

var (
	// ErrInvalidPrivateKey invalid private key
	ErrInvalidPrivateKey = errors.New("invalid private key")

	// ErrContractCall contract call failed
	ErrContractCall = errors.New("contract call failed")

	// ErrRPCConnection RPC connection failed
	ErrRPCConnection = errors.New("rpc connection failed")

	// ErrInvalidProposalID invalid proposal ID
	ErrInvalidProposalID = errors.New("invalid proposal ID")

	// ErrInvalidContentHash invalid content hash
	ErrInvalidContentHash = errors.New("invalid content hash")

	// ErrInvalidChainID invalid chain ID
	ErrInvalidChainID = errors.New("invalid chain ID")
)

// WrapError wraps error with additional message
func WrapError(err error, msg string) error {
	return errors.New(msg + ": " + err.Error())
}
