// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import "./erc20/Erc20.sol";

contract CertificateDao is ERC20 {

    // Dao pledge
    mapping (address => uint) _daoDeposits;

    // dao and certificate mapping
    mapping (address => string[]) _daoCn;

    mapping(string => userWallet) userWallets;

    //Initialize variables
    uint public _daoDepositAmount;

    struct userWallet {
        address user;
        uint8 status;
    }

    constructor() ERC20("DPKI Coin", "DPK") public {
        _mint(msg.sender, 100000000 * 10 ** decimals());
        _daoDepositAmount = 5000 * 10 ** decimals();
    }

    function getWallet(string memory uuid) external view returns(address, uint8){
        return (userWallets[uuid].user, userWallets[uuid].status);
    }

    function bindWallet(string memory uuid) external {
        require(userWallets[uuid].user == address(0));
        if (_daoDeposits[msg.sender] == 0) {
            userWallets[uuid] = userWallet(msg.sender, 1);
        }else{
            userWallets[uuid] = userWallet(msg.sender, 2);
        }
    }

    function unbindWallet(string memory uuid) external {
        require(userWallets[uuid].user == msg.sender);
        delete userWallets[uuid];
    }

    function verifyWallet(string memory uuid) external {
        require(_daoDeposits[msg.sender] > 0);
        require(userWallets[uuid].status == 1);
        userWallets[uuid].status = 2;
    }

    function changeWallet(string memory uuid, address newWallet) external {
        require(newWallet != address(0));
        if (userWallets[uuid].status == 1){
            userWallets[uuid].user = newWallet;
        }else{
            require(userWallets[uuid].user == msg.sender);
            userWallets[uuid].user = newWallet;
        }
    }

    function getUserInfo(string memory uuid) external view returns(bool){
        if(userWallets[uuid].status == 2){
            return _daoDeposits[userWallets[uuid].user] > 0;
        }else{
            return false;
        }
    }
    // /**
    //  *
    //  */
    function isDeposit() external view returns (bool) {
        return _daoDeposits[msg.sender] != 0;
    }

    // /**
    //  *
    //  */
    function getDeposit(address walletAddress) external view returns (uint) {
        return _daoDeposits[walletAddress];
    }

    function stakeAmount(address walletAddress, uint amount) external {
        require(balanceOf(msg.sender) >= amount, "Not enough DPK");
        _transfer(msg.sender, address(this), amount);
        _daoDeposits[walletAddress] += amount;
    }
    // /**
    //  *
    //  */
    function stake() external {
        require(_daoDeposits[msg.sender] == 0, "Already staked");
        require(balanceOf(msg.sender) >= _daoDepositAmount, "Not enough CSD");
        _transfer(msg.sender, address(this), _daoDepositAmount);
        _daoDeposits[msg.sender] += _daoDepositAmount;
    }
    // /**
    //  *
    //  */
    function withdraw() external {
        require(_daoDeposits[msg.sender] > 0);
        _transfer(address(this), msg.sender, _daoDeposits[msg.sender]);
        delete _daoDeposits[msg.sender];
    }

}