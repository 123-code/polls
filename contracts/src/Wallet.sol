pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/Address.sol";
import "./IMyToken.sol"; 

contract WalletImplementation {
    using Address for address;

    address public owner;
    IMyToken public nftContract;

    function initialize(address _owner, address _nftContractAddress) external {
        require(owner == address(0), "Already initialized");
        owner = _owner;
        nftContract = IMyToken(_nftContractAddress);
    }

    function execute(address target, uint256 value, bytes calldata data) external {
        require(msg.sender == owner, "Not authorized");
        (bool success, ) = target.call{value: value}(data);
        require(success, "Transaction failed");
    }

    function mintNFT() external payable {
        require(msg.sender == owner, "Not authorized");
        require(msg.value >= nftContract.price(), "Insufficient payment");

        uint256 newTokenId = nftContract.mint{value: msg.value}(address(this));
    }


    function withdraw() external {
        require(msg.sender == owner, "Not authorized");
        payable(owner).transfer(address(this).balance);
    }

    receive() external payable {}
}