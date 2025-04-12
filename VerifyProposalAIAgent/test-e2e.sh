#!/bin/sh
#
# TruthValidator End-to-End Test Script
#
# Tests the complete proposal validation workflow:
# 1. Contract deployment
# 2. Proposal submission
# 3. Manual voting
# 4. AI agent automated voting
#
# Requirements:
# - TruthValidatorAgent binary must be built
# - PRIVATE_KEY environment variable must be set
# - Local Ethereum node or testnet connection
#
# Usage: ./test-e2e.sh
#
# Exit codes:
# 0 - All tests passed
# 1 - Test failure
# 2 - Missing requirements

set -eu # Exit on error and undefined variables

# ========== Configuration ==========
# Private key of the first test account (must be set)
export PRIVATE_KEY=""

# Blockchain network to use (eth-localnet, sepolia, etc)
NETWORK="eth-localnet"

# Test proposal content
PROPOSAL_CONTENT="Should we implement enhanced privacy features?"

# ========== Test 1: Contract Deployment ==========
echo "=== Deploying contract ==="
ADDR=$(./TruthValidatorAgent deploy --network ${NETWORK})
echo "Contract deployed at: ${ADDR}"

# ========== Test 2: Proposal Submission ==========
echo "=== Submitting test proposal ==="
TX_HASH=$(./TruthValidatorAgent submitProposal --network ${NETWORK} "${ADDR}" "${PROPOSAL_CONTENT}")
echo "Proposal submitted with transaction hash: ${TX_HASH}"

# Wait for transaction confirmation
echo "Waiting for transaction to be mined (5 seconds)..."
sleep 5  # Adjust based on network speed

# ========== Test 3: Proposal Counter Verification ==========
echo "=== Verifying proposal counter ==="
COUNTER=$(./TruthValidatorAgent getProposalCounter --network ${NETWORK} "${ADDR}" | awk '/Proposal Counter: /{print $3}')
if [ "${COUNTER}" != "1" ]; then
    echo "Test failed! Expected proposal counter 1, got ${COUNTER}"
    exit 1
fi

PROPOSAL_ID=0  # First proposal has ID 0

# ========== Test 4: Manual Voting ==========
echo "=== Testing manual voting ==="
VOTE_TX_HASH=$(./TruthValidatorAgent vote --network ${NETWORK} "${ADDR}" "${PROPOSAL_ID}" "true" "The proposal is well-aligned with our goals.")
echo "Vote submitted with transaction hash: ${VOTE_TX_HASH}"

# Wait for vote transaction
echo "Waiting for vote transaction (5 seconds)..."
sleep 5

# ========== Test 5: Proposal Details Verification ==========
echo "=== Verifying proposal details ==="
PROPOSAL_DETAILS=$(./TruthValidatorAgent getProposal --network ${NETWORK} "${ADDR}" "${PROPOSAL_ID}")
echo "Proposal details:"
echo "${PROPOSAL_DETAILS}"

# Verify proposal content matches
PROPOSAL_CONTENT_GOT=$(echo "${PROPOSAL_DETAILS}" | awk -F ': ' '/Content:/ {print $2}')
if [ "${PROPOSAL_CONTENT_GOT}" != "${PROPOSAL_CONTENT}" ]; then
    echo "Test failed! Expected proposal content \"${PROPOSAL_CONTENT}\", got \"${PROPOSAL_CONTENT_GOT}\"."
    exit 1
fi

# Verify vote count incremented
YES_VOTES_GOT=$(echo "${PROPOSAL_DETAILS}" | awk -F ': ' '/Yes Votes:/ {print $2}')
if [ "${YES_VOTES_GOT}" != "1" ]; then
    echo "Test failed! Expected 1 Yes vote, got ${YES_VOTES_GOT}."
    exit 1
fi

