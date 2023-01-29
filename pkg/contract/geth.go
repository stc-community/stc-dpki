package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stc-community/stc-dpki/pkg/errors"
)

type Config struct {
	PrivateKey string
	Address    string
	RpcUrl     string
}

type EthClient struct {
	Client   *ethclient.Client
	Auth     *bind.TransactOpts
	Instance *Certificate
}

func NewEthClient(ctx context.Context, cfg *Config) (*EthClient, error) {
	client, err := InitGethClient(ctx, cfg)
	if err != nil {
		return nil, err
	}
	auth, err := InitGethAuth(ctx, client, cfg)
	if err != nil {
		return nil, err
	}
	contractAdd := common.HexToAddress(cfg.Address)
	instance, err := NewCertificate(contractAdd, client)
	if err != nil {
		return nil, err
	}
	result := &EthClient{
		Client:   client,
		Auth:     auth,
		Instance: instance,
	}
	return result, err
}

func InitGethClient(ctx context.Context, cfg *Config) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, cfg.RpcUrl)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return client, nil
}

func InitGethAuth(ctx context.Context, client *ethclient.Client, cfg *Config) (*bind.TransactOpts, error) {
	chanID, err := client.ChainID(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chanID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return auth, nil
}

// DeployContract 部署合约
func DeployContract(ctx context.Context, cfg *Config) error {
	fmt.Println("Rpc Url:", cfg.RpcUrl)
	fmt.Println("Private Key:", cfg.PrivateKey)
	fmt.Println()
	client, err := InitGethClient(ctx, cfg)
	if err != nil {
		return err
	}
	auth, err := InitGethAuth(ctx, client, cfg)
	if err != nil {
		return err
	}
	address, tx, _, err := DeployCertificate(auth, client)
	if err != nil {
		return err
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	return nil
}
