package permission

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/json"
)

func createPermissionSetup() *api.BaseAPI {
	newPermission := Permission{
		ResourceType: "VIEW",
		Permission: "WRITE",
		Role:       "DNS Admin",
	}
	createPermissionAPI := NewCreate(newPermission)
	createPermissionAPI.SetResponseObject("test response object")
	return createPermissionAPI
}

func TestCreatePermissionMethod(t *testing.T) {
	newPermission := createPermissionSetup()
	assert.Equal(t, "POST", newPermission.Method())
}

func TestCreatePermissionEndpoint(t *testing.T) {
	newPermission := createPermissionSetup()
	assert.Equal(t, "/wapi/v2.3.1/permission", newPermission.Endpoint())
}


func TestCreatePermissionMarshalling(t *testing.T) {
	newPermission := createPermissionSetup()
	expectedJSON := `{"permission":"WRITE","resource_type":"VIEW","role":"DNS Admin"}`
	jsonBytes, err := json.Marshal(newPermission.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestCreateResponseObject(t *testing.T)  {
	response := createPermissionSetup().ResponseObject().(string)
	assert.Equal(t,"test response object",response)

 }

func getPermissionSetup() *api.BaseAPI {
	objRef := "permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5yb2xlJEROUyBBZG1pbi4:DNS%20Admin/WRITE"
	return NewGet(objRef)
}

func TestGetPermissionMethod(t *testing.T)  {
	getPermission := getPermissionSetup()
	assert.Equal(t,"GET",getPermission.Method())
}

func TestGetPermissionEndpoint(t *testing.T)  {
	getPermission := getPermissionSetup()
	assert.Equal(t,"/wapi/v2.3.1/permission/b25lLmhpZXJfcnVsZSQuY29tLmluZm9ibG94LmRucy56b25lJC4uLi4uY29tLmluZm9ibG94Lm9uZS5yb2xlJEROUyBBZG1pbi4:DNS%20Admin/WRITE?_return_fields=group,object,permission,resource_type,role",getPermission.Endpoint())
}

