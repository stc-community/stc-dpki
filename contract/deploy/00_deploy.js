require("hardhat-deploy")
require("hardhat-deploy-ethers")

const { networkConfig } = require("../helper-hardhat-config")


const private_key = network.config.accounts[0]
const wallet = new ethers.Wallet(private_key, ethers.provider)

module.exports = async ({ deployments }) => {
    console.log("Wallet Ethereum Address:", wallet.address)
    const chainId = network.config.chainId
    const tokensToBeMinted = networkConfig[chainId]["tokensToBeMinted"]

    //deploy Certificate
    const Certificate = await ethers.getContractFactory('Certificate', wallet);
    console.log('Deploying Certificate...');
    const certificate = await Certificate.deploy();
    await certificate.deployed()
    console.log('Certificate deployed to:', certificate.address);
}