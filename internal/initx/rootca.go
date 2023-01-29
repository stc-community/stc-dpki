package initx

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/stc-community/stc-fvm-dpki/internal/config"
	"github.com/stc-community/stc-fvm-dpki/pkg/util"
	"math/big"
	"os"
	"time"
)

// InitRootCert 初始化根证书
func InitRootCert() error {
	cfg := config.C.TLS
	if cfg.CertFile == "" && cfg.KeyFile == "" {
		// 初始化证书
		// 生成根pem文件
		config.C.TLS.CertFile = "./cert.crt"
		config.C.TLS.KeyFile = "./cert.key"
		ok, err := util.PathExists(cfg.CertFile)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		err = GenerateRootPemFile("Deca", cfg.CertFile, cfg.KeyFile)
		if err != nil {
			return fmt.Errorf("生成根证书文件失败：%v", err)
		}
	}
	return nil
}

// GenerateRootPemFile 生成新的根证书
func GenerateRootPemFile(host, certFile, keyFile string) error {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country:            []string{"CN"},         // 证书所属的国家
			Organization:       []string{"company"},    // 证书存放的公司名称
			OrganizationalUnit: []string{"department"}, // 证书所属的部门名称
			Province:           []string{"BeiJing"},    // 证书签发机构所在省
			CommonName:         host,
			Locality:           []string{"BeiJing"}, // 证书签发机构所在市
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		Issuer: pkix.Name{
			CommonName: host,
		},
	}
	priKey, err := GenerateKeyPair()
	if err != nil {
		return err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, &priKey.PublicKey, priKey)
	if err != nil {
		return err
	}

	// 将私钥写入.key文件
	keyFd, _ := os.OpenFile(keyFile, os.O_WRONLY|os.O_CREATE, os.ModePerm.Perm())
	defer func() {
		err = keyFd.Close()
	}()
	keyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priKey),
	}
	_ = pem.Encode(keyFd, keyBlock)
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	}
	// 将证书写入.crt文件
	certFd, _ := os.OpenFile(certFile, os.O_WRONLY|os.O_CREATE, os.ModePerm.Perm())
	defer func() {
		_ = certFd.Close()
	}()
	_ = pem.Encode(certFd, certBlock)
	return nil
}

// GenerateKeyPair 生成一对具有指定字位数的RSA密钥
func GenerateKeyPair() (*rsa.PrivateKey, error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, errors.New("密钥对生成失败")
	}
	return priKey, nil
}
