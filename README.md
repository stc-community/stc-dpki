<p align="center">
<img src="https://user-images.githubusercontent.com/52234994/165200623-c60e956b-5805-4088-bf58-f97ebd8ae8b4.png" 
    width="30%" border="0" alt="CA">
</p>

# Stc-DPKI
一种基于区块链结构的 DPKI 系统的跨域身份认证架构，它将在线身份的控制权返回到他们所属的实体，使用智能合约来管理数字证书的生命周期，使用用户
自行生成的证书来减轻传统 CA 证书管理繁琐的操作和管理负担，以实现轻量级的认证架构，避免复杂的组织和管理，从而可以实现更加便捷高效的跨域信息共享
通过这样做，DPKI 解决了许多困扰传统公钥基础设施的可用性和安全性挑战。

DPKI 在 PKI 生命周期的每个阶段都有优势，它使在线身份的自主控制成为可能，并提供简单创建更强大的 SSL 证书，
在使用中，它可以帮助实体最终由于将公钥管理降级为安全去中心化，进行了加密数据存储。

# 系统架构
在此DPKI认证架构中，系统架构由六个主要实体组成：用户、颁发者、验证者、星际文件系统（IPFS）节点、区块链网络和服务提供商，该系统以多个发行方节点为主要区块链节点，组成去中心化网络。 

用户将身份信息提交给颁发者进行认证，颁发者对信息进行校验，将其所属的CA智能合约地址作为用户自生成证书的颁发者。 然后发行者生成可验证的证书凭据（VC）并将其存储在该智能合约中，以方便验证者快速验证信息的真实性。 使用区块链智能合约旨在提高证书生命周期（即颁发、凭证管理和验证）的自动化程度。 区块链确保系统没有单点故障、更高的可扩展性、更高的问责制和可靠性。 验证者可以直接从IPFS地址中提取证书，从区块链中提取VC来验证证书的真实性，可以减少验证者与颁发者之间的交互，有利于提高验证效率。
![image](https://user-images.githubusercontent.com/52234994/215308087-de3a4043-6269-4517-a5aa-7847c10e05d9.png)

# Quick Start
## Build
```
$ git clone git@github.com:stc-community/fvm-dpki.git
$ make
```

## Deploy contract
Deploy contracts through blockchain RPC addresses and account keys
```
$ bin/dpki deploy -u https://goerli.infura.io/v3/81fec17192394537a0408c47ef5ec0b8 -k <private key>
```
## Configuration
Obtain the deployed contract address, modify the contract information in the configuration file, and other necessary information
```toml
# Contract Configuration
[Contract]
# Contract address
Address = "<contract address>"
# RPC Address
RpcUrl = "https://goerli.infura.io/v3/81fec17192394537a0408c47ef5ec0b8"
# Account private key - used for signing requests
PrivateKey = "<private key>"
```

## Start DPKI
```
$ bin/dpki tls -c configs/config.toml
```

## DPKI SDK Demo
Support decentralized certificate application, certificate online status query, certificate chain verification

[Demo](https://github.com/stc-community/fvm-dpki-casdk/examples/vm-mtls)