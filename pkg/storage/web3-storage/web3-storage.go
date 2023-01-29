package web3_storage

import (
	"context"
	"fmt"
	"github.com/stc-community/stc-fvm-dpki/pkg/errors"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/stc-community/stc-fvm-dpki/internal/config"

	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	"github.com/wumansgy/goEncrypt"
)

type Web3Storage struct {
	db         w3s.Client
	httpClient *http.Client
	timeout    time.Duration
	filename   string
	key        []byte
}

func New() (*Web3Storage, error) {
	cfg := config.C.Web3Storage
	db, err := w3s.NewClient(w3s.WithToken(cfg.Token))
	if err != nil {
		return nil, err
	}
	return &Web3Storage{
		db:       db,
		timeout:  time.Duration(cfg.Timeout) * time.Second,
		filename: "filename",
		httpClient: &http.Client{
			Timeout: time.Duration(cfg.Timeout) * time.Second,
		},
		key: []byte(cfg.EncryptKey),
	}, nil
}

func (a *Web3Storage) Put(ctx context.Context, data []byte) (cid string, err error) {
	file, err := a.dataToFile(data, a.filename, a.key)
	defer os.Remove(file.Name())
	if err != nil {
		return
	}
	cidObj, err := a.db.Put(ctx, file)
	if err != nil {
		return
	}
	return cidObj.String(), nil
}

func (a *Web3Storage) Get(ctx context.Context, cidStr string) ([]byte, error) {
	var result []byte
	var err error
	result, err = a.GetByW3sLink(cidStr, a.filename, a.key)
	if err != nil {
		log.Println("Failed to obtain from w3s link, err: ", err.Error())
		ctx, cancel := context.WithTimeout(ctx, a.timeout*time.Second)
		defer cancel()
		result, err = a.GetByW3sClient(ctx, cidStr, a.key)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	return result, nil
}

func (a *Web3Storage) GetByW3sLink(cid string, filename string, key []byte) ([]byte, error) {
	url := fmt.Sprintf("https://%s.ipfs.w3s.link/%s", cid, filename)
	resp, err := a.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request error, url:%s, code:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data, err := goEncrypt.DesCbcDecrypt(body, key[:], nil) // 解密得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Web3Storage) GetByW3sClient(ctx context.Context, cidStr string, key []byte) (data []byte, err error) {
	cidObj, err := cid.Decode(cidStr)
	if err != nil {
		return nil, err
	}
	res, err := a.db.Get(ctx, cidObj)
	if err != nil {
		return nil, err
	}
	_, fsys, err := res.Files()
	if err != nil {
		return nil, err
	}
	err = fs.WalkDir(fsys, "/", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			file, err := fsys.Open(path)
			if err != nil {
				return err
			}
			data, err = ioutil.ReadAll(file)
			if err != nil {
				return err
			}
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	data, err = goEncrypt.DesCbcDecrypt(data, key[:], nil) // 解密得到密文,可以自己传入初始化向量,如果不传就使用默认的初始化向量,8字节
	if err != nil {
		return nil, err
	}
	return
}

func (a *Web3Storage) Close() error {
	return nil
}

func (a *Web3Storage) dataToFile(data []byte, filename string, key []byte) (file *os.File, err error) {
	file, err = os.Create(filename)
	if err != nil {
		return
	}
	// 对数据进行加密
	cryptText, err := goEncrypt.DesCbcEncrypt(data, key[:], nil)
	if err != nil {
		return
	}
	err = os.WriteFile(file.Name(), cryptText, 0o644)
	if err != nil {
		return
	}
	return
}
