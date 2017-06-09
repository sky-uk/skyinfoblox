package zoneauth

// DNSZone : Contains zone configuration. Reference is used during updates and when retriving the zone.
type DNSZone struct {
	Reference string `json:"_ref,omitempty"`
	FQDN      string `json:"fqdn,omitempty"`
	View      string `json:"view,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// DNSZoneReference : A zone, it's reference and associated FQDN used for finding a zone when getting a list of all zones
type DNSZoneReference struct {
	Reference string `json:"_ref"`
	FQDN      string `json:"fqdn"`
}

// DNSZoneReferences : A list of zone references
type DNSZoneReferences []DNSZoneReference
