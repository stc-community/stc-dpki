package certificate

import (
	"time"

	"github.com/stc-community/stc-fvm-dpki/internal/schema"
	"github.com/stc-community/stc-fvm-dpki/pkg/util/structure"
)

type SchemaCertificate schema.Certificate

func (a SchemaCertificate) ToCertificate() *Certificate {
	item := new(Certificate)
	structure.Copy(a, item)
	return item
}

type Certificate struct {
	SerialNumber           string                 `json:"serial_number"`
	AuthorityKeyIdentifier string                 `json:"authority_key_identifier"`
	CertPem                string                 `json:"cert_pem"`
	NotBefore              time.Time              `json:"not_before"`
	NotAfter               time.Time              `json:"not_after"`
	MetaData               map[string]interface{} `json:"meta_data"`
}

func (a Certificate) ToSchemaCertificate() *schema.Certificate {
	item := new(schema.Certificate)
	structure.Copy(a, item)
	return item
}

type Certificates []*Certificate

func (a Certificates) ToSchemaCertificates() []*schema.Certificate {
	list := make([]*schema.Certificate, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaCertificate()
	}
	return list
}
