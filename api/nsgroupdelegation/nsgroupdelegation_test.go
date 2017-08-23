package nsgroupdelegation

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createNSGroupDelegationAPI, getNSGroupDelegationAPI, getAllNSGroupDelegationAPI, updateNSGroupDelegationAPI, deleteNSGroupDelegationAPI *api.BaseAPI
var nsGroupDelegationObject NSGroupDelegation
var nsGroupDelegationObjectList []NSGroupDelegation
var externalServerObject common.ExternalServer
var externalServerList []common.ExternalServer

func setupNSGroupDelegationTest(testType string) {

	externalServerObject = common.ExternalServer{
		Name:    "ns1.example.com",
		Address: "192.168.1.100",
	}
	externalServerList = append(externalServerList, externalServerObject)

	nsGroupDelegationObject = NSGroupDelegation{
		Reference:  "nsgroup:delegation/ZG5zOm5zX2dyb2VwJAByaW1hcnlfWm9uZV9YRlI:TEST_NS_GROUP_DELEGATION",
		Comment:    "Test NS Group Delegation",
		Name:       "test-ns-group-delegation",
		DelegateTo: externalServerList,
	}
	nsGroupDelegationObjectList = append(nsGroupDelegationObjectList, nsGroupDelegationObject)

	switch testType {
	case "create":
		createNSGroupDelegationAPI = NewCreate(nsGroupDelegationObject)
		createNSGroupDelegationAPI.SetResponseObject(&nsGroupDelegationObject.Reference)
	case "get":
		getNSGroupDelegationAPI = NewGet(nsGroupDelegationObject.Reference, RequestReturnFields)
		getNSGroupDelegationAPI.SetResponseObject(&nsGroupDelegationObject)
	case "getall":
		getAllNSGroupDelegationAPI = NewGetAll()
		getAllNSGroupDelegationAPI.SetResponseObject(&nsGroupDelegationObjectList)
	case "update":
		updateNSGroupDelegationAPI = NewUpdate(nsGroupDelegationObject, RequestReturnFields)
		updateNSGroupDelegationAPI.SetResponseObject(&nsGroupDelegationObject)
	case "delete":
		deleteNSGroupDelegationAPI = NewDelete(nsGroupDelegationObject.Reference)
	case "default":
		fmt.Println("Option not implemented")
	}
}

func TestNameServerGroupDelegationNewCreateMethod(t *testing.T) {
	setupNSGroupDelegationTest("create")
	assert.Equal(t, http.MethodPost, createNSGroupDelegationAPI.Method())
}

func TestNameServerGroupDelegationNewCreateEndpoint(t *testing.T) {
	setupNSGroupDelegationTest("create")
	assert.Equal(t, wapiVersion+nsGroupDelegationEndpoint, createNSGroupDelegationAPI.Endpoint())
}

func TestNameServerGroupDelegationNewCreateResponse(t *testing.T) {
	setupNSGroupDelegationTest("create")
	response := *createNSGroupDelegationAPI.ResponseObject().(*string)
	assert.Equal(t, nsGroupDelegationObject.Reference, response)
}

func TestNameServerGroupDelegationNewGetMethod(t *testing.T) {
	setupNSGroupDelegationTest("get")
	assert.Equal(t, http.MethodGet, getNSGroupDelegationAPI.Method())
}

func TestNameServerGroupDelegationNewGetEndpoint(t *testing.T) {
	setupNSGroupDelegationTest("get")
	assert.Equal(t, wapiVersion+"/"+nsGroupDelegationObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), getNSGroupDelegationAPI.Endpoint())
}

func TestNameServerGroupDelegationNewGetResponse(t *testing.T) {
	setupNSGroupDelegationTest("get")
	response := getNSGroupDelegationAPI.ResponseObject().(*NSGroupDelegation)

	assert.Equal(t, "test-ns-group-delegation", response.Name)
	assert.Equal(t, "Test NS Group Delegation", response.Comment)
	assert.Equal(t, "ns1.example.com", response.DelegateTo[0].Name)
	assert.Equal(t, "192.168.1.100", response.DelegateTo[0].Address)
}

func TestNameServerGroupDelegationNewGetAllMethod(t *testing.T) {
	setupNSGroupDelegationTest("getall")
	assert.Equal(t, http.MethodGet, getAllNSGroupDelegationAPI.Method())
}

func TestNameServerGroupDelegationNewGetAllEndpoint(t *testing.T) {
	setupNSGroupDelegationTest("getall")
	assert.Equal(t, wapiVersion+nsGroupDelegationEndpoint, getAllNSGroupDelegationAPI.Endpoint())
}

func TestNameServerGroupDelegationNewGetAllResponse(t *testing.T) {
	setupNSGroupDelegationTest("getall")
	response := *getAllNSGroupDelegationAPI.ResponseObject().(*[]NSGroupDelegation)

	assert.Equal(t, "nsgroup:delegation/ZG5zOm5zX2dyb2VwJAByaW1hcnlfWm9uZV9YRlI:TEST_NS_GROUP_DELEGATION", response[0].Reference)
	assert.Equal(t, "test-ns-group-delegation", response[0].Name)
}

func TestNameServerGroupDelegationNewUpdateMethod(t *testing.T) {
	setupNSGroupDelegationTest("update")
	assert.Equal(t, http.MethodPut, updateNSGroupDelegationAPI.Method())
}

func TestNameServerGroupDelegationNewUpdateEndpoint(t *testing.T) {
	setupNSGroupDelegationTest("update")
	assert.Equal(t, wapiVersion+"/"+nsGroupDelegationObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), updateNSGroupDelegationAPI.Endpoint())
}

func TestNameServerGroupDelegationNewUpdateResponse(t *testing.T) {
	setupNSGroupDelegationTest("update")
	response := updateNSGroupDelegationAPI.ResponseObject().(*NSGroupDelegation)

	assert.Equal(t, "test-ns-group-delegation", response.Name)
	assert.Equal(t, "Test NS Group Delegation", response.Comment)
	assert.Equal(t, "ns1.example.com", response.DelegateTo[0].Name)
	assert.Equal(t, "192.168.1.100", response.DelegateTo[0].Address)
}

func TestNameServerGroupDelegationNewDeleteMethod(t *testing.T) {
	setupNSGroupDelegationTest("delete")
	assert.Equal(t, http.MethodDelete, deleteNSGroupDelegationAPI.Method())
}

func TestNameServerGroupDelegationNewDeleteEndpoint(t *testing.T) {
	setupNSGroupDelegationTest("delete")
	assert.Equal(t, wapiVersion+"/"+nsGroupDelegationObject.Reference, deleteNSGroupDelegationAPI.Endpoint())
}
