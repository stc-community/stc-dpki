package ipfs

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/stc-community/stc-fvm-dpki/internal/config"
	"github.com/stc-community/stc-fvm-dpki/pkg/errors"
	shell "github.com/ipfs/go-ipfs-api"
)

type Ipfs struct {
	sh *shell.Shell
}

func New() (*Ipfs, error) {
	cfg := config.C.Ipfs
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	sh := shell.NewShell(addr)
	if !sh.IsUp() {
		return nil, errors.New("ipfs is not up")
	}
	return &Ipfs{
		sh: sh,
	}, nil
}

func (a *Ipfs) Put(ctx context.Context, data []byte) (string, error) {
	hash, err := a.sh.Add(bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (a *Ipfs) Get(ctx context.Context, hash string) ([]byte, error) {
	read, err := a.sh.Cat(hash)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(read)
	return body, nil
}

func (a *Ipfs) Close() error {
	return nil
}
