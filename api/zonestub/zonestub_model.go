package zonestub

const WapiVersion = "/wapi/v2.6.1"
const Endpoint = "/zone_forward"

type ZoneStub struct {
	Ref               string `json:"_ref,omitempty"`
	Comment           string `json:"comment,omitempty"`
	Disable           *bool  `json:"disable,omitempty"`
	DisableForwarding *bool  `json:"disable_forwarding,omitempty"`
	ExternalNSGroup   string `json:"external_ns_group,omitempty"`
	FQDN              string `json:"fqdn,omitempty"`
	Locked            *bool  `json:"locked,omitempty"`
	MaskPrefix        string `json:"mask_prefix,omitempty"`
	NsGroup           string `json:"ns_group,omitempty"`
}
