package records

// ARecord : ARecord data structure
type ARecord struct {
	Ref  string `json:"_ref"`
	IPv4 string `json:"ipv4addr"`
	Name string `json:"name"`
	View string `json:"view"`
}

// CNAMERecord : CNAMERecord data structure
type CNAMERecord struct {
	Ref       string `json:"_ref"`
	Canonical string `json:"canonical"`
	Name      string `json:"name"`
	View      string `json:"view"`
}

// TXTRecord : TXTRecord data structure
type TXTRecord struct {
	Ref  string `json:"_ref"`
	Name string `json:"name"`
	Text string `json:"text"`
	View string `json:"view"`
}

// SRVRecord : SRVRecord data structure
type SRVRecord struct {
	Ref      string `json:"_ref"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Target   string `json:"target"`
	View     string `json:"view"`
	Weight   int    `json:"weight"`
}
