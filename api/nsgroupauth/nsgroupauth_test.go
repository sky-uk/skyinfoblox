package nsgroupauth

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createNSGroupAuthAPI, getNSGroupAuthAPI, getAllNSGroupAuthAPI, updateNSGroupAuthAPI, deleteNSGroupAuthAPI *api.BaseAPI
var nsGroupAuthObject NSGroupAuth
var nsGroupAuthObjectList []NSGroupAuth
var returnFields []string

func setupNSGroupAuthTest(testType string) {

	gridDefault := true
	useExternalPrimary := true
	nsGroupAuthObject = NSGroupAuth{
		Reference:          "nsgroup/ZG5zOm5zX2dyb2VwJAByaW1hcnlfWm9uZV9YRlI:TEST_NS_GROUP_AUTH",
		Comment:            "Test NS Group Auth",
		GridDefault:        &gridDefault,
		Name:               "test-ns-group-auth",
		UseExternalPrimary: &useExternalPrimary,
	}
	nsGroupAuthObjectList = make([]NSGroupAuth, 0)
	nsGroupAuthObjectList = append(nsGroupAuthObjectList, nsGroupAuthObject)
	returnFields = []string{"comment", "external_primaries", "external_secondaries", "grid_primary", "grid_secondaries", "is_grid_default", "name", "use_external_primary"}

	switch testType {
	case "create":
		createNSGroupAuthAPI = NewCreate(nsGroupAuthObject)
		createNSGroupAuthAPI.SetResponseObject(&nsGroupAuthObject.Reference)
	case "get":
		getNSGroupAuthAPI = NewGet(nsGroupAuthObject.Reference, returnFields)
		getNSGroupAuthAPI.SetResponseObject(&nsGroupAuthObject)
	case "getall":
		getAllNSGroupAuthAPI = NewGetAll()
		getAllNSGroupAuthAPI.SetResponseObject(&nsGroupAuthObjectList)
	case "update":
		updateNSGroupAuthAPI = NewUpdate(nsGroupAuthObject, returnFields)
		updateNSGroupAuthAPI.SetResponseObject(&nsGroupAuthObject)
	case "delete":
		deleteNSGroupAuthAPI = NewDelete(nsGroupAuthObject.Reference)
	case "default":
		fmt.Println("Option not implemented")
	}
}

func TestNameServerGroupAuthNewCreateMethod(t *testing.T) {
	setupNSGroupAuthTest("create")
	assert.Equal(t, http.MethodPost, createNSGroupAuthAPI.Method())
}

func TestNameServerGroupAuthNewCreateEndpoint(t *testing.T) {
	setupNSGroupAuthTest("create")
	assert.Equal(t, wapiVersion+nsGroupEndpoint, createNSGroupAuthAPI.Endpoint())
}

func TestNameServerGroupAuthNewCreateResponse(t *testing.T) {
	setupNSGroupAuthTest("create")
	response := *createNSGroupAuthAPI.ResponseObject().(*string)
	assert.Equal(t, nsGroupAuthObject.Reference, response)
}

func TestNameServerGroupAuthNewGetMethod(t *testing.T) {
	setupNSGroupAuthTest("get")
	assert.Equal(t, http.MethodGet, getNSGroupAuthAPI.Method())
}

func TestNameServerGroupAuthNewGetEndpoint(t *testing.T) {
	setupNSGroupAuthTest("get")
	assert.Equal(t, wapiVersion+"/"+nsGroupAuthObject.Reference+"?_return_fields="+strings.Join(returnFields, ","), getNSGroupAuthAPI.Endpoint())
}

func TestNameServerGroupAuthNewGetResponse(t *testing.T) {
	setupNSGroupAuthTest("get")
	response := getNSGroupAuthAPI.ResponseObject().(*NSGroupAuth)

	assert.Equal(t, "test-ns-group-auth", response.Name)
	assert.Equal(t, "Test NS Group Auth", response.Comment)
	assert.Equal(t, true, *response.GridDefault)
	assert.Equal(t, true, *response.UseExternalPrimary)
}

func TestNameServerGroupAuthNewGetAllMethod(t *testing.T) {
	setupNSGroupAuthTest("getall")
	assert.Equal(t, http.MethodGet, getAllNSGroupAuthAPI.Method())
}

func TestNameServerGroupAuthNewGetAllEndpoint(t *testing.T) {
	setupNSGroupAuthTest("getall")
	assert.Equal(t, wapiVersion+nsGroupEndpoint, getAllNSGroupAuthAPI.Endpoint())
}

func TestNameServerGroupAuthNewGetAllResponse(t *testing.T) {
	setupNSGroupAuthTest("getall")
	response := *getAllNSGroupAuthAPI.ResponseObject().(*[]NSGroupAuth)

	assert.Equal(t, 1, len(response))
	assert.Equal(t, "nsgroup/ZG5zOm5zX2dyb2VwJAByaW1hcnlfWm9uZV9YRlI:TEST_NS_GROUP_AUTH", response[0].Reference)
	assert.Equal(t, "test-ns-group-auth", response[0].Name)
}

func TestNameServerGroupAuthNewUpdateMethod(t *testing.T) {
	setupNSGroupAuthTest("update")
	assert.Equal(t, http.MethodPut, updateNSGroupAuthAPI.Method())
}

func TestNameServerGroupAuthNewUpdateEndpoint(t *testing.T) {
	setupNSGroupAuthTest("update")
	assert.Equal(t, wapiVersion+"/"+nsGroupAuthObject.Reference+"?_return_fields="+strings.Join(returnFields, ","), updateNSGroupAuthAPI.Endpoint())
}

func TestNameServerGroupAuthNewUpdateResponse(t *testing.T) {
	setupNSGroupAuthTest("update")
	response := updateNSGroupAuthAPI.ResponseObject().(*NSGroupAuth)

	assert.Equal(t, "test-ns-group-auth", response.Name)
	assert.Equal(t, "Test NS Group Auth", response.Comment)
	assert.Equal(t, true, *response.GridDefault)
	assert.Equal(t, true, *response.UseExternalPrimary)
}

func TestNameServerGroupAuthNewDeleteMethod(t *testing.T) {
	setupNSGroupAuthTest("delete")
	assert.Equal(t, http.MethodDelete, deleteNSGroupAuthAPI.Method())
}

func TestNameServerGroupAuthNewDeleteEndpoint(t *testing.T) {
	setupNSGroupAuthTest("delete")
	assert.Equal(t, wapiVersion+"/"+nsGroupAuthObject.Reference, deleteNSGroupAuthAPI.Endpoint())
}
