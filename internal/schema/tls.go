package schema

import (
	"github.com/flowshield/cfssl/signer"
	"math/big"
)

type TlsShowResult struct {
	ID          string `json:"id"`
	Certificate string `json:"certificate"`
}

type LocalSignRequest struct {
	Hostname string                 `json:"hostname"`
	Hosts    []string               `json:"hosts"`
	Request  string                 `json:"certificate_request"`
	Subject  *signer.Subject        `json:"subject,omitempty"`
	Profile  string                 `json:"profile"`
	Label    string                 `json:"label"`
	Serial   *big.Int               `json:"serial,omitempty"`
	Bundle   bool                   `json:"bundle"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (a *LocalSignRequest) ToSignerSignRequest() signer.SignRequest {
	result := signer.SignRequest{
		Hosts:    a.Hosts,
		Request:  a.Request,
		Profile:  a.Profile,
		Label:    a.Label,
		Serial:   a.Serial,
		Metadata: a.Metadata,
	}
	if a.Hostname != "" {
		result.Hosts = signer.SplitHosts(a.Hostname)
	}
	if a.Subject != nil {
		result.Subject = a.Subject
	}
	return result
}

// RevokeParams
type RevokeParams struct {
	Serial  string `json:"serial" binding:"required"`
	AKI     string `json:"authority_key_id"`
	Reason  string `json:"reason"`
	Nonce   string `json:"nonce"`
	Sign    string `json:"sign"`
	AuthKey string `json:"auth_key" binding:"required"`
	Profile string `json:"profile"`
}
