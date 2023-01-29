package service

import (
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	TlsSet,
	OcspSet,
	CertificateSet,
) // end
