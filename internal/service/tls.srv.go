package service

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/flowshield/cfssl/auth"
	cfssl_config "github.com/flowshield/cfssl/config"
	"github.com/flowshield/cfssl/helpers"
	"github.com/flowshield/cfssl/signer"
	"github.com/stc-community/stc-dpki/internal/dao"
	"github.com/stc-community/stc-dpki/internal/initx"
	"github.com/stc-community/stc-dpki/internal/schema"
	"github.com/stc-community/stc-dpki/pkg/attrmgr"
	"github.com/stc-community/stc-dpki/pkg/errors"
	"github.com/stc-community/stc-dpki/pkg/logger"
	"github.com/stc-community/stc-dpki/pkg/util/json"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var TlsSet = wire.NewSet(wire.Struct(new(TlsSrv), "*"))

type TlsSrv struct {
	CertificateRepo *dao.CertificateRepo
	CfsslHandler    *initx.CfsslHandler
}

func (a *TlsSrv) Info(c *gin.Context) {
	a.CfsslHandler.InfoHandler.ServeHTTP(c.Writer, c.Request)
}

func (a *TlsSrv) Revoke(ctx context.Context, params schema.RevokeParams) error {
	// 数据库查询
	//_, err := a.CertificateRepo.GetC(ctx, schema.SnCidKey(params.Serial))
	//if err != nil {
	//	return err
	//}
	//if params.Profile == "" {
	//	return errors.New("profile 未指定")
	//}
	//if authKey, ok := config.C.Cfssl.Config.AuthKeys[params.Profile]; ok {
	//	if authKey.Key != params.AuthKey {
	//		return errors.New("非法操作")
	//	}
	//}
	//data := schema.CertificateRevoke{
	//	SerialNumber: params.Serial,
	//	RevokeAt:     time.Now(),
	//}
	//err = a.CertificateRepo.PutC(ctx, schema.SnRevokeKey(params.Serial), data.String())
	//if err != nil {
	//	return err
	//}
	return nil
}

func (a *TlsSrv) AuthSign(ctx context.Context, params auth.AuthenticatedRequest) (*schema.TlsShowResult, error) {
	var lsr schema.LocalSignRequest
	err := json.Unmarshal(params.Request, &lsr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Sanity checks to ensure that we have a valid policy. This
	// should have been checked in NewAuthHandler.
	policy := a.CfsslHandler.LocalSigner.Policy()
	if policy == nil {
		return nil, errors.New("invalid policy")
	}

	profile, err := signer.Profile(a.CfsslHandler.LocalSigner, lsr.Profile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if profile.Provider == nil {
		return nil, errors.New("no authentication provider")
	}

	validAuth := false
	if profile.Provider.Verify(&params) {
		validAuth = true
	} else if profile.PrevProvider != nil && profile.PrevProvider.Verify(&params) {
		validAuth = true
	}
	if !validAuth {
		return nil, errors.New("invalid token")
	}

	signReq := lsr.ToSignerSignRequest()
	if signReq.Request == "" {
		return nil, errors.New("missing parameter 'certificate_rest'")
	}
	err = a.genExpiryByCsr(&signReq, profile)
	if err != nil {
		return nil, err
	}
	cert, err := a.CfsslHandler.LocalSigner.Sign(signReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// 存储
	// save storage
	item, err := a.saveCertificate(ctx, cert, signReq.Metadata)
	if err != nil {
		return nil, err
	}
	logger.Infof("cid:%s", item.ID)
	return &schema.TlsShowResult{
		ID:          item.ID,
		Certificate: string(cert),
	}, nil
}

// 存储
func (a *TlsSrv) saveCertificate(ctx context.Context, certPem []byte, metaData map[string]interface{}) (*schema.IDResult, error) {
	cert, err := helpers.ParseCertificatePEM(certPem)
	if err != nil {
		return nil, err
	}
	sn := cert.SerialNumber.String()
	certificate := schema.Certificate{
		SerialNumber:           sn,
		AuthorityKeyIdentifier: hex.EncodeToString(cert.AuthorityKeyId),
		CertPem:                string(certPem),
		NotBefore:              cert.NotBefore,
		NotAfter:               cert.NotAfter,
		MetaData:               metaData,
	}
	idResult, err := a.CertificateRepo.PutS(ctx, certificate)
	if err != nil {
		return nil, err
	}
	tran, err := a.CertificateRepo.PutBlockChain(
		ctx,
		certificate.SerialNumber,
		hex.EncodeToString(cert.SubjectKeyId),
		hex.EncodeToString(cert.AuthorityKeyId),
		idResult.ID,
		certificate.Hash())
	if err != nil {
		return nil, err
	}
	logger.Infof("save to chain:%s", tran.Hash())
	return schema.NewIDResult(idResult.ID), nil
}

// 添加过期时间
func (a *TlsSrv) genExpiryByCsr(sr *signer.SignRequest, profile *cfssl_config.SigningProfile) error {
	csr, err := helpers.ParseCSRPEM([]byte(sr.Request))
	if err != nil {
		return errors.WithStack(err)
	}
	if v := attrmgr.GetExpiryValue(csr); v > 0 {
		var backdate time.Duration
		if backdate = profile.Backdate; backdate == 0 {
			backdate = -5 * time.Minute
		} else {
			backdate = -1 * profile.Backdate
		}
		notBefore := time.Now().Round(time.Minute).Add(backdate)
		notBefore = notBefore.UTC()
		notAfter := notBefore.Add(v)
		notAfter = notAfter.UTC()
		sr.NotBefore = notBefore
		sr.NotAfter = notAfter
	}
	return nil
}
