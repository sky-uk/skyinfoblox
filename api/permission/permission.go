package permission

import (
	"github.com/sky-uk/go-rest-api"
	"net/http"
)

const permissionEndpoint = "/wapi/v2.3.1/permission"
const returnFields = "?_return_fields=group,object,permission,resource_type,role"

func NewGet(permissionRef string) *rest.BaseAPI {
	permissionGetAPI := rest.NewBaseAPI(http.MethodGet, permissionEndpoint+permissionRef+returnFields, nil, new(Permission), nil)
	return permissionGetAPI
}

func NewGetAll() *rest.BaseAPI {
	permissionGetAllAPI := rest.NewBaseAPI(http.MethodGet, permissionEndpoint+returnFields, nil, new([]Permission), nil)
	return permissionGetAllAPI
}

func NewCreate(newPermission Permission) *rest.BaseAPI {
	permissionCreateAPI := rest.NewBaseAPI(http.MethodPost, permissionEndpoint, newPermission, new(string), nil)
	return permissionCreateAPI
}

func NewUpdate(updatedPermission Permission) *rest.BaseAPI {
	permissionCreateAPI := rest.NewBaseAPI(http.MethodPut, permissionEndpoint, updatedPermission, new(string), nil)
	return permissionCreateAPI
}

func NewDelete(permissionRef string) *rest.BaseAPI {
	permissionDeleteAPI := rest.NewBaseAPI(http.MethodDelete, permissionEndpoint+permissionRef, nil, new(string), nil)
	return permissionDeleteAPI
}
