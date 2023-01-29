// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.12;

/**
 * @title Certificate
 */
contract Certificate {

    uint8 constant CERT_STATUS_GOOD = 1;
    uint8 constant CERT_STATUS_REVOKE = 2;

    struct Cert {
        string ski;
        string aki;
        uint8 status; // 状态1:正常；2:吊销
        string cid;
        string cidDocHash;
    }

    // 证书存储结构: sn=>cert
    mapping (string => Cert) _certs;
    // 证书识别号映射: ski=>sn
    mapping (string => string) _skiSn;

    /**
     * @dev 存储证书信息
     * @param sn 证书识别号
     * @param cid 线下存储cid
     * @param cidDocHash 线下存储内容hash
     */
    function save(string memory sn, string memory ski, string memory aki, string memory cid, string memory cidDocHash) public {
        require(bytes(sn).length > 0);
        require(bytes(ski).length > 0);
        require(bytes(cid).length > 0);
        require(bytes(cidDocHash).length > 0);

        _certs[sn] = Cert(ski, aki, CERT_STATUS_GOOD, cid, cidDocHash);
        _skiSn[ski] = sn;
    }

    /**
     * @dev 吊销证书
     * @param sn 证书识别号
     */
    function revoke(string memory sn) public {
        require(bytes(sn).length > 0);
        require(_certs[sn].status != 0, "Certificate does not exist");
        _certs[sn].status = CERT_STATUS_REVOKE;
    }

    /**
     * @dev 获取证书信息
     * @param sn 证书识别号
     * @return cert 证书信息
     */
    function get(string memory sn) public view returns (Cert memory){
        return _certs[sn];
    }

    /**
     * @dev 验证证书-包括证书链及状态
     * @param sn 证书识别号
     */
    function verify(string memory sn) public view returns (Cert memory){
        require(bytes(sn).length > 0);
        require(_certs[sn].status != 0, string.concat("Certificate does not exist, sn:", sn));
        require(_certs[sn].status == CERT_STATUS_GOOD, string.concat("Certificate unauthorized, sn:", sn, ", ski:", _certs[sn].ski));
        if (bytes(_certs[sn].aki).length > 0) {
            require(bytes(_skiSn[_certs[sn].aki]).length > 0, string.concat("Certificate Ski does not exist:", sn));
            // 递归检测父级证书
            verify(_skiSn[_certs[sn].aki]);
        }
        return _certs[sn];
    }
}