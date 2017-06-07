package network

// Network : base DHCP Network object model
// See : https://h1infoblox.devops.int.ovp.bskyb.com/wapidoc/objects/network.html
type Network struct {
	Ref         string `json:"_ref"`
	Network     string `json:"network"`
	NetworkView string `json:"network_view"`
	Comment     string `json:"comment,omitempty"`
}
