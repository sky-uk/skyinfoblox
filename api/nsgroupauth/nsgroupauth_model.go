package nsgroupauth

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupEndpoint = "/nsgroup"

// NSGroupAuth : Name Server Group Authoritative object type
type NSGroupAuth struct {
	Reference           string                  `json:"_ref,omitempty"`
	Comment             string                  `json:"comment,omitempty"`
	ExternalPrimaries   []common.ExternalServer `json:"external_primaries,omitempty"`
	ExternalSecondaries []common.ExternalServer `json:"external_secondaries,omitempty"`
	GridPrimary         []common.MemberServer   `json:"grid_primary,omitempty"`
	GridSecondaries     []common.MemberServer   `json:"grid_secondaries,omitempty"`
	GridDefault         *bool                   `json:"is_grid_default,omitempty"`
	Name                string                  `json:"name,omitempty"`
	UseExternalPrimary  *bool                   `json:"use_external_primary,omitempty"`
}
