package network

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetNetworkAPI base object.
type GetNetworkAPI struct {
	*api.BaseAPI
}

// NewGetNetwork returns a new object of type GetNetworkAPI.
func NewGetNetwork(objRef string) *GetNetworkAPI {
	this := new(GetNetworkAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+objRef, nil, new(Network))
	return this
}

// GetResponse casts the response object and
// returns the single network object
func (gn GetNetworkAPI) GetResponse() Network {
	return *gn.ResponseObject().(*Network)
}
