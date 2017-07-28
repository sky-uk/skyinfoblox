package adminuser

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

//AdminUserAPI - Base struce
type BaseAdminUserAPI struct {
	*api.BaseAPI
}

var endPoint string

//NewCreateAdminUser - Create function
func NewCreateAdminUser(newUser AdminUser) *BaseAdminUserAPI {
	this := new(BaseAdminUserAPI)
	endPoint = "/wapi/v2.2.2/adminuser"
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, endPoint, newUser, new(string))
	return this
}

//NewGetAdminUser - Get a User
func NewGetAdminUser(ref string, returnFields []string) *BaseAdminUserAPI {
	this := new(BaseAdminUserAPI)
	if returnFields != nil && len(returnFields) > 0 {
		endPoint = fmt.Sprintf("/wapi/v2.2.2/%s/?_return_fields=%s", ref, strings.Join(returnFields, ","))
	} else {
		endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", ref)
	}
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, endPoint, nil, new(AdminUser))
	return this
}

//NewDeleteAdminUse - Deletes the user
func NewDeleteAdminUser(ref string) *BaseAdminUserAPI {
	this := new(BaseAdminUserAPI)
	endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", ref)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, endPoint, nil, new(AdminUser))
	return this
}


// NewUpdateAdminUser - Updates the user
func NewUpdateAdminUser(updateUser AdminUser) *BaseAdminUserAPI {
	this := new(BaseAdminUserAPI)
	endPoint = fmt.Sprintf("/wapi/v2.2.2/%s", updateUser.Ref)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, endPoint, nil, new(AdminUser))
	return this
}

// GetResponse casts the response object to string
func (gu BaseAdminUserAPI) GetResponse() string {
	return *gu.ResponseObject().(*string)
}
