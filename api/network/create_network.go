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

// NewCreateNetwork returns a new object of type network.API.
func NewCreateNetwork(net Network) *CreateNetworkAPI {
	this := new(CreateNetworkAPI)
	qPath := fmt.Sprintf("/wapi/v2.3.1/network?network=%s", net.Network)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, qPath, nil, new(string))
	return this
}

// GetResponse casts the response object and
// returns ResponseObject of GetAllARecordsAPI.
func (ga CreateNetworkAPI) GetResponse() string {
	return *ga.ResponseObject().(*string)
}
