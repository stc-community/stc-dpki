package certificate

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/wire"
	"github.com/stc-community/fvm-dpki/internal/schema"
	"github.com/stc-community/fvm-dpki/pkg/contract"
	"github.com/stc-community/fvm-dpki/pkg/errors"
	"github.com/stc-community/fvm-dpki/pkg/storage"
	"github.com/stc-community/fvm-dpki/pkg/util/json"
)

var CertificateSet = wire.NewSet(wire.Struct(new(CertificateRepo), "DB", "Eth"))

type CertificateRepo struct {
	DB  storage.ExecCloser
	Eth *contract.EthClient
}

func (a *CertificateRepo) GetS(ctx context.Context, id string) (*schema.Certificate, error) {
	errCount := 0
retry:
	str, err := a.DB.Get(ctx, id)
	if err != nil {
		if errCount < 3 {
			errCount++
			goto retry
		}
		return nil, err
	}
	var result Certificate
	err = json.Unmarshal(str, &result)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return result.ToSchemaCertificate(), err
}

func (a *CertificateRepo) PutS(ctx context.Context, item schema.Certificate) (*schema.IDResult, error) {
	itemByte, err := json.Marshal(item)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	errCount := 0
retry:
	id, err := a.DB.Put(ctx, itemByte)
	if err != nil {
		if errCount < 3 {
			errCount++
			goto retry
		}
		return nil, errors.WithStack(err)
	}
	return schema.NewIDResult(id), nil
}

func (a *CertificateRepo) PutBlockChain(ctx context.Context, sn, ski, aki, cid, cidDocHash string) (*types.Transaction, error) {
	v, err := a.Eth.Instance.Save(a.Eth.Auth, sn, ski, aki, cid, cidDocHash)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return v, err
}

func (a *CertificateRepo) GetBlockChain(ctx context.Context, sn string) (*contract.CertificateCert, error) {
	v, err := a.Eth.Instance.Get(&bind.CallOpts{Pending: true, From: a.Eth.Auth.From}, sn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &v, err
}

func (a *CertificateRepo) Verify(ctx context.Context, sn string) (*contract.CertificateCert, error) {
	v, err := a.Eth.Instance.Verify(&bind.CallOpts{Pending: true, From: a.Eth.Auth.From}, sn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &v, err
}

func (a *CertificateRepo) Revoke(ctx context.Context, sn string) (*types.Transaction, error) {
	v, err := a.Eth.Instance.Revoke(a.Eth.Auth, sn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return v, err
}
