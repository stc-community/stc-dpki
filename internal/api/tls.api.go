package api

import (
	"github.com/flowshield/cfssl/api"
	"github.com/flowshield/cfssl/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/stc-community/fvm-dpki/internal/ginx"
	"github.com/stc-community/fvm-dpki/internal/schema"
	"github.com/stc-community/fvm-dpki/internal/service"
)

var TlsSet = wire.NewSet(wire.Struct(new(TlsAPI), "*"))

type TlsAPI struct {
	TlsSrv *service.TlsSrv
}

// 利用cfssl内部handle
func (a *TlsAPI) Info(c *gin.Context) {
	a.TlsSrv.Info(c)
}

func (a *TlsAPI) AuthSign(c *gin.Context) {
	ctx := c.Request.Context()
	var params auth.AuthenticatedRequest
	if err := ginx.ParseJSON(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	result, err := a.TlsSrv.AuthSign(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	if err := api.SendResponse(c.Writer, result); err != nil {
		ginx.ResError(c, err)
		return
	}
}

// 利用cfssl内部handle
func (a *TlsAPI) Revoke(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.RevokeParams
	if err := ginx.ParseJSON(c, &params); err != nil {
		ginx.ResError(c, err)
		return
	}
	err := a.TlsSrv.Revoke(ctx, params)
	if err != nil {
		ginx.ResError(c, err)
		return
	}
	api.SendResponse(c.Writer, map[string]string{})
}
