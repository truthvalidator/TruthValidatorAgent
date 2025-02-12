package cmd

import (
	"fmt"
	"math/big"
)

// ParseUint256 converts a string to a *big.Int
func ParseUint256(value string) (*big.Int, error) {
	id, success := new(big.Int).SetString(value, 10)
	if !success {
		return nil, fmt.Errorf("failed to parse '%s' as a uint256", value)
	}
	return id, nil
}
