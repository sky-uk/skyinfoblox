package zonedelegated

import (
	//"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"github.com/sky-uk/skyinfoblox/api/common"
)

func setupZoneDelegated(action string) *api.BaseAPI {

	delegateObject := common.ExternalServer{
		Address:"172.16.0.1",
		Name: "dns1.testdomain.com",
	}
	disableZone := false
	newZone := ZoneDelegated{
		Address: "10.0.0.1",
		Comment: "This is a comment",
		DelegateTo: []common.ExternalServer{delegateObject},
		DelegatedTTL: 1234,
		Disable: &disableZone,
		DnsFqdn: "example.com",
		View: "default",
		ZoneFormat: "FORWARD",
	}

	switch action {

	case "create":
		zoneAPI := NewCreateZoneDelegated(newZone)
		response := "/wapi/v2.6.1/zone_delegated/blablalba:example.com/default"
		zoneAPI.SetResponseObject(&response)
		return zoneAPI
	case "get":
		returnFields := []string{}
		zoneAPI := NewGetZoneDelegated("blablalba:example.com",returnFields)
		return  zoneAPI
	case "getWithFields":
		returnFields := []string{"delegate_to","fqdn"}
		zoneAPI := NewGetZoneDelegated("blablalba:example.com",returnFields)
		return  zoneAPI

	case "delete":
		zoneAPI := NewDeleteZoneDelegated("blablalba:example.com")
		return zoneAPI
	case "update":
		zoneAPI := NewUpdateZoneDelegated("blablalba:example.com",newZone)
		return  zoneAPI
	default:
		return nil
	}


}

func TestCreateZoneDelegatedMethod(t *testing.T){
	newZone := setupZoneDelegated("create")
	assert.Equal(t, http.MethodPost, newZone.Method() )
}

func TestCreateZoneDelegatedEndpoint(t *testing.T){
	newZone := setupZoneDelegated("create")
	assert.Equal(t, "/wapi/v2.3.1/zone_delegated", newZone.Endpoint() )
}

func TestGetZoneDelegatedMethod(t *testing.T) {
	newZone := setupZoneDelegated("get")
	assert.Equal(t, http.MethodGet, newZone.Method() )
}

func TestGetZoneDelegatedEndpoint(t *testing.T) {
	newZone := setupZoneDelegated("get")
	assert.Equal(t, "/wapi/v2.3.1/blablalba:example.com", newZone.Endpoint())
}

func TestGetZoneDelegatedEndpointWithFields(t *testing.T) {
	newZone := setupZoneDelegated("getWithFields")
	assert.Equal(t, "/wapi/v2.3.1/blablalba:example.com/?_return_fields=delegate_to,fqdn", newZone.Endpoint())
}


func TestDeleteZoneDelegatedMethod(t *testing.T) {
	newZone := setupZoneDelegated("delete")
	assert.Equal(t, http.MethodDelete, newZone.Method())
}

func TestDeleteZoneDelegatedEndpoint(t *testing.T) {
	newZone := setupZoneDelegated("delete")
	assert.Equal(t, "/wapi/v2.3.1/blablalba:example.com", newZone.Endpoint())
}

func TestUpdateZoneDelegatedMethod(t *testing.T) {
	newZone := setupZoneDelegated("update")
	assert.Equal(t, http.MethodPut, newZone.Method())
}

func TestUpdateZoneDelegatedEndpoing(t *testing.T) {
	newZone := setupZoneDelegated("update")
	assert.Equal(t, "/wapi/v2.3.1/blablalba:example.com", newZone.Endpoint())
}