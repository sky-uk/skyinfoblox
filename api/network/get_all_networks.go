package network

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetAllNetworksAPI base object.
type GetAllNetworksAPI struct {
	*api.BaseAPI
}

// NewGetAllNetworks returns a new object of GetAllARecordsAPI.
func NewGetAllNetworks() *GetAllNetworksAPI {
	this := new(GetAllNetworksAPI)
	this.BaseAPI = api.NewBaseAPI(
		http.MethodGet,
		"/wapi/v2.3.1/network",
		nil,
		new([]Network),
	)
	return this
}

// GetResponse casts the response object and
// returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllNetworksAPI) GetResponse() []Network {
	return *ga.ResponseObject().(*[]Network)
}
