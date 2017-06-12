package network

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// UpdateNetworkAPI base object.
type UpdateNetworkAPI struct {
	*api.BaseAPI
}

// NewUpdateNetwork returns a new object of type UpdateNetworkAPI.
// It accepts in input a valid network object reference and a
// string representing the key,value pairs of the paramenters to
// update (in the form k1=v1&k2=v2&...)
func NewUpdateNetwork(objRef string, fieldsTOUpdate string) *UpdateNetworkAPI {
	this := new(UpdateNetworkAPI)
	qPath := fmt.Sprintf("/wapi/v2.3.1/%s", objRef)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, qPath, fieldsTOUpdate, new(string))
	return this
}

// GetResponse casts the response object and
// returns the string representing the updated object reference
// or nil in case of errors
func (ga UpdateNetworkAPI) GetResponse() string {
	return ga.ResponseObject().(string)
}
