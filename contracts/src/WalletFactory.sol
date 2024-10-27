pragma solidity ^0.8.19;

import "@openzeppelin/contracts/proxy/Clones.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WalletFactory is Ownable{
    address public implementationContract;
    address public nftContractAddress;
    mapping(string=>address) public userWallets;

    event WalletCreated(string indexed userId,address walletAddress);

    constructor(address _implementationContract,address _nftContractAddress) Ownable(msg.sender){
        implementationContract = _implementationContract;
        nftContractAddress = _nftContractAddress;
    }

    function createWallet(string memory userId,address userAddress) external onlyOwner returns(address){
        require(userWallets[userId] == address(0),"Biletera ya existe para este usuario");
        address wallet = Clones.clone(implementationContract);
        userWallets[userId] = wallet;
        (bool success, ) = wallet.call(
            abi.encodeWithSignature(
                "initialize(address,address)",
                userAddress,
                nftContractAddress
            )
        );
        require(success,"fallo de inicializacion de billetera");
        emit WalletCreated(userId,wallet);
        return wallet;
    }

}