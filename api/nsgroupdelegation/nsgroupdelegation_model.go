package nsgroupdelegation

import "github.com/sky-uk/skyinfoblox/api/common"

const wapiVersion = "/wapi/v2.6.1"
const nsGroupDelegationEndpoint = "/nsgroup:delegation"

// NSGroupDelegation : Name Server Group Delegation object type
type NSGroupDelegation struct {
	Reference  string                  `json:"_ref,omitempty"`
	Comment    string                  `json:"comment,omitempty"`
	DelegateTo []common.ExternalServer `json:"delegate_to,omitempty"`
	Name       string                  `json:"name,omitempty"`
}
