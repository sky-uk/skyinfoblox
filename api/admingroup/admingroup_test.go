package admingroup

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createAdminGroupAPI, getAllAdminGroupAPI, getAdminGroupAPI, updateAdminGroupAPI, deleteAdminGroupAPI *api.BaseAPI
var returnFields []string
var reference string

func setupAdminGroupTest() {

	superUser := true
	disable := false
	reference = "admingroup/b25lLmFkbWluX2dyb3VwJC5jbG91ZC1hcGktb25seQ:test"
	adminGroup := IBXAdminGroup{
		Reference:      reference,
		AccessMethod:   []string{"API"},
		Comment:        "API Access only",
		Disable:        &disable,
		EmailAddresses: []string{"test@example-test.com"},
		Name:           "test",
		Roles:          []string{"test-role"},
		SuperUser:      &superUser,
	}
	returnFields = []string{"name", "comment", "access_method", "disable", "email_addresses", "roles", "superuser"}

	getAllAdminGroupReference := IBXAdminGroupReference{
		Reference:      "admingroup/b25lLmFkbWluX2dyb3VwJC5jbG91ZC1hcGktb25seQ:test",
		AdminGroupName: "test",
	}
	adminGroupList := make([]IBXAdminGroupReference, 0)
	adminGroupList = append(adminGroupList, getAllAdminGroupReference)

	createAdminGroupAPI = NewCreate(adminGroup)
	createAdminGroupAPI.SetResponseObject(&reference)

	getAllAdminGroupAPI = NewGetAll()
	getAllAdminGroupAPI.SetResponseObject(&adminGroupList)

	getAdminGroupAPI = NewGet(reference, returnFields)
	getAdminGroupAPI.SetResponseObject(&adminGroup)

	updateAdminGroupAPI = NewUpdate(adminGroup, returnFields)
	updateAdminGroupAPI.SetResponseObject(&adminGroup)

	deleteAdminGroupAPI = NewDelete(reference)
}

func TestAdminGroupNewCreateMethod(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, http.MethodPost, createAdminGroupAPI.Method())
}

func TestAdminGroupNewCreateEndpoint(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, wapiVersion+adminGroupEndpoint, createAdminGroupAPI.Endpoint())
}

func TestAdminGroupNewCreateResponse(t *testing.T) {
	setupAdminGroupTest()
	response := *createAdminGroupAPI.ResponseObject().(*string)

	assert.Equal(t, reference, response)
}

func TestAdminGroupNewGetAllMethod(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, http.MethodGet, getAllAdminGroupAPI.Method())
}

func TestAdminGroupNewGetAllEndpoint(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, wapiVersion+adminGroupEndpoint, getAllAdminGroupAPI.Endpoint())
}

func TestAdminGroupNewGetAllResponse(t *testing.T) {
	setupAdminGroupTest()
	response := *getAllAdminGroupAPI.ResponseObject().(*[]IBXAdminGroupReference)

	assert.Equal(t, "admingroup/b25lLmFkbWluX2dyb3VwJC5jbG91ZC1hcGktb25seQ:test", response[0].Reference)
	assert.Equal(t, "test", response[0].AdminGroupName)
}

func TestAdminGroupNewGetMethod(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, http.MethodGet, getAdminGroupAPI.Method())
}

func TestAdminGroupNewGetEndpoint(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, wapiVersion+"/"+reference+"?_return_fields="+strings.Join(returnFields, ","), getAdminGroupAPI.Endpoint())
}

func TestAdminGroupNewGetResponse(t *testing.T) {
	setupAdminGroupTest()
	response := getAdminGroupAPI.ResponseObject().(*IBXAdminGroup)

	assert.Equal(t, "test", response.Name)
	assert.Equal(t, "API", response.AccessMethod[0])
	assert.Equal(t, "API Access only", response.Comment)
	assert.Equal(t, false, *response.Disable)
	assert.Equal(t, "test@example-test.com", response.EmailAddresses[0])
	assert.Equal(t, "test-role", response.Roles[0])
	assert.Equal(t, true, *response.SuperUser)
}

func TestAdminGroupNewUpdateMethod(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, http.MethodPut, updateAdminGroupAPI.Method())
}

func TestAdminGroupNewUpdateEndpoint(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, wapiVersion+"/"+reference+"?_return_fields="+strings.Join(returnFields, ","), updateAdminGroupAPI.Endpoint())
}

func TestAdminGroupNewUpdateResponse(t *testing.T) {
	setupAdminGroupTest()

	response := updateAdminGroupAPI.ResponseObject().(*IBXAdminGroup)

	assert.Equal(t, "test", response.Name)
	assert.Equal(t, "API", response.AccessMethod[0])
	assert.Equal(t, "API Access only", response.Comment)
	assert.Equal(t, false, *response.Disable)
	assert.Equal(t, "test@example-test.com", response.EmailAddresses[0])
	assert.Equal(t, "test-role", response.Roles[0])
	assert.Equal(t, true, *response.SuperUser)
}

func TestAdminGroupNewDeleteMethod(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, http.MethodDelete, deleteAdminGroupAPI.Method())
}

func TestAdminGroupNewDeleteEndpoint(t *testing.T) {
	setupAdminGroupTest()
	assert.Equal(t, wapiVersion+"/"+reference, deleteAdminGroupAPI.Endpoint())
}
