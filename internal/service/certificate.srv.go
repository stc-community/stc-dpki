package service

import (
	"context"
	"encoding/hex"
	"github.com/flowshield/cfssl/helpers"
	"github.com/stc-community/stc-fvm-dpki/internal/config"
	"github.com/stc-community/stc-fvm-dpki/internal/dao"
	"github.com/stc-community/stc-fvm-dpki/internal/schema"
	"github.com/stc-community/stc-fvm-dpki/pkg/errors"
	"github.com/stc-community/stc-fvm-dpki/pkg/logger"
	"github.com/google/wire"
	"os"
)

var CertificateSet = wire.NewSet(wire.Struct(new(CertificateSrv), "*"))

type CertificateSrv struct {
	CertificateRepo *dao.CertificateRepo
}

func (a *CertificateSrv) GetS(ctx context.Context, id string) (*schema.Certificate, error) {
	item, err := a.CertificateRepo.GetS(ctx, id)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *CertificateSrv) GetC(ctx context.Context, sn string) (interface{}, error) {
	item, err := a.CertificateRepo.GetBlockChain(ctx, sn)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}
	return item, nil
}

func (a *CertificateSrv) Revoke(ctx context.Context, sn string) (interface{}, error) {
	result, err := a.CertificateRepo.Revoke(ctx, sn)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a *CertificateSrv) Verify(ctx context.Context, sn string) (interface{}, error) {
	item, err := a.CertificateRepo.Verify(ctx, sn)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (a *CertificateSrv) CreateRootCa(ctx context.Context) error {
	cfg := config.C.TLS
	certPEM, err := os.ReadFile(cfg.CertFile)
	if err != nil {
		return err
	}
	certs, err := helpers.ParseCertificatesPEM(certPEM)
	for _, cert := range certs {
		certificate := schema.Certificate{
			SerialNumber:           cert.SerialNumber.String(),
			AuthorityKeyIdentifier: hex.EncodeToString(cert.AuthorityKeyId),
			CertPem:                string(helpers.EncodeCertificatePEM(cert)),
			NotBefore:              cert.NotBefore,
			NotAfter:               cert.NotAfter,
		}
		res, err := a.CertificateRepo.PutS(ctx, certificate)
		if err != nil {
			return err
		}
		item, err := a.CertificateRepo.GetBlockChain(ctx, cert.SerialNumber.String())
		if err != nil {
			return err
		}
		if item != nil && (certificate.Hash() != item.CidDocHash) {
			tx, err := a.CertificateRepo.PutBlockChain(
				ctx,
				cert.SerialNumber.String(),
				hex.EncodeToString(cert.SubjectKeyId),
				hex.EncodeToString(cert.AuthorityKeyId),
				res.ID,
				certificate.Hash(),
			)
			if err != nil {
				return err
			}
			logger.Infof("Rootca Save tx:%s; sn:%s", tx.Hash(), cert.SerialNumber.String())
		}
	}
	return nil
}
