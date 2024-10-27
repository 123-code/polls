// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
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

    function execute(
        address target,
        uint256 value,
        bytes calldata data,
        bytes memory signature
    ) external {
        require(_verifySignature(target, value, data, signature), "Invalid signature");

        (bool success, ) = target.call{value: value}(data);
        require(success, "Transaction failed");
    }

    function mintNFT() external payable {
        require(msg.sender == owner, "Not authorized");

        uint256 price = nftContract.price();
        require(msg.value >= price, "Insufficient payment");
        require(address(this).balance >= msg.value, "Insufficient contract balance");

        uint256 newTokenId = nftContract.mint{value: msg.value}(address(this));
        require(newTokenId != 0, "Mint failed");
    }

    function withdraw() external {
        require(msg.sender == owner, "Not authorized");
        payable(owner).transfer(address(this).balance);
    }

    function _verifySignature(
        address target,
        uint256 value,
        bytes calldata data,
        bytes memory signature
    ) internal view returns (bool) {

        bytes32 messageHash = keccak256(abi.encodePacked(target, value, data, address(this)));


        bytes32 ethSignedMessageHash = MessageHashUtils.toEthSignedMessageHash(messageHash);


address signer = ECDSA.recover(ethSignedMessageHash, signature);


        return signer == owner;
    }

    receive() external payable {}
}