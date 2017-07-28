package adminuser

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func userSetup(action string) *api.BaseAPI {
	disableUser := true
	groupList := []string{"Group1"}
	adminUser := AdminUser{
		Ref:      "adminuser/blablalba:SVC-APP-UNIT-TEST",
		Name:     "SVC-APP-UNIT-TEST",
		Comment:  "This is a unit test user",
		Groups:   groupList,
		Password: "Ultr@S@f3P@SSw0rD",
		Disable:  &disableUser,
		Email:    "user@domain.com",
	}
	switch action {
	case "create":
		userAPI := NewCreateAdminUser(adminUser)
		response := "/wapi/v2.2.2/adminuser/blablalba:SVC-APP-UNIT-TEST"
		userAPI.SetResponseObject(&response)
		return userAPI
	case "delete":
		userAPI := NewDeleteAdminUser("adminuser/blablalba:SVC-APP-UNIT-TEST")
		return userAPI
	case "get":
		returnFields := []string{"name", "comment"}
		userAPI := NewGetAdminUser("adminuser/blablalba:SVC-APP-UNIT-TEST", returnFields)
		return userAPI
	case "update":
		userAPI := NewUpdateAdminUser(adminUser)
		return userAPI
	default:
		return nil
	}
}

// TestUserCreateMethod - Test User Creation Method
func TestUserCreateMethod(t *testing.T) {
	newUser := userSetup("create")
	assert.Equal(t, http.MethodPost, newUser.Method())
}

// TestUserCreateEndpoint - Test User Creation Endpoint
func TestUserCreateEndpoint(t *testing.T) {
	newUser := userSetup("create")
	assert.Equal(t, "/wapi/v2.2.2/adminuser", newUser.Endpoint())
}

// TestUserDeleteMethod - Test User Deletion method
func TestUserDeleteMethod(t *testing.T) {
	newUser := userSetup("delete")
	assert.Equal(t, http.MethodDelete, newUser.Method())
}

// TestUserDeleteEndpoint - Test user deletion endpoint
func TestUserDeleteEndpoint(t *testing.T) {
	newUser := userSetup("delete")
	assert.Equal(t, "/wapi/v2.2.2/adminuser/blablalba:SVC-APP-UNIT-TEST", newUser.Endpoint())

}

func TestUserResponse(t *testing.T) {
	newUser := userSetup("create")
	fmt.Println(newUser)
	response := *newUser.ResponseObject().(*string)
	assert.Equal(t, "/wapi/v2.2.2/adminuser/blablalba:SVC-APP-UNIT-TEST", response)
}

func TestGetUserMethod(t *testing.T) {
	newUser := userSetup("get")
	assert.Equal(t, http.MethodGet, newUser.Method())
}

func TestGetUserEndpoint(t *testing.T) {
	newUser := userSetup("get")
	assert.Equal(t, "/wapi/v2.2.2/adminuser/blablalba:SVC-APP-UNIT-TEST/?_return_fields=name,comment", newUser.Endpoint())
}

func TestUpdateUserMethod(t *testing.T) {
	newUser := userSetup("update")
	assert.Equal(t, http.MethodPut, newUser.Method())
}

func TestUpdateUserEndpoint(t *testing.T) {
	newUser := userSetup("update")
	assert.Equal(t, "/wapi/v2.2.2/adminuser/blablalba:SVC-APP-UNIT-TEST", newUser.Endpoint())
}
