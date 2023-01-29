package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/stc-community/fvm-dpki/internal/api"
)

var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

type Router struct {
	TlsAPI         *api.TlsAPI
	OcspAPI        *api.OcspAPI
	CertificateAPI *api.CertificateAPI
} // end

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) Prefixes() []string {
	return []string{
		"/api/",
	}
}

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		gCfssl := v1.Group("cfssl")
		{
			gCfssl.POST("info", a.TlsAPI.Info)
			gCfssl.POST("authsign", a.TlsAPI.AuthSign)
			gCfssl.POST("revoke", a.TlsAPI.Revoke)
		}

		gCert := v1.Group("certificate")
		{
			gCert.GET("get/storage/:id", a.CertificateAPI.GetS)
			gCert.GET("get/chain/:id", a.CertificateAPI.GetC)
			gCert.GET("revoke/:id", a.CertificateAPI.Revoke)
			gCert.GET("verify/:id", a.CertificateAPI.Verify)
		}
	} // v1 end
}
