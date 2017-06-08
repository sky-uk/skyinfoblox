package zoneauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// DeleteZoneAPI : Zone API for deleting a zone
type DeleteZoneAPI struct {
	*api.BaseAPI
}

// NewDelete : delete a resource by it's reference - this function can probably be common to all Infoblox resource types.
func NewDelete(ref string) *DeleteZoneAPI {
	this := new(DeleteZoneAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/wapi/v2.3.1/"+ref, nil, new(string))
	return this
}
