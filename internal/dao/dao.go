package dao

import (
	"github.com/stc-community/fvm-dpki/internal/dao/certificate"

	"github.com/google/wire"
) // end

// RepoSet repo injection
var RepoSet = wire.NewSet(
	certificate.CertificateSet,
) // end

// Define repo type alias
type (
	CertificateRepo = certificate.CertificateRepo
) // end
