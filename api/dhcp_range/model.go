package dhcprange

type DHCPRange struct {
	Ref               string          `json:"_ref"`
	Start             string          `json:"start_addr"`
	End               string          `json:"end_addr"`
	Network           string          `json:"network"`
	NetworkView       string          `json:"network_view"`
	Restart           bool            `json:"restart_if_needed"`
	ServerAssociation string          `json:"server_association_type"`
	Member            DHCPRangeMember `json:"member"`
}

type DHCPRangeMember struct {
	InternalType string `json:"_struct"`
	Address      string `json:"ipv4addr"`
	Name         string `json:"name"`
}
