package zonedelegated

import (
	"github.com/sky-uk/skyinfoblox/api/common"
)

type ZoneDelegated struct {
	Ref                    string                  `json:"_ref,omitempty"`
	Address                string                  `json:"address,omitempty"`
	Comment                string                  `json:"comment,omitempty"`
	DelegateTo             []common.ExternalServer `json:"delegate_to"`
	DelegatedTTL           uint                    `json:"delegated_ttl,omitempty"`
	Disable                *bool                   `json:"disable,omitempty"`
	//DisplayDomain          string                  `json:"display_domain,omitempty"`
	DnsFqdn                string                  `json:"dns_fqdn,omitempty"`
	EnableRFC2317Exclusion *bool                   `json:"enable_rfc2317_exclusion,omitempty"`
	Fqdn                   string                  `json:"fqdn"`
	Locked                 *bool                   `json:"locked,omitempty"`
	NameServerGroup        string                  `json:"ns_group,omitempty"`
	Prefix                 string                  `json:"prefix,omitempty"`
	UseDelegatedTTL        *bool                   `json:"use_delegated_ttl,omitempty"`
	View                   string                  `json:"view"`
	ZoneFormat             string                  `json:"zone_format,omitempty"`
}
