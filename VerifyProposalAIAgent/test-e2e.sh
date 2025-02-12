#!/bin/sh

set -eu

# Private key of the first test account.
export PRIVATE_KEY=""

# Network to use
NETWORK="eth-localnet"

# Proposal content
PROPOSAL_CONTENT="Should we implement enhanced privacy features?"

# Deploy the contract and get address
ADDR=$(./TruthValidatorAgent deploy --network ${NETWORK})
echo "Contract deployed at: ${ADDR}"

# Submit a new proposal
echo "Submitting proposal..."
TX_HASH=$(./TruthValidatorAgent submitProposal --network ${NETWORK} "${ADDR}" "${PROPOSAL_CONTENT}")
echo "Proposal submitted with transaction hash: ${TX_HASH}"

# Wait for the transaction to be mined
echo "Waiting for transaction to be mined..."
sleep 5  # Adjust the sleep time as needed

# Verify proposal counter
COUNTER=$(./TruthValidatorAgent getProposalCounter --network ${NETWORK} "${ADDR}" | awk '/Proposal Counter: /{print $3}')
if [ "${COUNTER}" != "1" ]; then
    echo "Test failed! Expected proposal counter 1, got ${COUNTER}"
    exit 1
fi

PROPOSAL_ID=0  # First proposal ID

# Vote on the proposal
echo "Voting on proposal..."
VOTE_TX_HASH=$(./TruthValidatorAgent vote --network ${NETWORK} "${ADDR}" "${PROPOSAL_ID}" "true" "The proposal is well-aligned with our goals.")
echo "Vote submitted with transaction hash: ${VOTE_TX_HASH}"

# Wait for the transaction to be mined
echo "Waiting for transaction to be mined..."
sleep 5  # Adjust the sleep time as needed

# Retrieve proposal details
echo "Retrieving proposal details..."
PROPOSAL_DETAILS=$(./TruthValidatorAgent getProposal --network ${NETWORK} "${ADDR}" "${PROPOSAL_ID}")
echo "Proposal details:"
echo "${PROPOSAL_DETAILS}"

# Verify the proposal content
PROPOSAL_CONTENT_GOT=$(echo "${PROPOSAL_DETAILS}" | awk -F ': ' '/Content:/ {print $2}')
if [ "${PROPOSAL_CONTENT_GOT}" != "${PROPOSAL_CONTENT}" ]; then
    echo "Test failed! Expected proposal content \"${PROPOSAL_CONTENT}\", got \"${PROPOSAL_CONTENT_GOT}\"."
    exit 1
fi

# Verify vote count
YES_VOTES_GOT=$(echo "${PROPOSAL_DETAILS}" | awk -F ': ' '/Yes Votes:/ {print $2}')
if [ "${YES_VOTES_GOT}" != "1" ]; then
    echo "Test failed! Expected 1 Yes vote, got ${YES_VOTES_GOT}."
    exit 1
fi

# Verify vote details (including Reason)
VOTE_DETAILS=$(echo "${PROPOSAL_DETAILS}" | grep "Voter:")
EXPECTED_REASON="The proposal is well-aligned with our goals."
REASON_GOT=$(echo "${VOTE_DETAILS}" | awk -F ', Reason: ' '{print $2}')
if [ "${REASON_GOT}" != "${EXPECTED_REASON}" ]; then
    echo "Test failed! Expected vote reason \"${EXPECTED_REASON}\", got \"${REASON_GOT}\"."
    exit 1
fi

# Test AI agent listening and voting
echo "Starting AI agent to listen and vote..."
# Run the AI agent in the background
./TruthValidatorAgent aigent_listenAndVote --network ${NETWORK} "${ADDR}" &
AI_AGENT_PID=$!

# Wait for the AI agent to start listening
echo "Waiting for the AI agent to start listening..."
sleep 5  # Adjust the sleep time as needed

# Submit another proposal
echo "Submitting another proposal..."
NEW_PROPOSAL_CONTENT="Should we upgrade the system to version 2.0?"
NEW_TX_HASH=$(./TruthValidatorAgent submitProposal --network ${NETWORK} "${ADDR}" "${NEW_PROPOSAL_CONTENT}")
if [ -z "${NEW_TX_HASH}" ]; then
    echo "Failed to submit the new proposal. Exiting..."
    kill ${AI_AGENT_PID} || true
    exit 1
fi
echo "New proposal submitted with transaction hash: ${NEW_TX_HASH}"

# Wait for the transaction to be mined
echo "Waiting for the transaction to be mined..."
sleep 5  # Adjust the sleep time as needed

# Wait for the AI agent to vote
echo "Waiting for the AI agent to vote..."
sleep 10  # Adjust the sleep time as needed

# Retrieve the new proposal details
NEW_PROPOSAL_ID=1  # Second proposal ID
echo "Retrieving new proposal details..."
NEW_PROPOSAL_DETAILS=$(./TruthValidatorAgent getProposal --network ${NETWORK} "${ADDR}" "${NEW_PROPOSAL_ID}")
if [ -z "${NEW_PROPOSAL_DETAILS}" ]; then
    echo "Failed to retrieve new proposal details. Exiting..."
    kill ${AI_AGENT_PID} || true
    exit 1
fi
echo "New proposal details:"
echo "${NEW_PROPOSAL_DETAILS}"

# Verify the new proposal content
NEW_PROPOSAL_CONTENT_GOT=$(echo "${NEW_PROPOSAL_DETAILS}" | awk -F ': ' '/Content:/ {print $2}')
if [ "${NEW_PROPOSAL_CONTENT_GOT}" != "${NEW_PROPOSAL_CONTENT}" ]; then
    echo "Test failed! Expected new proposal content \"${NEW_PROPOSAL_CONTENT}\", got \"${NEW_PROPOSAL_CONTENT_GOT}\"."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

# Verify vote count for the new proposal
NEW_YES_VOTES_GOT=$(echo "${NEW_PROPOSAL_DETAILS}" | awk -F ': ' '/Yes Votes:/ {print $2}')
if [ "${NEW_YES_VOTES_GOT}" != "1" ]; then
    echo "Test failed! Expected 1 Yes vote for the new proposal, got ${NEW_YES_VOTES_GOT}."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

# Verify vote details (including Reason) for the new proposal
NEW_VOTE_DETAILS=$(echo "${NEW_PROPOSAL_DETAILS}" | grep "Voter:")
NEW_EXPECTED_REASON="The proposal contains keywords related to enhancement: ${NEW_PROPOSAL_CONTENT}"
NEW_REASON_GOT=$(echo "${NEW_VOTE_DETAILS}" | awk -F ', Reason: ' '{print $2}')
if [ "${NEW_REASON_GOT}" != "${NEW_EXPECTED_REASON}" ]; then
    echo "Test failed! Expected vote reason \"${NEW_EXPECTED_REASON}\", got \"${NEW_REASON_GOT}\"."
    kill ${AI_AGENT_PID} || true
    exit 1
fi

echo "All tests passed!"
kill ${AI_AGENT_PID} || true
exit 0
