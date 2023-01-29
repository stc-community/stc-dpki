// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CertificateCert is an auto generated low-level Go binding around an user-defined struct.
type CertificateCert struct {
	Ski        string
	Aki        string
	Status     uint8
	Cid        string
	CidDocHash string
}

// CertificateMetaData contains all meta data concerning the Certificate contract.
var CertificateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sn\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ski\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aki\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidDocHash\",\"type\":\"string\"}],\"internalType\":\"structCertificate.Cert\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sn\",\"type\":\"string\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sn\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ski\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aki\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidDocHash\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"sn\",\"type\":\"string\"}],\"name\":\"verify\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ski\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"aki\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cidDocHash\",\"type\":\"string\"}],\"internalType\":\"structCertificate.Cert\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611556806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806365b2a86314610051578063693ec85e1461006d57806372f4a9bf1461009d578063bb9c6c3e146100b9575b600080fd5b61006b60048036038101906100669190610cb5565b6100e9565b005b61008760048036038101906100829190610cb5565b6101a8565b6040516100949190610e24565b60405180910390f35b6100b760048036038101906100b29190610e46565b610445565b005b6100d360048036038101906100ce9190610cb5565b610575565b6040516100e09190610e24565b60405180910390f35b60008151116100f757600080fd5b600080826040516101089190610f89565b908152602001604051809103902060020160009054906101000a900460ff1660ff160361016a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016190610ffd565b60405180910390fd5b600260008260405161017c9190610f89565b908152602001604051809103902060020160006101000a81548160ff021916908360ff16021790555050565b6101b0610b29565b6000826040516101c09190610f89565b90815260200160405180910390206040518060a00160405290816000820180546101e99061104c565b80601f01602080910402602001604051908101604052809291908181526020018280546102159061104c565b80156102625780601f1061023757610100808354040283529160200191610262565b820191906000526020600020905b81548152906001019060200180831161024557829003601f168201915b5050505050815260200160018201805461027b9061104c565b80601f01602080910402602001604051908101604052809291908181526020018280546102a79061104c565b80156102f45780601f106102c9576101008083540402835291602001916102f4565b820191906000526020600020905b8154815290600101906020018083116102d757829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660ff1660ff16815260200160038201805461032a9061104c565b80601f01602080910402602001604051908101604052809291908181526020018280546103569061104c565b80156103a35780601f10610378576101008083540402835291602001916103a3565b820191906000526020600020905b81548152906001019060200180831161038657829003601f168201915b505050505081526020016004820180546103bc9061104c565b80601f01602080910402602001604051908101604052809291908181526020018280546103e89061104c565b80156104355780601f1061040a57610100808354040283529160200191610435565b820191906000526020600020905b81548152906001019060200180831161041857829003601f168201915b5050505050815250509050919050565b600085511161045357600080fd5b600084511161046157600080fd5b600082511161046f57600080fd5b600081511161047d57600080fd5b6040518060a00160405280858152602001848152602001600160ff168152602001838152602001828152506000866040516104b89190610f89565b908152602001604051809103902060008201518160000190816104db9190611233565b5060208201518160010190816104f19190611233565b5060408201518160020160006101000a81548160ff021916908360ff16021790555060608201518160030190816105289190611233565b50608082015181600401908161053e9190611233565b50905050846001856040516105539190610f89565b9081526020016040518091039020908161056d9190611233565b505050505050565b61057d610b29565b600082511161058b57600080fd5b6000808360405161059c9190610f89565b908152602001604051809103902060020160009054906101000a900460ff1660ff161415826040516020016105d1919061132b565b60405160208183030381529060405290610621576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610618919061138a565b60405180910390fd5b50600160ff166000836040516106379190610f89565b908152602001604051809103902060020160009054906101000a900460ff1660ff16148260008460405161066b9190610f89565b908152602001604051809103902060000160405160200161068d92919061147b565b604051602081830303815290604052906106dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d4919061138a565b60405180910390fd5b50600080836040516106ef9190610f89565b9081526020016040518091039020600101805461070b9061104c565b9050111561089457600060016000846040516107279190610f89565b908152602001604051809103902060010160405161074591906114bd565b9081526020016040518091039020805461075e9061104c565b9050118260405160200161077291906114fa565b604051602081830303815290604052906107c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b9919061138a565b60405180910390fd5b5061089260016000846040516107d89190610f89565b90815260200160405180910390206001016040516107f691906114bd565b9081526020016040518091039020805461080f9061104c565b80601f016020809104026020016040519081016040528092919081815260200182805461083b9061104c565b80156108885780601f1061085d57610100808354040283529160200191610888565b820191906000526020600020905b81548152906001019060200180831161086b57829003601f168201915b5050505050610575565b505b6000826040516108a49190610f89565b90815260200160405180910390206040518060a00160405290816000820180546108cd9061104c565b80601f01602080910402602001604051908101604052809291908181526020018280546108f99061104c565b80156109465780601f1061091b57610100808354040283529160200191610946565b820191906000526020600020905b81548152906001019060200180831161092957829003601f168201915b5050505050815260200160018201805461095f9061104c565b80601f016020809104026020016040519081016040528092919081815260200182805461098b9061104c565b80156109d85780601f106109ad576101008083540402835291602001916109d8565b820191906000526020600020905b8154815290600101906020018083116109bb57829003601f168201915b505050505081526020016002820160009054906101000a900460ff1660ff1660ff168152602001600382018054610a0e9061104c565b80601f0160208091040260200160405190810160405280929190818152602001828054610a3a9061104c565b8015610a875780601f10610a5c57610100808354040283529160200191610a87565b820191906000526020600020905b815481529060010190602001808311610a6a57829003601f168201915b50505050508152602001600482018054610aa09061104c565b80601f0160208091040260200160405190810160405280929190818152602001828054610acc9061104c565b8015610b195780601f10610aee57610100808354040283529160200191610b19565b820191906000526020600020905b815481529060010190602001808311610afc57829003601f168201915b5050505050815250509050919050565b6040518060a001604052806060815260200160608152602001600060ff16815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bc282610b79565b810181811067ffffffffffffffff82111715610be157610be0610b8a565b5b80604052505050565b6000610bf4610b5b565b9050610c008282610bb9565b919050565b600067ffffffffffffffff821115610c2057610c1f610b8a565b5b610c2982610b79565b9050602081019050919050565b82818337600083830152505050565b6000610c58610c5384610c05565b610bea565b905082815260208101848484011115610c7457610c73610b74565b5b610c7f848285610c36565b509392505050565b600082601f830112610c9c57610c9b610b6f565b5b8135610cac848260208601610c45565b91505092915050565b600060208284031215610ccb57610cca610b65565b5b600082013567ffffffffffffffff811115610ce957610ce8610b6a565b5b610cf584828501610c87565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610d38578082015181840152602081019050610d1d565b60008484015250505050565b6000610d4f82610cfe565b610d598185610d09565b9350610d69818560208601610d1a565b610d7281610b79565b840191505092915050565b600060ff82169050919050565b610d9381610d7d565b82525050565b600060a0830160008301518482036000860152610db68282610d44565b91505060208301518482036020860152610dd08282610d44565b9150506040830151610de56040860182610d8a565b5060608301518482036060860152610dfd8282610d44565b91505060808301518482036080860152610e178282610d44565b9150508091505092915050565b60006020820190508181036000830152610e3e8184610d99565b905092915050565b600080600080600060a08688031215610e6257610e61610b65565b5b600086013567ffffffffffffffff811115610e8057610e7f610b6a565b5b610e8c88828901610c87565b955050602086013567ffffffffffffffff811115610ead57610eac610b6a565b5b610eb988828901610c87565b945050604086013567ffffffffffffffff811115610eda57610ed9610b6a565b5b610ee688828901610c87565b935050606086013567ffffffffffffffff811115610f0757610f06610b6a565b5b610f1388828901610c87565b925050608086013567ffffffffffffffff811115610f3457610f33610b6a565b5b610f4088828901610c87565b9150509295509295909350565b600081905092915050565b6000610f6382610cfe565b610f6d8185610f4d565b9350610f7d818560208601610d1a565b80840191505092915050565b6000610f958284610f58565b915081905092915050565b600082825260208201905092915050565b7f436572746966696361746520646f6573206e6f74206578697374000000000000600082015250565b6000610fe7601a83610fa0565b9150610ff282610fb1565b602082019050919050565b6000602082019050818103600083015261101681610fda565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061106457607f821691505b6020821081036110775761107661101d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026110df7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826110a2565b6110e986836110a2565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061113061112b61112684611101565b61110b565b611101565b9050919050565b6000819050919050565b61114a83611115565b61115e61115682611137565b8484546110af565b825550505050565b600090565b611173611166565b61117e818484611141565b505050565b5b818110156111a25761119760008261116b565b600181019050611184565b5050565b601f8211156111e7576111b88161107d565b6111c184611092565b810160208510156111d0578190505b6111e46111dc85611092565b830182611183565b50505b505050565b600082821c905092915050565b600061120a600019846008026111ec565b1980831691505092915050565b600061122383836111f9565b9150826002028217905092915050565b61123c82610cfe565b67ffffffffffffffff81111561125557611254610b8a565b5b61125f825461104c565b61126a8282856111a6565b600060209050601f83116001811461129d576000841561128b578287015190505b6112958582611217565b8655506112fd565b601f1984166112ab8661107d565b60005b828110156112d3578489015182556001820191506020850194506020810190506112ae565b868310156112f057848901516112ec601f8916826111f9565b8355505b6001600288020188555050505b505050505050565b7f436572746966696361746520646f6573206e6f742065786973742c20736e3a00815250565b600061133682611305565b601f820191506113468284610f58565b915081905092915050565b600061135c82610cfe565b6113668185610fa0565b9350611376818560208601610d1a565b61137f81610b79565b840191505092915050565b600060208201905081810360008301526113a48184611351565b905092915050565b7f436572746966696361746520756e617574686f72697a65642c20736e3a000000815250565b7f2c20736b693a0000000000000000000000000000000000000000000000000000815250565b600081546114058161104c565b61140f8186610f4d565b9450600182166000811461142a576001811461143f57611472565b60ff1983168652811515820286019350611472565b6114488561107d565b60005b8381101561146a5781548189015260018201915060208101905061144b565b838801955050505b50505092915050565b6000611486826113ac565b601d820191506114968285610f58565b91506114a1826113d2565b6006820191506114b182846113f8565b91508190509392505050565b60006114c982846113f8565b915081905092915050565b7f436572746966696361746520536b6920646f6573206e6f742065786973743a00815250565b6000611505826114d4565b601f820191506115158284610f58565b91508190509291505056fea26469706673582212205508786503e002308095ee471630a032931b2665f175dd705aeeddcf37d9002c64736f6c63430008110033",
}

// CertificateABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateMetaData.ABI instead.
var CertificateABI = CertificateMetaData.ABI

// CertificateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertificateMetaData.Bin instead.
var CertificateBin = CertificateMetaData.Bin

// DeployCertificate deploys a new Ethereum contract, binding an instance of Certificate to it.
func DeployCertificate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Certificate, error) {
	parsed, err := CertificateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertificateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// Certificate is an auto generated Go binding around an Ethereum contract.
type Certificate struct {
	CertificateCaller     // Read-only binding to the contract
	CertificateTransactor // Write-only binding to the contract
	CertificateFilterer   // Log filterer for contract events
}

// CertificateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateSession struct {
	Contract     *Certificate      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateCallerSession struct {
	Contract *CertificateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CertificateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateTransactorSession struct {
	Contract     *CertificateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CertificateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateRaw struct {
	Contract *Certificate // Generic contract binding to access the raw methods on
}

// CertificateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateCallerRaw struct {
	Contract *CertificateCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateTransactorRaw struct {
	Contract *CertificateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificate creates a new instance of Certificate, bound to a specific deployed contract.
func NewCertificate(address common.Address, backend bind.ContractBackend) (*Certificate, error) {
	contract, err := bindCertificate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Certificate{CertificateCaller: CertificateCaller{contract: contract}, CertificateTransactor: CertificateTransactor{contract: contract}, CertificateFilterer: CertificateFilterer{contract: contract}}, nil
}

// NewCertificateCaller creates a new read-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateCaller(address common.Address, caller bind.ContractCaller) (*CertificateCaller, error) {
	contract, err := bindCertificate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateCaller{contract: contract}, nil
}

// NewCertificateTransactor creates a new write-only instance of Certificate, bound to a specific deployed contract.
func NewCertificateTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateTransactor, error) {
	contract, err := bindCertificate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTransactor{contract: contract}, nil
}

// NewCertificateFilterer creates a new log filterer instance of Certificate, bound to a specific deployed contract.
func NewCertificateFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateFilterer, error) {
	contract, err := bindCertificate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateFilterer{contract: contract}, nil
}

// bindCertificate binds a generic wrapper to an already deployed contract.
func bindCertificate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.CertificateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.CertificateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Certificate *CertificateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Certificate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Certificate *CertificateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Certificate *CertificateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Certificate.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateCaller) Get(opts *bind.CallOpts, sn string) (CertificateCert, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "get", sn)

	if err != nil {
		return *new(CertificateCert), err
	}

	out0 := *abi.ConvertType(out[0], new(CertificateCert)).(*CertificateCert)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateSession) Get(sn string) (CertificateCert, error) {
	return _Certificate.Contract.Get(&_Certificate.CallOpts, sn)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateCallerSession) Get(sn string) (CertificateCert, error) {
	return _Certificate.Contract.Get(&_Certificate.CallOpts, sn)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateCaller) Verify(opts *bind.CallOpts, sn string) (CertificateCert, error) {
	var out []interface{}
	err := _Certificate.contract.Call(opts, &out, "verify", sn)

	if err != nil {
		return *new(CertificateCert), err
	}

	out0 := *abi.ConvertType(out[0], new(CertificateCert)).(*CertificateCert)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateSession) Verify(sn string) (CertificateCert, error) {
	return _Certificate.Contract.Verify(&_Certificate.CallOpts, sn)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string sn) view returns((string,string,uint8,string,string))
