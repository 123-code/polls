// SPDX-License-Identifier: MIT
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
pragma solidity ^0.8.19;

// Contract Address: 0xF554f9646581F79aF174F144Ea4De42AE13AF9c1
// rpc: https://optimism-goerli.infura.io/v3/c24c8ebb1b7c447aa3e95e28e11e6532



contract MyToken is ERC721URIStorage {
    uint256 private _nextTokenId;
    
    mapping(address => uint) public tokencount;
    mapping(uint => address) public tokenids;
    
    event Burned(address indexed owner, uint256 indexed tokenId);
    
    constructor() ERC721("ECVote", "VTT") {
        _nextTokenId = 1;
    }
    
    function mint(address to) public payable returns (uint256) {
 
        
        uint256 newItemId = _nextTokenId;
        _mint(to, newItemId);
        
        _nextTokenId += 1;
        tokencount[to] += 1;
        tokenids[newItemId] = msg.sender;
        
        return newItemId;
    }
    
    function getnftbalance(address _requester) public view returns (uint256) {
        return balanceOf(_requester);
    }
    
    function TransferNFT(uint256 tokenId, address to) external {
        require(msg.sender == tokenids[tokenId], "You don't own this NFT");
        transferFrom(msg.sender, to, tokenId);
    }
    
    function burn(uint256 tokenId) external {
        require(msg.sender == ownerOf(tokenId), "You don't own this NFT");
        _burn(tokenId);
        emit Burned(msg.sender, tokenId);
    }
}