package zonestub

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setupZoneStub(action string) *api.BaseAPI {
	zoneDisable := true
	forwardingDisable := true
	lockedZone := false
	stubfrom := common.ExternalServer{
		Address: "1.2.3.4",
		Name:    "externalname.testdomain.com",
	}

	zoneStub := ZoneStub{
		Ref:               "zone_stub/blablalba:example.com/default",
		Comment:           "this is a comment",
		Disable:           &zoneDisable,
		DisableForwarding: &forwardingDisable,
		ExternalNSGroup:   "NS-Group-1",
		FQDN:              "stubtest.com",
		Locked:            &lockedZone,
		NsGroup:           "NS-Group-2",
		Prefix:            "prefix",
		StubFrom:          []common.ExternalServer{stubfrom},
		View:              "default",
		ZoneFormat:        "FORWARD",
	}

	switch action {

	case "create":
		zoneAPI := NewCreate(zoneStub)
		response := "zone_stub/blablalba:example.com/default"
		zoneAPI.SetResponseObject(&response)
		return zoneAPI
	case "get":
		returnFields := []string{}
		zoneAPI := NewGet("zone_stub/blablalba:example.com", returnFields)
		return zoneAPI
	case "getall":
		returnFields := []string{}
		zoneAPI := NewGetAll(returnFields)
		return zoneAPI
	case "getallWithFields":
		returnFields := []string{"stub_from", "fqdn"}
		zoneAPI := NewGetAll(returnFields)
		return zoneAPI

	case "getWithFields":
		returnFields := []string{"stub_from", "fqdn"}
		zoneAPI := NewGet("zone_stub/blablalba:example.com", returnFields)
		return zoneAPI
	case "delete":
		zoneAPI := NewDelete("zone_stub/blablalba:example.com")
		return zoneAPI
	case "update":
		zoneAPI := NewUpdate(zoneStub)
		return zoneAPI
	default:
		return nil
	}
}

func TestZoneStubCreateMethod(t *testing.T) {
	newStubZone := setupZoneStub("create")
	assert.Equal(t, http.MethodPost, newStubZone.Method())
}

func TestZoneStubCreateEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("create")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub", newStubZone.Endpoint())
}

func TestZoneStubCreateResponse(t *testing.T) {
	newStubZone := setupZoneStub("create")
	assert.Equal(t, "zone_stub/blablalba:example.com/default", *newStubZone.ResponseObject().(*string))
}

func TestZoneStubGetMethod(t *testing.T) {
	newStubZone := setupZoneStub("get")
	assert.Equal(t, http.MethodGet, newStubZone.Method())
}

func TestZoneStubGetEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("get")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub/blablalba:example.com", newStubZone.Endpoint())
}
func TestZoneStubGetAllMethod(t *testing.T) {
	newStubZone := setupZoneStub("getall")
	assert.Equal(t, http.MethodGet, newStubZone.Method())
}

func TestZoneStubGetAllEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("getall")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub", newStubZone.Endpoint())
}

func TestZoneStubGetAllWithFieldsMethod(t *testing.T) {
	newStubZone := setupZoneStub("getallWithFields")
	assert.Equal(t, http.MethodGet, newStubZone.Method())
}

func TestZoneStubGetAllWithFieldsEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("getallWithFields")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub?_return_fields=stub_from,fqdn", newStubZone.Endpoint())
}

func TestZoneStubGetEndpointWithFields(t *testing.T) {
	newStubZone := setupZoneStub("getWithFields")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub/blablalba:example.com?_return_fields=stub_from,fqdn", newStubZone.Endpoint())
}

func TestZoneStubUpdateMethod(t *testing.T) {
	newStubZone := setupZoneStub("update")
	assert.Equal(t, http.MethodPut, newStubZone.Method())
}

func TestZoneStubUpdateEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("update")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub/blablalba:example.com/default", newStubZone.Endpoint())
}

func TestZoneStubDeleteMethod(t *testing.T) {
	newStubZone := setupZoneStub("delete")
	assert.Equal(t, http.MethodDelete, newStubZone.Method())
}

func TestZoneStubDeleteEndpoint(t *testing.T) {
	newStubZone := setupZoneStub("delete")
	assert.Equal(t, "/wapi/v2.6.1/zone_stub/blablalba:example.com", newStubZone.Endpoint())
}
