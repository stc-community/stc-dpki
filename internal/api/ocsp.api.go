package api

import (
	"net/http"

	"github.com/stc-community/stc-fvm-dpki/internal/service"
	"github.com/google/wire"
	stdocsp "golang.org/x/crypto/ocsp"
)

var OcspSet = wire.NewSet(wire.Struct(new(OcspAPI), "*"))

type OcspAPI struct {
	OcspSrv *service.OcspSrv
}

// 利用cfssl内部handle
func (a *OcspAPI) Response(req *stdocsp.Request) ([]byte, http.Header, error) {
	return a.OcspSrv.Query(req)
}
