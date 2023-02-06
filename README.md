<p align="center">
<img src="https://user-images.githubusercontent.com/52234994/215312990-066258c1-1864-4d13-ae69-492f76e6e401.png" 
    width="30%" border="0" alt="CA">
</p>

# STC-DPKI
A cross-domain identity authentication architecture of DPKI system based on blockchain structure, which returns the control of online identities to their entities, uses smart contracts to manage the life cycle of digital certificates, and uses user-generated certificates to reduce the cumbersome operation and management burden of traditional CA certificate management, so as to achieve lightweight authentication architecture and avoid complex organization and management, Thus, more convenient and efficient cross-domain information sharing can be achieved. By doing so, DPKI has solved many usability and security challenges that have plagued traditional public key infrastructure.

DPKI has advantages at every stage of the PKI life cycle. It makes the autonomous control of online identity possible, and provides a simple and more powerful SSL certificate. In use, it can help entities finally store encrypted data by degrading public key management to security decentralization.

# System Architecture 
In this DPKI authentication architecture, the system architecture is composed of six main entities: users, issuers, verifiers, interstellar file system (IPFS) nodes, blockchain networks and service providers. The system takes multiple issuer nodes as the main blockchain nodes, forming a decentralized network.
![image](https://user-images.githubusercontent.com/52234994/215308087-de3a4043-6269-4517-a5aa-7847c10e05d9.png)

The user submits the identity information to the issuer for authentication. The issuer verifies the information and takes the CA smart contract address to which it belongs as the issuer of the user's self-generated certificate. Then the issuer generates verifiable certificate credentials (VC) and stores them in the smart contract to facilitate the verifier to quickly verify the authenticity of the information. The use of blockchain smart contracts aims to improve the degree of automation of certificate lifecycle (i.e. issuance, certificate management and verification). Blockchain ensures that the system has no single point of failure, higher scalability, higher accountability and reliability. The verifier can extract the certificate directly from the IPFS address, and extract VC from the blockchain to verify the authenticity of the certificate, which can reduce the interaction between the verifier and the issuer, and is conducive to improving the verification efficiency.

# How It's Made

The stc-dpki system is mainly composed of Dpki services, storage engines, smart contracts, and dpki-sdk components. The stc-dpki system takes multiple Dpki service nodes as the main blockchain nodes to form a decentralized network. The details are as follows:

1. DPKI service (core service in DPKI architecture)
Anyone can run a DPKI service, which carries the core of decentralized public key infrastructure, provides a decentralized authentication model based on blockchain structure that is compatible with traditional PKI digital identity standards, receives information from identity entities and issues ID cards, encapsulates certificates as VC information in smart contracts, and stores certificate public key information in IPFS.

2. Storage engine (identity metadata encrypted storage engine)
The Stc-dpki storage engine supports storage through Filecoin's web3.storege and IPFS services. It mainly stores the public key information of encrypted identity entities, reduces gas costs by reducing the volume of on-chain information, and has better confidentiality, integrity and availability.

3. Smart contract (identity life cycle management and DAO Tools model)
Stc-dpki manages the life cycle of digital certificates through smart contracts, including the storage, revocation, verification and other operations of certificate VC information, and replaces traditional certificate chain and trust chain verification through a verifiable registry structure, which solves the mutual trust problem of CA and improves verification efficiency.

4. DPKI SDK (trusted authentication tool for identity entities)
User entities use stc dpki sdk decentralized registration information and issue identity cards to provide an integrated trusted authentication tool based on blockchain smart contracts, which has good ease of use and authentication efficiency.

# Quick Start
## Build
```
$ git clone git@github.com:stc-community/fvm-dpki.git
$ make
```

## Deploy contract

```
cd contract
yarn install
```
Add your private key as an environment variable by running this command:
```
export PRIVATE_KEY='abcdef'
```
You can also add private key using .env file. Next we get the deployer address
```
yarn hardhat get-address
```
Go to the [Hyperspace testnet faucet](https://hyperspace.yoga/#faucet), and paste in the Ethereum address from the previous step. This will send some hyperspace testnet FIL to the account.

Wait for the hyperspace testnet FIL to be sent, then deploy the contract.

```
yarn hardhat deploy
```
This will compile all the contracts in the contracts folder and deploy them to the Hyperspace test network automatically!

## Configuration
Obtain the deployed contract address, modify the contract information in the configuration file, and other necessary information
```toml
# Contract Configuration
[Contract]
# Contract address
Address = "<contract address>"
# RPC Address
RpcUrl = "https://api.hyperspace.node.glif.io/rpc/v1"
# Account private key - used for signing requests
PrivateKey = "<private key>"
```

## Start DPKI
```
$ bin/dpki tls -c configs/config.toml
```

## DPKI SDK Demo
Support decentralized certificate application, certificate online status query, certificate chain verification

[Demo](https://github.com/stc-community/fvm-dpki-casdk/tree/main/test)
