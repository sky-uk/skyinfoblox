package zone_auth


// DNSZone : Contains zone configuration.
type DNSZone struct {
	Reference string `json:"_ref,omitempty"`
	FQDN      string `json:"fqdn,omitempty"`
	View      string `json:"view,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// DNSZoneReference : A zone, it's reference and associated FQDN.
type DNSZoneReference struct {
	Reference string `json:"_ref"`
	FQDN      string `json:"fqdn"`
}

// DNSZoneReferences : A list of zone references
type DNSZoneReferences []DNSZoneReference
