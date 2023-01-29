<p align="center">
<img src="https://user-images.githubusercontent.com/52234994/165200623-c60e956b-5805-4088-bf58-f97ebd8ae8b4.png" 
    width="40%" border="0" alt="CA">
</p>

# DeCA
DeCA 是一个兼容 X.509 的去中心化的 PKI 框架。
DeCA 可执行 X.509 PKI 标准的所有关键功能，即注册、确认、撤销和验证 TLS 证书。
DeCA 兼容现有的 PKI 标准，即 X.509。它存储、颁发和验证 X.509 格式的证书内容，而不是创建其自定义实现。

# 背景
PKI 的安全性很大程度上依赖于这些第三方 CA 的可靠性，这对 PKI 来说是一个单点故障。
过去曾发生过多起流行的 CA 违规事件，其中 CA 的中心化运营模式因流氓证书的传播而引发了
大量的针对性攻击。

我们的目标是使 CA 池完全去中心化，并同时构建我们的去中心化解决方案与已建立的 
PKI 标准（即 X.509）合作，以实现有效的现实世界集成。

# 架构
DeCA利用IPFS CRDT技术提出了一个名为 DeCA 的去中心化 PKI 框架，
该框架在去中心化的 CA 组之间提供数据同步，隐匿同步策略，基础数据低延迟同步。
证书信息利用IPFS的特性存储在IPFS中，不可篡改，有效防止第三方攻击。
![image](https://user-images.githubusercontent.com/52234994/192089294-d5891f90-16ac-497d-9efe-a09eb38b0ced.png)

# Get Start
```
$ go get github.com/stc-community/fvm-dpki
$ make
$ bin/ca tls -c configs/config.toml
```

### SDK Installation
```
$ go get github.com/stc-community/fvm-dpki
```

CA SDK的经典用法是客户端和服务器使用CA中心颁发的证书进行加密通信。以下是客户端和服务器之间sdk的用法。

See：[Demo](https://github.com/FlowShield/casdk/tree/main/examples)

