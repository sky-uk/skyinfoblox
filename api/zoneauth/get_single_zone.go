package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetSingleZoneAuthAPI type
type GetSingleZoneAuthAPI struct {
	*api.BaseAPI
}

// NewGetSingleZone : returns a zone's details.
func NewGetSingleZone(ref string) *GetSingleZoneAuthAPI {
	this := new(GetSingleZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+ref+"?_return_fields=comment,fqdn,view", nil, new(DNSZone))
	return this
}

// GetResponse : returns response obeject from GetSingleZone
func (gsz GetSingleZoneAuthAPI) GetResponse() *DNSZone {
	return gsz.ResponseObject().(*DNSZone)
}
