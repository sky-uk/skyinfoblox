package nsgroupstub

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createNSGroupStubAPI, getNSGroupStubAPI, getAllNSGroupStubAPI, updateNSGroupStubAPI, deleteNSGroupStubAPI *api.BaseAPI
var nsGroupStubObject NSGroupStub
var nsGroupStubObjectList []NSGroupStub
var nsGroupStubMemberServerObject common.MemberServer
var nsGroupStubMemberServerList []common.MemberServer
var nsGroupStubExternalServerObject common.ExternalServer
var nsGroupStubExternalServerList []common.ExternalServer

func setupNSGroupStubTest(testType string) {

	nsGroupStubExternalServerObject.Name = "ns1.example.com"
	nsGroupStubExternalServerObject.Address = "192.168.0.1"
	nsGroupStubExternalServerList = append(nsGroupStubExternalServerList, nsGroupStubExternalServerObject)

	nsGroupStubMemberServerObject.Name = "grid-member01.example.com"
	nsGroupStubMemberServerObject.PreferredPrimaries = nsGroupStubExternalServerList
	nsGroupStubMemberServerList = append(nsGroupStubMemberServerList, nsGroupStubMemberServerObject)

	nsGroupStubObject.Name = "test-ns-group-stub"
	nsGroupStubObject.Comment = "Testing NS Group Stub"
	nsGroupStubObject.Reference = "nsgroup:stubmember/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-stub"

	nsGroupStubObjectList = append(nsGroupStubObjectList, nsGroupStubObject)

	switch testType {
	case "create":
		createNSGroupStubAPI = NewCreate(nsGroupStubObject)
		createNSGroupStubAPI.SetResponseObject(&nsGroupStubObject.Reference)
	case "get":
		getNSGroupStubAPI = NewGet(nsGroupStubObject.Reference, RequestReturnFields)
		getNSGroupStubAPI.SetResponseObject(&nsGroupStubObject)
	case "getall":
		getAllNSGroupStubAPI = NewGetAll()
		getAllNSGroupStubAPI.SetResponseObject(&nsGroupStubObjectList)
	case "update":
		updateNSGroupStubAPI = NewUpdate(nsGroupStubObject, RequestReturnFields)
		updateNSGroupStubAPI.SetResponseObject(&nsGroupStubObject)
	case "delete":
		deleteNSGroupStubAPI = NewDelete(nsGroupStubObject.Reference)
	case "default":
		fmt.Println("Option not implemented")
	}
}

func TestNameServerGroupStubNewCreateMethod(t *testing.T) {
	setupNSGroupStubTest("create")
	assert.Equal(t, http.MethodPost, createNSGroupStubAPI.Method())
}

func TestNameServerGroupStubNewCreateEndpoint(t *testing.T) {
	setupNSGroupStubTest("create")
	assert.Equal(t, wapiVersion+nsGroupStubEndpoint, createNSGroupStubAPI.Endpoint())
}

func TestNameServerGroupStubNewCreateResponse(t *testing.T) {
	setupNSGroupStubTest("create")
	response := *createNSGroupStubAPI.ResponseObject().(*string)
	assert.Equal(t, nsGroupStubObject.Reference, response)
}

func TestNameServerGroupStubNewGetMethod(t *testing.T) {
	setupNSGroupStubTest("get")
	assert.Equal(t, http.MethodGet, getNSGroupStubAPI.Method())
}

func TestNameServerGroupStubNewGetEndpoint(t *testing.T) {
	setupNSGroupStubTest("get")
	assert.Equal(t, wapiVersion+"/"+nsGroupStubObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), getNSGroupStubAPI.Endpoint())
}

func TestNameServerGroupStubNewGetResponse(t *testing.T) {
	setupNSGroupStubTest("get")
	response := getNSGroupStubAPI.ResponseObject().(*NSGroupStub)

	assert.Equal(t, "test-ns-group-stub", response.Name)
	assert.Equal(t, "Testing NS Group Stub", response.Comment)
}

func TestNameServerGroupStubNewGetAllMethod(t *testing.T) {
	setupNSGroupStubTest("getall")
	assert.Equal(t, http.MethodGet, getAllNSGroupStubAPI.Method())
}

func TestNameServerGroupStubNewGetAllEndpoint(t *testing.T) {
	setupNSGroupStubTest("getall")
	assert.Equal(t, wapiVersion+nsGroupStubEndpoint, getAllNSGroupStubAPI.Endpoint())
}

func TestNameServerGroupStubNewGetAllResponse(t *testing.T) {
	setupNSGroupStubTest("getall")
	response := *getAllNSGroupStubAPI.ResponseObject().(*[]NSGroupStub)

	assert.Equal(t, "nsgroup:stubmember/ZG5zLoL2zX2dyb3VwJHRlc3Q:test-ns-group-stub", response[0].Reference)
	assert.Equal(t, "test-ns-group-stub", response[0].Name)
}

func TestNameServerGroupStubNewUpdateMethod(t *testing.T) {
	setupNSGroupStubTest("update")
	assert.Equal(t, http.MethodPut, updateNSGroupStubAPI.Method())
}

func TestNameServerGroupStubNewUpdateEndpoint(t *testing.T) {
	setupNSGroupStubTest("update")
	assert.Equal(t, wapiVersion+"/"+nsGroupStubObject.Reference+"?_return_fields="+strings.Join(RequestReturnFields, ","), updateNSGroupStubAPI.Endpoint())
}

func TestNameServerGroupStubNewUpdateResponse(t *testing.T) {
	setupNSGroupStubTest("update")
	response := updateNSGroupStubAPI.ResponseObject().(*NSGroupStub)

	assert.Equal(t, "test-ns-group-stub", response.Name)
	assert.Equal(t, "Testing NS Group Stub", response.Comment)
}

func TestNameServerGroupStubNewDeleteMethod(t *testing.T) {
	setupNSGroupStubTest("delete")
	assert.Equal(t, http.MethodDelete, deleteNSGroupStubAPI.Method())
}

func TestNameServerGroupStubNewDeleteEndpoint(t *testing.T) {
	setupNSGroupStubTest("delete")
	assert.Equal(t, wapiVersion+"/"+nsGroupStubObject.Reference, deleteNSGroupStubAPI.Endpoint())
}
