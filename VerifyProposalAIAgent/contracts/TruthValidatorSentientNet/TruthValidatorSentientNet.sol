// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TruthValidatorSentientNet {
    // Struct to store proposal details
    struct Proposal {
        uint256 id;          // Unique ID of the proposal
        string content;      // Content of the proposal
        uint256 yesVotes;    // Number of "yes" votes
        uint256 noVotes;     // Number of "no" votes
        bool isFinalized;    // Whether the proposal is finalized
    }

    // Struct to store vote details
    struct Vote {
        address voter;       // Address of the voter
        bool isApproved;     // Whether the vote is in favor (true) or against (false)
        string reason;       // Reason provided for the vote
    }

    address public admin;                              // Admin address with special privileges
    uint256 public proposalCounter;                   // Counter to generate unique proposal IDs
    mapping(uint256 => Proposal) public proposals;    // Mapping of proposal IDs to Proposal structs
    mapping(uint256 => mapping(address => Vote)) public voterVotes;  // Mapping of proposal IDs to voter addresses to Vote structs (replaces Vote array)
    uint256 public voteThreshold = 3;                  // Minimum number of votes required to finalize a proposal

    // Events
    event ProposalSubmitted(uint256 indexed proposalId, string content); // Emitted when a proposal is submitted
    event VoteCast(uint256 indexed proposalId, address voter, bool isApproved, string reason); // Emitted when a vote is cast
    event ProposalFinalized(uint256 indexed proposalId, bool finalResult, Vote[] voterResults); // Emitted when a proposal is finalized, including all voter results

    // Modifier to restrict access to the admin
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can call this function");
        _;
    }

    // Constructor to set the admin as the contract deployer
    constructor() {
        admin = msg.sender;
    }

    // Function to submit a new proposal
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

    mapping(uint256 => address[]) public proposalVoters;  // Mapping of proposal IDs to array of voter addresses

    // Function to cast a vote on a proposal
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

    // Internal function to finalize a proposal
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

    // Function to retrieve a voter's vote for a specific proposal
    function getVote(uint256 _proposalId, address _voter) external view returns (Vote memory) {
        return voterVotes[_proposalId][_voter];
    }

    function getVoters(uint256 _proposalId) external view returns (address[] memory) {
        return proposalVoters[_proposalId];
    }

    // Function to retrieve vote counts for a specific proposal
    function getVoteCounts(uint256 _proposalId) external view returns (uint256 yesVotes, uint256 noVotes) {
        return (proposals[_proposalId].yesVotes, proposals[_proposalId].noVotes);
    }

    // Function to set the vote threshold (only callable by the admin)
    function setVoteThreshold(uint256 _newThreshold) external onlyAdmin {
        require(_newThreshold > 0, "Threshold must be greater than 0");
        voteThreshold = _newThreshold;
    }
}
