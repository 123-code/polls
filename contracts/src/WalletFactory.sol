pragma solidity ^0.8.19;

import "@openzeppelin/contracts/proxy/Clones.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WalletFactory is Ownable {
    address public implementationContract;
    address public nftContractAddress;
    mapping(string => address) public userWallets;

    event WalletCreated(string indexed userId, address walletAddress);

    constructor(address _implementationContract, address _nftContractAddress) Ownable(msg.sender) {
        implementationContract = _implementationContract;
        nftContractAddress = _nftContractAddress;
    }

    function createWallet(string memory userId) external onlyOwner returns (address) {
        require(userWallets[userId] == address(0), "Wallet already exists for this user");

        address wallet = Clones.clone(implementationContract);
        userWallets[userId] = wallet;

        (bool success, ) = wallet.call(abi.encodeWithSignature("initialize(address,address)", msg.sender, nftContractAddress));
        require(success, "Wallet initialization failed");

        emit WalletCreated(userId, wallet);
        return wallet;
    }

    function updateImplementation(address newImplementation) external onlyOwner {
        implementationContract = newImplementation;
    }

    function updateNFTContract(address newNFTContract) external onlyOwner {
        nftContractAddress = newNFTContract;
    }
}