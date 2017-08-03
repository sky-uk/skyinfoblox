package permission

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createPermissionAPI, getPermissionAPI, getAllPermissionsAPI, updatePermissionAPI, deletePermissionAPI *api.BaseAPI
var reference string
var permissionOne, permissionTwo Permission
var permissionListResponse []Permission

func setUpPermissionTest() {

	newPermission := Permission{
		Reference:    "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5hZG1pbl9ncm91cCQuY2R0ZXN0Z3JvdXAyLmRucy5iaW5kX2E:cdtestgroup2/READ",
		ResourceType: "VIEW",
		Permission:   "WRITE",
		Role:         "DNS Admin",
	}

	reference = "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5hZG1pbl9ncm91cCQuY2R0ZXN0Z3JvdXAyLmRucy5iaW5kX2E:cdtestgroup2/READ"

	createPermissionAPI = NewCreate(newPermission)
	createPermissionAPI.SetResponseObject(reference)

	getPermissionAPI = NewGet(reference)
	getPermissionAPI.SetResponseObject(newPermission)

	permissionOne = Permission{
		Reference:    "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5hZG1pbl9ncm91cCQuY2R0ZXN0Z3JvdXAyLmRucy5iaW5kX2E:cdtestgroup2/READ",
		ResourceType: "VIEW",
		Permission:   "WRITE",
		Role:         "DNS Admin",
	}

	permissionTwo = Permission{
		Reference:    "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5yb2xlJHRlc3Ryb2xlMi5kbnMuYmluZF9h:testrole2/READ",
		ResourceType: "A",
		Permission:   "READ",
		Role:         "testrole2",
	}

	permissionListResponse = append(permissionListResponse, permissionOne)
	permissionListResponse = append(permissionListResponse, permissionTwo)

	getAllPermissionsAPI = NewGetAll()
	getAllPermissionsAPI.SetResponseObject(&permissionListResponse)

	updatePermissionAPI = NewUpdate(reference, newPermission)
	updatePermissionAPI.SetResponseObject(newPermission)

	deletePermissionAPI = NewDelete(reference)
}

func TestCreatePermissionMethod(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, http.MethodPost, createPermissionAPI.Method())
}

func TestCreatePermissionEndpoint(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, permissionEndpoint+"permission", createPermissionAPI.Endpoint())
}

func TestCreateResponseObject(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, reference, createPermissionAPI.ResponseObject().(string))
}

func TestGetPermissionMethod(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, http.MethodGet, getPermissionAPI.Method())
}

func TestGetPermissionEndpoint(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, permissionEndpoint+reference+returnFields, getPermissionAPI.Endpoint())
}

func TestGetPermissionResponse(t *testing.T) {
	setUpPermissionTest()
	response := getPermissionAPI.ResponseObject().(Permission)

	assert.Equal(t, "WRITE", response.Permission)
	assert.Equal(t, "DNS Admin", response.Role)
	assert.Equal(t, "VIEW", response.ResourceType)
	assert.Equal(t, "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5hZG1pbl9ncm91cCQuY2R0ZXN0Z3JvdXAyLmRucy5iaW5kX2E:cdtestgroup2/READ", response.Reference)
}

func TestGetAllPermissionsMethod(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, http.MethodGet, getAllPermissionsAPI.Method())
}

func TestGetAllPermissionEndpoint(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, permissionEndpoint+"permission"+returnFields, getAllPermissionsAPI.Endpoint())
}

func TestGetAllPermissionResponse(t *testing.T) {
	setUpPermissionTest()
	response := getAllPermissionsAPI.ResponseObject().(*[]Permission)
	assert.Equal(t, permissionOne, (*response)[0])
	assert.Equal(t, permissionTwo, (*response)[1])

}

func TestUpdatePermissionMethod(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, http.MethodPut, updatePermissionAPI.Method())
}

func TestUpdatePermissionEndpoint(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, permissionEndpoint+reference+returnFields, updatePermissionAPI.Endpoint())
}

func TestUpdatePermissionResponse(t *testing.T) {
	setUpPermissionTest()
	response := updatePermissionAPI.ResponseObject().(Permission)

	assert.Equal(t, "WRITE", response.Permission)
	assert.Equal(t, "DNS Admin", response.Role)
	assert.Equal(t, "VIEW", response.ResourceType)
	assert.Equal(t, "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5hZG1pbl9ncm91cCQuY2R0ZXN0Z3JvdXAyLmRucy5iaW5kX2E:cdtestgroup2/READ", response.Reference)
}

func TestDeletePermissionMethod(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, http.MethodDelete, deletePermissionAPI.Method())
}

func TestDeletePermissionEndpoint(t *testing.T) {
	setUpPermissionTest()
	assert.Equal(t, permissionEndpoint+reference, deletePermissionAPI.Endpoint())
}
