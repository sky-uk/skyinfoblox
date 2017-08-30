package mxrecord

//WapiVersion : WAPI version related with this data model
const WapiVersion = "/wapi/v2.6.1"

// Endpoint - resource WAPI endpoint
const MXRecordEndpoint = "record:mx"

type MxRecord struct {
	Ref               string `json:"_ref,omitempty"`
	Comment           string `json:"comment,omitempty"`
	DDNSPrincipal     string `json:"ddns_principal,omitempty"`
	DDNSProtected     bool   `json:"ddns_protected"`
	Disable           bool   `json:"disable"`
	ForbidReclamation bool   `json:"forbid_reclamation"`
	MailExchanger     string `json:"mail_exchanger"`
	Name              string `json:"name"`
	Preference        uint   `json:"preference"`
	TTL               uint   `json:"ttl,omitempty"`
	UseTTL            bool   `json:"use_ttl"`
	View              string `json:"view,omitempty"`
	Zone              string `json:"zone,omitempty"`
}
