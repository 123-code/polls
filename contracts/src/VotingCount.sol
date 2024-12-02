// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VotingSystem {

    struct Candidate {
        string name;
        uint256 voteCount;
    }



    mapping(string=>uint256) public candidateIndex;

    Candidate[] public candidates;

  
    event VoteCast(address voter, uint256 candidateIndex);

    address public owner;


    constructor(string[] memory _candidateNames) {
        owner = msg.sender;

        for (uint256 i = 0; i < _candidateNames.length; i++) {
            candidates.push(Candidate({
                name: _candidateNames[i],
                voteCount: 0
            }));
            candidateIndex[_candidateNames[i]] = i;
        }
    }
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the contract owner can perform this action");
        _;
    }


    function vote(uint256 _candidateIndex) public {
 
        require(_candidateIndex < candidates.length, "Invalid candidate index");
        candidates[_candidateIndex].voteCount++;
  

        emit VoteCast(msg.sender, _candidateIndex);
    }

    function getLeader() public view returns (string memory, uint256) {
        require(candidates.length > 0, "No candidates available");
        
        uint256 leadingVoteCount = 0;
        uint256 leadingCandidateIndex = 0;

        for (uint256 i = 0; i < candidates.length; i++) {
            if (candidates[i].voteCount > leadingVoteCount) {
                leadingVoteCount = candidates[i].voteCount;
                leadingCandidateIndex = i;
            }
        }

        return (
            candidates[leadingCandidateIndex].name, 
            candidates[leadingCandidateIndex].voteCount
        );
    }

    function getAllCandidates() public view returns (Candidate[] memory) {
        return candidates;
    }

    function addCandidate(string memory _name) public onlyOwner {
        candidates.push(Candidate({
            name: _name,
            voteCount: 0
        }));
    }
}
