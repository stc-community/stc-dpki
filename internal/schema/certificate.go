package schema

import (
	"github.com/stc-community/fvm-dpki/pkg/util"
	"time"

	"github.com/stc-community/fvm-dpki/pkg/util/json"
)

type Certificate struct {
	SerialNumber           string                 `json:"serial_number"`
	AuthorityKeyIdentifier string                 `json:"authority_key_identifier"`
	CertPem                string                 `json:"cert_pem"`
	NotBefore              time.Time              `json:"not_before"`
	NotAfter               time.Time              `json:"not_after"`
	MetaData               map[string]interface{} `json:"meta_data"`
}

func (a *Certificate) String() string {
	return json.MarshalToString(a)
}

func (a *Certificate) Hash() string {
	return util.Md5(json.MarshalToString(a))
}

// CertificateRevoke 证书状态信息
type CertificateRevoke struct {
	SerialNumber string    `json:"serial_number"`
	RevokeAt     time.Time `json:"revoke_at"`
}

func (a *CertificateRevoke) String() string {
	return json.MarshalToString(a)
}
