package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	cfssl_config "github.com/flowshield/cfssl/config"

	"github.com/koding/multiconfig"

	"github.com/stc-community/stc-fvm-dpki/pkg/util/json"
)

var (
	// C Global configuration (Must Load first, otherwise the configuration will not be available)
	C    = new(Config)
	once sync.Once
)

// MustLoad load config
func MustLoad(fpaths ...string) error {
	once.Do(func() {
		loaders := []multiconfig.Loader{
			&multiconfig.TagLoader{},
			&multiconfig.EnvironmentLoader{},
		}

		for _, fpath := range fpaths {
			if strings.HasSuffix(fpath, "toml") {
				loaders = append(loaders, &multiconfig.TOMLLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "json") {
				loaders = append(loaders, &multiconfig.JSONLoader{Path: fpath})
			}
			if strings.HasSuffix(fpath, "yaml") {
				loaders = append(loaders, &multiconfig.YAMLLoader{Path: fpath})
			}
		}
		m := multiconfig.DefaultLoader{
			Loader:    multiconfig.MultiLoader(loaders...),
			Validator: multiconfig.MultiValidator(&multiconfig.RequiredValidator{}),
		}
		m.MustLoad(C)
	})

	err := ParseConfigByEnv()
	if err != nil {
		return err
	}
	PrintWithJSON()
	return nil
}

func ParseConfigByEnv() error {
	// TLS
	if v := os.Getenv("CA_TLS_HOST"); v != "" {
		C.TLS.Host = v
	}
	if v := os.Getenv("CA_TLS_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("environment variable [%s] parsing error:%v", "CA_TLS_PORT", err)
		}
		C.TLS.Port = p
	}
	if v := os.Getenv("CA_TLS_CERT_FILE"); v != "" {
		C.TLS.CertFile = v
	}
	if v := os.Getenv("CA_TLS_KEY_FILE"); v != "" {
		C.TLS.KeyFile = v
	}

	// Contract
	if v := os.Getenv("CA_CONTRACT_ADDRESS"); v != "" {
		C.Contract.Address = v
	}
	if v := os.Getenv("CA_CONTRACT_RPC_URL"); v != "" {
		C.Contract.RpcUrl = v
	}
	if v := os.Getenv("CA_CONTRACT_PRIVATE_KEY"); v != "" {
		C.Contract.PrivateKey = v
	}

	// OCSP
	if v := os.Getenv("CA_OCSP_HOST"); v != "" {
		C.OCSP.Host = v
	}
	if v := os.Getenv("CA_OCSP_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("environment variable [%s] parsing error:%v", "CA_OCSP_PORT", err)
		}
		C.OCSP.Port = p
	}
	// Log
	if v := os.Getenv("CA_LOG_HOOK_ENABLED"); v == "true" {
		C.Log.EnableHook = true
	}
	if v := os.Getenv("CA_LOG_REDIS_ADDR"); v != "" {
		C.LogRedisHook.Addr = v
	}
	if v := os.Getenv("CA_LOG_REDIS_KEY"); v != "" {
		C.LogRedisHook.Key = v
	}
	// Storage
	if v := os.Getenv("CA_STORAGE_TYPE"); v != "" {
		C.Storage.Type = v
	}
	// IPFS
	if v := os.Getenv("CA_IPFS_HOST"); v != "" {
		C.Ipfs.Host = v
	}
	if v := os.Getenv("CA_IPFS_PORT"); v != "" {
		p, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("environment variable [%s] parsing error:%v", "CA_IPFS_PORT", err)
		}
		C.Ipfs.Port = p
	}

	// Cfssl
	cfg, err := cfssl_config.LoadFile(C.Cfssl.ConfigFile)
	if err != nil {
		return err
	}
	cfg.Signing.Default.OCSP = C.Cfssl.OCSPHost
	C.Cfssl.Config = cfg
	return nil
}

func PrintWithJSON() {
	if C.PrintConfig {
		b, err := json.MarshalIndent(C, "", " ")
		if err != nil {
			os.Stdout.WriteString("[CONFIG] JSON marshal error: " + err.Error())
			return
		}
		os.Stdout.WriteString(string(b) + "\n")
	}
}

type Config struct {
	RunMode      string
	PrintConfig  bool
	TLS          TLS
	OCSP         OCSP
	Log          Log
	LogRedisHook LogRedisHook
	Cfssl        Cfssl
	Storage      Storage
	Ipfs         Ipfs
	Web3Storage  Web3Storage
	Contract     Contract
}

func (c *Config) IsDebugMode() bool {
	return c.RunMode == "debug"
}

func (c *Config) IsReleaseMode() bool {
	return c.RunMode == "release"
}

type LogHook string

func (h LogHook) IsRedis() bool {
	return h == "redis"
}

type Log struct {
	Level         int
	Format        string
	Output        string
	OutputFile    string
	EnableHook    bool
	HookLevels    []string
	Hook          LogHook
	HookMaxThread int
	HookMaxBuffer int
}

type LogGormHook struct {
	DBType       string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	Table        string
}

type LogRedisHook struct {
	Addr string
	Key  string
}

type Redis struct {
	Addr     string
	Password string
}

// TLS Configuration parameters
type TLS struct {
	Host               string
	Port               int
	CertFile           string
	KeyFile            string
	Cert               []byte
	Key                []byte
	ShutdownTimeout    int
	MaxContentLength   int64
	MaxReqLoggerLength int `default:"1024"`
	MaxResLoggerLength int
}

// OCSP Configuration parameters
type OCSP struct {
	Host               string
	Port               int
	ShutdownTimeout    int
	MaxContentLength   int64
	MaxReqLoggerLength int `default:"1024"`
	MaxResLoggerLength int
}

type Cfssl struct {
	ConfigFile string
	OCSPHost   string
	Config     *cfssl_config.Config
}

type Storage struct {
	Type string
}

func (h Storage) IsIpfs() bool {
	return h.Type == "ipfs"
}

func (h Storage) IsWeb3Storage() bool {
	return h.Type == "web3.storage"
}

type Ipfs struct {
	Host string
	Port int
}

type Web3Storage struct {
	Token      string
	EncryptKey string
	Timeout    int
	RetryCount int
}

type Contract struct {
	Address    string
	PrivateKey string
	RpcUrl     string
}
