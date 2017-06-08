package zone_auth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateZoneAPI : Update zone API
type UpdateZoneAPI struct {
	*api.BaseAPI
}

// NewUpdate : Update zone
func NewUpdate(updateDNSZone DNSZone, zoneReference string) *UpdateZoneAPI {
	this := new(UpdateZoneAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/wapi/v2.3.1/"+zoneReference+"?_return_fields=fqdn,view,comment", updateDNSZone, new(DNSZone))
	return this
}

// GetResponse : returns the response from UpdateZoneAPI
func (UpdateZoneAPI UpdateZoneAPI) GetResponse() string {
	return UpdateZoneAPI.ResponseObject().(string)
}
