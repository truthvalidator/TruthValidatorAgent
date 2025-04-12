// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title TruthValidatorSentientNet
 * @notice A decentralized truth verification system where:
 * - Participants submit proposals about facts/claims
 * - Community members vote on proposal validity
 * - AI agents can participate as voters
 * - Results are transparently recorded on-chain
 */
contract TruthValidatorSentientNet {
    /// @notice Proposal structure capturing all details of a truth claim
    /// @dev Stored in proposals mapping by proposalId
    struct Proposal {
        uint256 id;          /// @notice Auto-incrementing unique proposal ID
        string content;      /// @notice The actual claim/content being verified
        uint256 yesVotes;    /// @notice Count of validating votes
        uint256 noVotes;     /// @notice Count of invalidating votes
        bool isFinalized;    /// @notice Flag indicating final voting status
    }

    /// @notice Vote structure capturing individual voting decisions
    /// @dev Stored in voterVotes nested mapping (proposalId => voter => Vote)
    struct Vote {
        address voter;       /// @notice Voter's Ethereum address
        bool isApproved;     /// @notice True if vote validates the claim
        string reason;       /// @notice Justification for the vote decision
    }

    address public admin;                              /// @notice Privileged address for admin functions
    uint256 public proposalCounter;                   /// @notice Auto-incrementing counter for proposal IDs
    mapping(uint256 => Proposal) public proposals;    /// @notice Lookup of all proposals by ID
    mapping(uint256 => mapping(address => Vote)) public voterVotes;  /// @notice Nested vote tracking (proposalId => voter => vote)
    uint256 public voteThreshold = 3;                  /// @notice Minimum votes needed to finalize a proposal

    /// @notice Emitted when a new truth claim is submitted
    /// @param proposalId The unique ID assigned to this proposal
    /// @param content The actual claim/content being verified
    event ProposalSubmitted(uint256 indexed proposalId, string content);

    /// @notice Emitted when a participant votes on a proposal
    /// @param proposalId The proposal being voted on
    /// @param voter The address casting the vote
    /// @param isApproved True if vote validates the claim
    /// @param reason Justification for the vote decision
    event VoteCast(uint256 indexed proposalId, address voter, bool isApproved, string reason);

    /// @notice Emitted when a proposal reaches vote threshold and is finalized
    /// @param proposalId The finalized proposal ID
    /// @param finalResult True if majority validated the claim
    /// @param voterResults Array of all votes cast on this proposal
    event ProposalFinalized(uint256 indexed proposalId, bool finalResult, Vote[] voterResults);

    /// @notice Restricts function access to only the admin
    modifier onlyAdmin() {
        require(msg.sender == admin, "TruthValidator: caller is not admin");
        _;
    }

    /// @notice Initializes contract with deployer as admin
    constructor() {
        admin = msg.sender;
    }

    /// @notice Allows anyone to submit a new truth claim for verification
    /// @param _content The factual claim being submitted for community validation
    /// @dev Emits ProposalSubmitted event on success
    function submitProposal(string memory _content) external {
        uint256 proposalId = proposalCounter++; // Generate a new proposal ID
        proposals[proposalId] = Proposal({
            id: proposalId,
            content: _content,
            yesVotes: 0,
            noVotes: 0,
            isFinalized: false
        });
        emit ProposalSubmitted(proposalId, _content); // Emit the ProposalSubmitted event
    }

    mapping(uint256 => address[]) public proposalVoters;  /// @notice Tracks all voters per proposal

    /// @notice Casts a vote on a proposal with justification
    /// @param _proposalId The proposal being voted on
    /// @param _isApproved True if vote validates the claim
    /// @param _reason Detailed justification for the vote
    /// @dev Emits VoteCast event and may trigger finalization
    function vote(uint256 _proposalId, bool _isApproved, string memory _reason) external {
        Proposal storage p = proposals[_proposalId]; // Get the proposal
        require(!p.isFinalized, "Proposal is finalized"); // Ensure the proposal is not already finalized
        require(voterVotes[_proposalId][msg.sender].voter == address(0), "Already voted"); // Check if the voter has already voted

        // Record the vote
        voterVotes[_proposalId][msg.sender] = Vote({
            voter: msg.sender,
            isApproved: _isApproved,
            reason: _reason
        });

        // Update the vote count
        if (_isApproved) {
            p.yesVotes++;
        } else {
            p.noVotes++;
        }

        // Add the voter to the list of voters for this proposal
        proposalVoters[_proposalId].push(msg.sender);

        emit VoteCast(_proposalId, msg.sender, _isApproved, _reason); // Emit the VoteCast event

        // Finalize the proposal if the total votes reach the threshold
        if (p.yesVotes + p.noVotes >= voteThreshold) {
            _finalizeProposal(_proposalId);
        }
    }

    /// @notice Internal function that finalizes a proposal when threshold reached
    /// @param _proposalId The proposal to finalize
    /// @dev Called automatically when vote threshold met
    function _finalizeProposal(uint256 _proposalId) internal {
        Proposal storage p = proposals[_proposalId];
        require(!p.isFinalized, "Proposal is already finalized");  // Ensure the proposal isn't already finalized

        p.isFinalized = true;
        bool finalResult = (p.yesVotes > p.noVotes);

        // Collect all voter results
        Vote[] memory voterResults = new Vote[](proposalVoters[_proposalId].length);
        for (uint256 i = 0; i < proposalVoters[_proposalId].length; i++) {
            address voter = proposalVoters[_proposalId][i];
            voterResults[i] = voterVotes[_proposalId][voter];
        }

        emit ProposalFinalized(_proposalId, finalResult, voterResults); // Emit the ProposalFinalized event with voter results
    }

    /// @notice Retrieves a specific vote from a voter on a proposal
    /// @param _proposalId The proposal to query
    /// @param _voter The voter's address
    /// @return Vote The voter's decision and reasoning
    function getVote(uint256 _proposalId, address _voter) external view returns (Vote memory) {
        return voterVotes[_proposalId][_voter];
    }

    /// @notice Gets all voters who participated in a proposal
    /// @param _proposalId The proposal to query
    /// @return address[] Array of voter addresses
    function getVoters(uint256 _proposalId) external view returns (address[] memory) {
        return proposalVoters[_proposalId];
    }

    /// @notice Gets the current vote tally for a proposal
    /// @param _proposalId The proposal to query
    /// @return yesVotes Count of validating votes
    /// @return noVotes Count of invalidating votes
    function getVoteCounts(uint256 _proposalId) external view returns (uint256 yesVotes, uint256 noVotes) {
        return (proposals[_proposalId].yesVotes, proposals[_proposalId].noVotes);
    }

    /// @notice Admin function to update the vote threshold
    /// @param _newThreshold New minimum votes required (must be >0)
    /// @dev Only callable by admin
    function setVoteThreshold(uint256 _newThreshold) external onlyAdmin {
        require(_newThreshold > 0, "TruthValidator: threshold must be >0");
        voteThreshold = _newThreshold;
    }
}