func (_Certificate *CertificateCallerSession) Verify(sn string) (CertificateCert, error) {
	return _Certificate.Contract.Verify(&_Certificate.CallOpts, sn)
}

// Revoke is a paid mutator transaction binding the contract method 0x65b2a863.
//
// Solidity: function revoke(string sn) returns()
func (_Certificate *CertificateTransactor) Revoke(opts *bind.TransactOpts, sn string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "revoke", sn)
}

// Revoke is a paid mutator transaction binding the contract method 0x65b2a863.
//
// Solidity: function revoke(string sn) returns()
func (_Certificate *CertificateSession) Revoke(sn string) (*types.Transaction, error) {
	return _Certificate.Contract.Revoke(&_Certificate.TransactOpts, sn)
}

// Revoke is a paid mutator transaction binding the contract method 0x65b2a863.
//
// Solidity: function revoke(string sn) returns()
func (_Certificate *CertificateTransactorSession) Revoke(sn string) (*types.Transaction, error) {
	return _Certificate.Contract.Revoke(&_Certificate.TransactOpts, sn)
}

// Save is a paid mutator transaction binding the contract method 0x72f4a9bf.
//
// Solidity: function save(string sn, string ski, string aki, string cid, string cidDocHash) returns()
func (_Certificate *CertificateTransactor) Save(opts *bind.TransactOpts, sn string, ski string, aki string, cid string, cidDocHash string) (*types.Transaction, error) {
	return _Certificate.contract.Transact(opts, "save", sn, ski, aki, cid, cidDocHash)
}

// Save is a paid mutator transaction binding the contract method 0x72f4a9bf.
//
// Solidity: function save(string sn, string ski, string aki, string cid, string cidDocHash) returns()
func (_Certificate *CertificateSession) Save(sn string, ski string, aki string, cid string, cidDocHash string) (*types.Transaction, error) {
	return _Certificate.Contract.Save(&_Certificate.TransactOpts, sn, ski, aki, cid, cidDocHash)
}

// Save is a paid mutator transaction binding the contract method 0x72f4a9bf.
//
// Solidity: function save(string sn, string ski, string aki, string cid, string cidDocHash) returns()
func (_Certificate *CertificateTransactorSession) Save(sn string, ski string, aki string, cid string, cidDocHash string) (*types.Transaction, error) {
	return _Certificate.Contract.Save(&_Certificate.TransactOpts, sn, ski, aki, cid, cidDocHash)
}