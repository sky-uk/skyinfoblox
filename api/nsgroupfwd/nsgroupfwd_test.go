package nsgroupfwd

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createNSGroupFwdAPI, getNSGroupFwdAPI, getAllNSGroupFwdAPI, updateNSGroupFwdAPI, deleteNSGroupFwdAPI *api.BaseAPI
var nsGroupFwdObject NSGroupFwd
var nsGroupFwdObjectList []NSGroupFwd
var nsForwardingExternalServerObject common.ExternalServer
var nsForwardingExternalServerList []common.ExternalServer
var nsForwardingServerObject common.ForwardingMemberServer
var nsForwardingServerList []common.ForwardingMemberServer

func setupNSGroupFwdTest(testType string) {

	nsForwardingExternalServerObject = common.ExternalServer{
		Address: "192.168.0.12",
		Name:    "ns1.example.com",
	}
	nsForwardingExternalServerList = append(nsForwardingExternalServerList, nsForwardingExternalServerObject)

	forwardersOnly := true
	useOverrideForwarders := false
	nsForwardingServerObject = common.ForwardingMemberServer{
		Name:                  "test-ns-group-fwd-servers",
		ForwardTo:             nsForwardingExternalServerList,
		ForwardersOnly:        &forwardersOnly,
		UseOverrideForwarders: &useOverrideForwarders,
	}
	nsForwardingServerList = append(nsForwardingServerList, nsForwardingServerObject)

	nsGroupFwdObject = NSGroupFwd{
		Name:              "test-ns-group-fwd",
		Comment:           "Testing NS Group Forwarding",
		ForwardingServers: nsForwardingServerList,
		Reference:         "nsgroup:forwardingmember/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-fwd",
	}
	nsGroupFwdObjectList = append(nsGroupFwdObjectList, nsGroupFwdObject)

	switch testType {
	case "create":
		createNSGroupFwdAPI = NewCreate(nsGroupFwdObject)
		createNSGroupFwdAPI.SetResponseObject(&nsGroupFwdObject.Reference)
	case "get":
		getNSGroupFwdAPI = NewGet(nsGroupFwdObject.Reference, RequestReturnFields)
		getNSGroupFwdAPI.SetResponseObject(&nsGroupFwdObject)
	case "getall":
		getAllNSGroupFwdAPI = NewGetAll()
		getAllNSGroupFwdAPI.SetResponseObject(&nsGroupFwdObjectList)
	case "update":
		updateNSGroupFwdAPI = NewUpdate(nsGroupFwdObject, RequestReturnFields)
		updateNSGroupFwdAPI.SetResponseObject(&nsGroupFwdObject)
	case "delete":
		deleteNSGroupFwdAPI = NewDelete(nsGroupFwdObject.Reference)
	case "default":
		fmt.Println("Option not implemented")
	}
}

func TestNameServerGroupFwdNewCreateMethod(t *testing.T) {
	setupNSGroupFwdTest("create")
	assert.Equal(t, http.MethodPost, createNSGroupFwdAPI.Method())
}

func TestNameServerGroupFwdNewCreateEndpoint(t *testing.T) {
	setupNSGroupFwdTest("create")
	assert.Equal(t, wapiVersion+nsGroupFwdEndpoint, createNSGroupFwdAPI.Endpoint())
}

func TestNameServerGroupFwdNewCreateResponse(t *testing.T) {
	setupNSGroupFwdTest("create")
	response := *createNSGroupFwdAPI.ResponseObject().(*string)
	assert.Equal(t, nsGroupFwdObject.Reference, response)
}

func TestNameServerGroupFwdNewGetMethod(t *testing.T) {
	setupNSGroupFwdTest("get")
	assert.Equal(t, http.MethodGet, getNSGroupFwdAPI.Method())
}

func TestNameServerGroupFwdNewGetEndpoint(t *testing.T) {
	setupNSGroupFwdTest("get")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), getNSGroupFwdAPI.Endpoint())
}

func TestNameServerGroupFwdNewGetResponse(t *testing.T) {
	setupNSGroupFwdTest("get")
	response := getNSGroupFwdAPI.ResponseObject().(*NSGroupFwd)

	assert.Equal(t, "test-ns-group-fwd", response.Name)
	assert.Equal(t, "Testing NS Group Forwarding", response.Comment)
	assert.Equal(t, "test-ns-group-fwd-servers", response.ForwardingServers[0].Name)
	assert.Equal(t, true, *response.ForwardingServers[0].ForwardersOnly)
	assert.Equal(t, false, *response.ForwardingServers[0].UseOverrideForwarders)
	assert.Equal(t, "ns1.example.com", response.ForwardingServers[0].ForwardTo[0].Name)
	assert.Equal(t, "192.168.0.12", response.ForwardingServers[0].ForwardTo[0].Address)
}

func TestNameServerGroupFwdNewGetAllMethod(t *testing.T) {
	setupNSGroupFwdTest("getall")
	assert.Equal(t, http.MethodGet, getAllNSGroupFwdAPI.Method())
}

func TestNameServerGroupFwdNewGetAllEndpoint(t *testing.T) {
	setupNSGroupFwdTest("getall")
	assert.Equal(t, wapiVersion+nsGroupFwdEndpoint, getAllNSGroupFwdAPI.Endpoint())
}

func TestNameServerGroupFwdNewGetAllResponse(t *testing.T) {
	setupNSGroupFwdTest("getall")
	response := *getAllNSGroupFwdAPI.ResponseObject().(*[]NSGroupFwd)

	assert.Equal(t, "nsgroup:forwardingmember/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-fwd", response[0].Reference)
	assert.Equal(t, "test-ns-group-fwd", response[0].Name)
}

func TestNameServerGroupFwdNewUpdateMethod(t *testing.T) {
	setupNSGroupFwdTest("update")
	assert.Equal(t, http.MethodPut, updateNSGroupFwdAPI.Method())
}

func TestNameServerGroupFwdNewUpdateEndpoint(t *testing.T) {
	setupNSGroupFwdTest("update")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), updateNSGroupFwdAPI.Endpoint())
}

func TestNameServerGroupFwdNewUpdateResponse(t *testing.T) {
	setupNSGroupFwdTest("update")
	response := updateNSGroupFwdAPI.ResponseObject().(*NSGroupFwd)

	assert.Equal(t, "test-ns-group-fwd", response.Name)
	assert.Equal(t, "Testing NS Group Forwarding", response.Comment)
	assert.Equal(t, "test-ns-group-fwd-servers", response.ForwardingServers[0].Name)
	assert.Equal(t, true, *response.ForwardingServers[0].ForwardersOnly)
	assert.Equal(t, false, *response.ForwardingServers[0].UseOverrideForwarders)
	assert.Equal(t, "ns1.example.com", response.ForwardingServers[0].ForwardTo[0].Name)
	assert.Equal(t, "192.168.0.12", response.ForwardingServers[0].ForwardTo[0].Address)
}

func TestNameServerGroupFwdNewDeleteMethod(t *testing.T) {
	setupNSGroupFwdTest("delete")
	assert.Equal(t, http.MethodDelete, deleteNSGroupFwdAPI.Method())
}

func TestNameServerGroupFwdNewDeleteEndpoint(t *testing.T) {
	setupNSGroupFwdTest("delete")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdObject.Reference, deleteNSGroupFwdAPI.Endpoint())
}
