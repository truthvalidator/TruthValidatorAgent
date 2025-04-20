#!/bin/bash
set -e

echo "Running end-to-end tests for Filecoin FEVM Calibration testnet..."

# Deploy contract to FEVM testnet
echo "Deploying contract to FEVM Calibration testnet..."
export PRIVATE_KEY="your_testnet_private_key_here"
DEPLOY_OUTPUT=$(./TruthValidatorAgent deploy --fevm)
CONTRACT_ADDRESS=$(echo "$DEPLOY_OUTPUT" | tail -n 1)
echo "Contract deployed at: $CONTRACT_ADDRESS"

# Test contract interaction
echo "Testing contract interaction..."
./TruthValidatorAgent submitProposal --fevm --address "$CONTRACT_ADDRESS" --content "Test proposal on FEVM"
./TruthValidatorAgent vote --fevm --address "$CONTRACT_ADDRESS" --proposalId 0 --approve --reason "Testing FEVM"

echo "FEVM tests completed successfully!"
