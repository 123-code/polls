// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract MyToken is ERC721URIStorage {
 
    uint256 private _nextTokenId;

    constructor() ERC721("ECVote", "VTT") {
        _nextTokenId = 1;
    }

    mapping(address => uint) public tokencount;
    mapping(uint => address) public tokenids;
    uint256 public immutable price = 0.000001 ether;

    event Burned(address indexed owner, uint256 indexed tokenId);

    function mint(address _to) public returns (uint256) {
        uint256 newItemID = _nextTokenId; 
        _mint(_to, newItemID);
        // _setTokenURI(newItemID, TokenURI); 
        _nextTokenId += 1;

        tokencount[_to] += 1;
        tokenids[newItemID] = msg.sender;

        return newItemID;
    }

    function getnftbalance(address _requester) public view returns (uint256) {
        return balanceOf(_requester);
    }


    function TransferNFT(uint256 _tokenID, address _to) external {
        require(msg.sender == tokenids[_tokenID], "You don't own this NFT");
        transferFrom(msg.sender, _to, _tokenID);
    }

    function burn(uint256 _tokenID) external {
        require(msg.sender == ownerOf(_tokenID), "You don't own this NFT");
        _burn(_tokenID);
        emit Burned(msg.sender, _tokenID);
    }
}