package zone_auth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetSingleZone type
type GetSingleZone struct {
	*api.BaseAPI
}

// NewGetSingleZone : returns a zone's details.
func NewGetSingleZone(ref string) *GetSingleZone {
	this := new(GetSingleZone)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+ref+"?_return_fields=comment,fqdn,view", nil, new(DNSZone))
	return this
}

// GetResponse : returns response obeject from GetSingleZone
func (gsz GetSingleZone) GetResponse() *DNSZone {
	return gsz.ResponseObject().(*DNSZone)
}