# Verify vote reason matches
VOTE_DETAILS=$(echo "${PROPOSAL_DETAILS}" | grep "Voter:")
EXPECTED_REASON="The proposal is well-aligned with our goals."
REASON_GOT=$(echo "${VOTE_DETAILS}" | awk -F ', Reason: ' '{print $2}')
if [ "${REASON_GOT}" != "${EXPECTED_REASON}" ]; then
    echo "Test failed! Expected vote reason \"${EXPECTED_REASON}\", got \"${REASON_GOT}\"."
    exit 1
fi

# ========== Test 6: AI Agent Voting ==========
echo "=== Testing AI agent automated voting ==="
echo "Starting AI agent in background..."
./TruthValidatorAgent aigent_listenAndVote --network ${NETWORK} "${ADDR}" &
AI_AGENT_PID=$!
echo "AI agent started with PID: ${AI_AGENT_PID}"

# Wait for agent initialization
echo "Waiting for AI agent to initialize (5 seconds)..."
sleep 5

# Submit second test proposal
echo "=== Submitting second proposal for AI evaluation ==="
NEW_PROPOSAL_CONTENT="Should we upgrade the system to version 2.0?"
NEW_TX_HASH=$(./TruthValidatorAgent submitProposal --network ${NETWORK} "${ADDR}" "${NEW_PROPOSAL_CONTENT}")
if [ -z "${NEW_TX_HASH}" ]; then
    echo "Failed to submit new proposal. Cleaning up..."
    kill ${AI_AGENT_PID} || true
    exit 1
fi
echo "New proposal submitted with transaction hash: ${NEW_TX_HASH}"

# Wait for mining and AI processing
echo "Waiting for AI agent to process and vote (15 seconds)..."
sleep 15

# Verify AI agent voted correctly
NEW_PROPOSAL_ID=1  # Second proposal has ID 1
echo "=== Verifying AI agent vote ==="
NEW_PROPOSAL_DETAILS=$(./TruthValidatorAgent getProposal --network ${NETWORK} "${ADDR}" "${NEW_PROPOSAL_ID}")
if [ -z "${NEW_PROPOSAL_DETAILS}" ]; then
    echo "Failed to retrieve new proposal details. Cleaning up..."
    kill ${AI_AGENT_PID} || true
    exit 1
fi
echo "New proposal details:"
echo "${NEW_PROPOSAL_DETAILS}"

# Verify proposal content
NEW_PROPOSAL_CONTENT_GOT=$(echo "${NEW_PROPOSAL_DETAILS}" | awk -F ': ' '/Content:/ {print $2}')
if [ "${NEW_PROPOSAL_CONTENT_GOT}" != "${NEW_PROPOSAL_CONTENT}" ]; then
    echo "Test failed! Expected new proposal content \"${NEW_PROPOSAL_CONTENT}\", got \"${NEW_PROPOSAL_CONTENT_GOT}\"."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

# Verify AI vote was recorded
NEW_YES_VOTES_GOT=$(echo "${NEW_PROPOSAL_DETAILS}" | awk -F ': ' '/Yes Votes:/ {print $2}")
if [ "${NEW_YES_VOTES_GOT}" != "1" ]; then
    echo "Test failed! Expected 1 Yes vote from AI, got ${NEW_YES_VOTES_GOT}."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

# Verify AI reasoning
NEW_VOTE_DETAILS=$(echo "${NEW_PROPOSAL_DETAILS}" | grep "Voter:")
NEW_EXPECTED_REASON="The proposal contains keywords related to enhancement: ${NEW_PROPOSAL_CONTENT}"
NEW_REASON_GOT=$(echo "${NEW_VOTE_DETAILS}" | awk -F ', Reason: ' '{print $2}')
if [ "${NEW_REASON_GOT}" != "${NEW_EXPECTED_REASON}" ]; then
    echo "Test failed! Expected AI vote reason \"${NEW_EXPECTED_REASON}\", got \"${NEW_REASON_GOT}\"."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

# Clean up and exit
echo "=== All tests passed successfully! ==="
kill ${AI_AGENT_PID} || true
exit 0
