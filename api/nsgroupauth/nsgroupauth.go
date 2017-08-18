package nsgroupauth

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new NSGroupAuth object
func NewCreate(nameServerGroupAuth NSGroupAuth) *api.BaseAPI {
	createNSGroupAuthAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsGroupEndpoint, nameServerGroupAuth, new(string))
	return createNSGroupAuthAPI
}

// NewGetAll : used to get a list of all NSGroupAuth objects
func NewGetAll() *api.BaseAPI {
	getAllNSGroupAuthAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsGroupEndpoint, nil, new([]NSGroupAuth))
	return getAllNSGroupAuthAPI
}

// NewGet : used to get a NSGroupAuth object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getNSGroupAuthAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+reference, nil, new(NSGroupAuth))
	return getNSGroupAuthAPI
}

// NewUpdate : used to update a NSGroupAuth object
func NewUpdate(nameServerGroupAuth NSGroupAuth, returnFields []string) *api.BaseAPI {
	reference := "/" + nameServerGroupAuth.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	updateNSGroupAuthAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+reference, nameServerGroupAuth, new(NSGroupAuth))
	return updateNSGroupAuthAPI
}

// NewDelete : used to delete a NSGroupAuth object
func NewDelete(reference string) *api.BaseAPI {
	deleteNSGroupAuthAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSGroupAuthAPI
}
