package adminuser

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)


var endPoint string

//NewCreateAdminUser - Create function
func NewCreateAdminUser(newUser AdminUser) *api.BaseAPI {
	endPoint = "/wapi/v2.2.2/adminuser"
	createUserAPI := api.NewBaseAPI(http.MethodPost, endPoint, newUser, new(string))
	return createUserAPI
}

//NewGetAdminUser - Get a User
func NewGetAdminUser(ref string, returnFields []string) *api.BaseAPI {
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("/wapi/v2.2.2/%s/?_return_fields=%s", ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", ref)
	}
	updateUserAPI := api.NewBaseAPI(http.MethodGet, endPoint, nil, new(AdminUser))
	return updateUserAPI
}

//NewDeleteAdminUser - Deletes the user
func NewDeleteAdminUser(ref string) *api.BaseAPI {
	endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", ref)
	deleteUserAPI  := api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(AdminUser))
	return deleteUserAPI
}

// NewUpdateAdminUser - Updates the user
func NewUpdateAdminUser(updateUser AdminUser) *api.BaseAPI {
	endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", updateUser.Ref)
	updateUserAPI := api.NewBaseAPI(http.MethodPut, endPoint, nil, new(AdminUser))
	return updateUserAPI
}


