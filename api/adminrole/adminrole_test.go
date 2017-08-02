package adminrole

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createAdminRoleAPI, getAllAdminRolesAPI, getAdminRoleAPI, updateAdminRoleAPI, deleteAdminRoleAPI *api.BaseAPI
var reference string
var adminRoleListResponse []AdminRole

func setUpAdminRoleTest() {

	disable := false
	adminRole := AdminRole{
		Reference: "adminrole/b25lLnJvbGUkdGVzdHJvbGUy:testrole2",
		Name:      "testRole",
		Comment:   "Admin Role for tests",
		Disable:   &disable,
	}

	reference = "adminrole/b25lLnJvbGUkdGVzdHJvbGUy:testrole2"

	createAdminRoleAPI = NewCreate(adminRole)
	createAdminRoleAPI.SetResponseObject(&reference)

	getAdminRoleAPI = NewGet(reference)
	getAdminRoleAPI.SetResponseObject(&adminRole)

	roleOne := AdminRole{
		Reference: "adminrole/b25lLnJvbGUkVGVycmFmb3JtIFVzZXI:Terraform%20User",
		Name:      "Terraform User",
	}

	roleTwo := AdminRole{
		Reference: "adminrole/b25lLnJvbGUkdGVzdHJvbGUy:testrole2",
		Name:      "testrole2",
	}

	adminRoleListResponse = append(adminRoleListResponse, roleOne)
	adminRoleListResponse = append(adminRoleListResponse, roleTwo)

	getAllAdminRolesAPI = NewGetAll()
	getAllAdminRolesAPI.SetResponseObject(adminRoleListResponse)

	updateAdminRoleAPI = NewUpdate(reference, adminRole)
	updateAdminRoleAPI.SetResponseObject(&adminRole)

	deleteAdminRoleAPI = NewDelete(reference)
}

func TestAdminRoleNewCreateMethod(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, http.MethodPost, createAdminRoleAPI.Method())
}

func TestAdminRoleNewCreateEndpoint(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, adminRoleEndpoint+"adminrole", createAdminRoleAPI.Endpoint())
}

func TestAdminRoleNewCreateResponse(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, reference, *createAdminRoleAPI.ResponseObject().(*string))
}

func TestAdminRoleNewGetMethod(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, http.MethodGet, getAdminRoleAPI.Method())
}

func TestAdminRoleNewGetEndpoint(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, adminRoleEndpoint+reference+returnFields, getAdminRoleAPI.Endpoint())
}

func TestAdminRoleGetAllMethod(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, http.MethodGet, getAllAdminRolesAPI.Method())
}

func TestAdminRoleGetAllEndpoint(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, adminRoleEndpoint+"adminrole"+returnFields, getAllAdminRolesAPI.Endpoint())
}

func TestAdminRoleNewGetResponse(t *testing.T) {
	setUpAdminRoleTest()
	response := getAdminRoleAPI.ResponseObject().(*AdminRole)

	assert.Equal(t, "testRole", response.Name)
	assert.Equal(t, "Admin Role for tests", response.Comment)
	assert.Equal(t, false, *response.Disable)
}

func TestAdminRoleUpdateMethod(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, http.MethodPut, updateAdminRoleAPI.Method())
}

func TestAdminRoleUpdateEndpoint(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, adminRoleEndpoint+reference+returnFields, updateAdminRoleAPI.Endpoint())
}

func TestAdminRoleUpdateResponse(t *testing.T) {
	setUpAdminRoleTest()
	response := updateAdminRoleAPI.ResponseObject().(*AdminRole)

	assert.Equal(t, "testRole", response.Name)
	assert.Equal(t, "Admin Role for tests", response.Comment)
	assert.Equal(t, false, *response.Disable)
}

func TestAdminRoleDeleteMethod(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, http.MethodDelete, deleteAdminRoleAPI.Method())
}

func TestAdminRoleDeleteEndpoint(t *testing.T) {
	setUpAdminRoleTest()
	assert.Equal(t, adminRoleEndpoint+reference, deleteAdminRoleAPI.Endpoint())
}
