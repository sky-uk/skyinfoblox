package zone_auth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetAllZones : all zones struct
type GetAllZones struct {
	*api.BaseAPI
}

// NewGetAll : returns an object containing all zones.
func NewGetAll() *GetAllZones {
	this := new(GetAllZones)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/zone_auth?_return_fields=fqdn", nil, new(DNSZoneReferences))
	return this
}

// GetResponse : returns the response object of GetAllZones
func (gaz GetAllZones) GetResponse() *DNSZoneReferences {
	return gaz.ResponseObject().(*DNSZoneReferences)
}
