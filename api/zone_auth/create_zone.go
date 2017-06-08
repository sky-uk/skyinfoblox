package zone_auth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateZoneAPI : Create zone API
type CreateZoneAPI struct {
	*api.BaseAPI
}

// NewCreate : Create a new zone
func NewCreate(newZone DNSZone) *CreateZoneAPI {
	this := new(CreateZoneAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/zone_auth?_return_fields=fqdn", newZone, new(DNSZone))
	return this
}

// GetResponse : get response object from created zone
func (cza CreateZoneAPI) GetResponse() string {
	return cza.ResponseObject().(string)
}
