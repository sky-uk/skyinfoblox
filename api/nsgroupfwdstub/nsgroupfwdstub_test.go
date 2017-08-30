package nsgroupfwdstub

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func setupNSGroupFwdStubTest(testType string) (*api.BaseAPI, *NSGroupFwdStub) {

	var nsGroupFwdStubObject NSGroupFwdStub
	var nsGroupFwdStubObjectList []NSGroupFwdStub
	var nsGroupFwdStubExternalServerObject common.ExternalServer
	var nsGroupFwdStubExternalServerList []common.ExternalServer

	nsGroupFwdStubExternalServerObject.Name = "ns1.example.com"
	nsGroupFwdStubExternalServerObject.Address = "192.168.0.1"
	nsGroupFwdStubExternalServerList = make([]common.ExternalServer, 0)
	nsGroupFwdStubExternalServerList = append(nsGroupFwdStubExternalServerList, nsGroupFwdStubExternalServerObject)

	nsGroupFwdStubObject.Name = "test-ns-group-fwd-stub"
	nsGroupFwdStubObject.Comment = "Testing NS Group Forward Stub"
	nsGroupFwdStubObject.Reference = "nsgroup:forwardstubserver/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-fwd-stub"
	nsGroupFwdStubObject.ExternalServers = nsGroupFwdStubExternalServerList

	nsGroupFwdStubObjectList = make([]NSGroupFwdStub, 0)
	nsGroupFwdStubObjectList = append(nsGroupFwdStubObjectList, nsGroupFwdStubObject)

	switch testType {
	case "create":
		createNSGroupFwdStubAPI := NewCreate(nsGroupFwdStubObject)
		createNSGroupFwdStubAPI.SetResponseObject(&nsGroupFwdStubObject.Reference)
		return createNSGroupFwdStubAPI, &nsGroupFwdStubObject
	case "get":
		getNSGroupFwdStubAPI := NewGet(nsGroupFwdStubObject.Reference, RequestReturnFields)
		getNSGroupFwdStubAPI.SetResponseObject(&nsGroupFwdStubObject)
		return getNSGroupFwdStubAPI, &nsGroupFwdStubObject
	case "getall":
		getAllNSGroupFwdStubAPI := NewGetAll()
		getAllNSGroupFwdStubAPI.SetResponseObject(&nsGroupFwdStubObjectList)
		return getAllNSGroupFwdStubAPI, nil
	case "update":
		updateNSGroupFwdStubAPI := NewUpdate(nsGroupFwdStubObject, RequestReturnFields)
		updateNSGroupFwdStubAPI.SetResponseObject(&nsGroupFwdStubObject)
		return updateNSGroupFwdStubAPI, &nsGroupFwdStubObject
	case "delete":
		deleteNSGroupFwdStubAPI := NewDelete(nsGroupFwdStubObject.Reference)
		return deleteNSGroupFwdStubAPI, &nsGroupFwdStubObject
	case "default":
		fmt.Println("Option not implemented")
	}
	return nil, nil
}

func TestNameServerGroupFwdStubNewCreateMethod(t *testing.T) {
	createNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("create")
	assert.Equal(t, http.MethodPost, createNSGroupFwdStubAPI.Method())
}

func TestNameServerGroupFwdStubNewCreateEndpoint(t *testing.T) {
	createNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("create")
	assert.Equal(t, wapiVersion+nsGroupFwdStubEndpoint, createNSGroupFwdStubAPI.Endpoint())
}

func TestNameServerGroupFwdStubNewCreateResponse(t *testing.T) {
	createNSGroupFwdStubAPI, nsGroupFwdStubObject := setupNSGroupFwdStubTest("create")
	response := *createNSGroupFwdStubAPI.ResponseObject().(*string)
	assert.Equal(t, nsGroupFwdStubObject.Reference, response)
}

func TestNameServerGroupFwdStubNewGetMethod(t *testing.T) {
	getNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("get")
	assert.Equal(t, http.MethodGet, getNSGroupFwdStubAPI.Method())
}

func TestNameServerGroupFwdStubNewGetEndpoint(t *testing.T) {
	getNSGroupFwdStubAPI, nsGroupFwdStubObject := setupNSGroupFwdStubTest("get")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdStubObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), getNSGroupFwdStubAPI.Endpoint())
}

func TestNameServerGroupFwdStubNewGetResponse(t *testing.T) {
	getNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("get")
	response := getNSGroupFwdStubAPI.ResponseObject().(*NSGroupFwdStub)

	assert.Equal(t, "test-ns-group-fwd-stub", response.Name)
	assert.Equal(t, "Testing NS Group Forward Stub", response.Comment)
	assert.Equal(t, 1, len(response.ExternalServers))
	assert.Equal(t, "ns1.example.com", response.ExternalServers[0].Name)
	assert.Equal(t, "192.168.0.1", response.ExternalServers[0].Address)
}

func TestNameServerGroupFwdStubNewGetAllMethod(t *testing.T) {
	getAllNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("getall")
	assert.Equal(t, http.MethodGet, getAllNSGroupFwdStubAPI.Method())
}

func TestNameServerGroupFwdStubNewGetAllEndpoint(t *testing.T) {
	getAllNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("getall")
	assert.Equal(t, wapiVersion+nsGroupFwdStubEndpoint, getAllNSGroupFwdStubAPI.Endpoint())
}

func TestNameServerGroupFwdStubNewGetAllResponse(t *testing.T) {
	getAllNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("getall")
	response := *getAllNSGroupFwdStubAPI.ResponseObject().(*[]NSGroupFwdStub)

	assert.Equal(t, 1, len(response))
	assert.Equal(t, "nsgroup:forwardstubserver/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-fwd-stub", response[0].Reference)
	assert.Equal(t, "test-ns-group-fwd-stub", response[0].Name)
}

func TestNameServerGroupFwdStubNewUpdateMethod(t *testing.T) {
	updateNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("update")
	assert.Equal(t, http.MethodPut, updateNSGroupFwdStubAPI.Method())
}

func TestNameServerGroupFwdStubNewUpdateEndpoint(t *testing.T) {
	updateNSGroupFwdStubAPI, nsGroupFwdStubObject := setupNSGroupFwdStubTest("update")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdStubObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), updateNSGroupFwdStubAPI.Endpoint())
}

func TestNameServerGroupFwdStubNewUpdateResponse(t *testing.T) {
	updateNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("update")
	response := updateNSGroupFwdStubAPI.ResponseObject().(*NSGroupFwdStub)

	assert.Equal(t, "test-ns-group-fwd-stub", response.Name)
	assert.Equal(t, "Testing NS Group Forward Stub", response.Comment)
	assert.Equal(t, 1, len(response.ExternalServers))
	assert.Equal(t, "ns1.example.com", response.ExternalServers[0].Name)
	assert.Equal(t, "192.168.0.1", response.ExternalServers[0].Address)
}

func TestNameServerGroupFwdStubNewDeleteMethod(t *testing.T) {
	deleteNSGroupFwdStubAPI, _ := setupNSGroupFwdStubTest("delete")
	assert.Equal(t, http.MethodDelete, deleteNSGroupFwdStubAPI.Method())
}

func TestNameServerGroupFwdStubNewDeleteEndpoint(t *testing.T) {
	deleteNSGroupFwdStubAPI, nsGroupFwdStubObject := setupNSGroupFwdStubTest("delete")
	assert.Equal(t, wapiVersion+"/"+nsGroupFwdStubObject.Reference, deleteNSGroupFwdStubAPI.Endpoint())
}
