package network

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateNetworkAPI base object.
type CreateNetworkAPI struct {
	*api.BaseAPI
}

// NewCreateNetwork returns a new object of type CreateNetworkAPI.
func NewCreateNetwork(ipAddr string, cidr string) *CreateNetworkAPI {
	this := new(CreateNetworkAPI)
	this.Method(http.MethodPost)
	qPath := fmt.Sprintf("/wapi/v2.3.1/network?network=\"%s/%s\"", ipAddr, cidr)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, qPath, nil, new(string))
	return this
}

// GetResponse casts the response object and
// returns ResponseObject of GetAllARecordsAPI.
func (ga CreateNetworkAPI) GetResponse() string {
	return ga.ResponseObject().(string)
}
