package zonedelegated

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

var endPoint string

// NewCreateZoneDelegated - Create a new Zone
func NewCreateZoneDelegated(newZoneDelegated ZoneDelegated) *api.BaseAPI {
	endPoint = "/wapi/v2.3.1/zone_delegated"
	createZoneAPI := api.NewBaseAPI(http.MethodPost, endPoint, newZoneDelegated, new(string))
	return createZoneAPI
}

// NewGetZoneDelegated - Read an existing zone
func NewGetZoneDelegated(ref string, returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("/wapi/v2.3.1/%s/?_return_fields=%s", ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("/wapi/v2.3.1/%s", ref)
	}
	getUserAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new(ZoneDelegated))
	return getUserAPI

}

// NewUpdateZoneDelegated - Update a zone
func NewUpdateZoneDelegated(ref string, updateZoneDelegated ZoneDelegated) *api.BaseAPI {
	endPoint := fmt.Sprintf("/wapi/v2.3.1/%s", ref)
	updateUserAPI := api.NewBaseAPI(http.MethodPut, endPoint, updateZoneDelegated, new(string))
	return updateUserAPI

}

// NewDeleteZoneDelegated - Delete a zone
func NewDeleteZoneDelegated(ref string) *api.BaseAPI {
	endPoint := fmt.Sprintf("/wapi/v2.3.1/%s", ref)
	deleteUserAPI := api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(string))
	return deleteUserAPI

}
