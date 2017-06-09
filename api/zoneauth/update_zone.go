package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateZoneAuthAPI : Update zone API
type UpdateZoneAuthAPI struct {
	*api.BaseAPI
}

// NewUpdate : Update zone
func NewUpdate(updateDNSZone DNSZone) *UpdateZoneAuthAPI {
	this := new(UpdateZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/wapi/v2.3.1/"+updateDNSZone.Reference+"?_return_fields=fqdn,view,comment", updateDNSZone, new(DNSZone))
	return this
}

// GetResponse : returns the response from UpdateZoneAPI
func (UpdateZoneAPI UpdateZoneAuthAPI) GetResponse() *DNSZone {
	return UpdateZoneAPI.ResponseObject().(*DNSZone)
}
