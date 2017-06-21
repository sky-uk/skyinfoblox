package network

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllNetworksAPI base object.
type GetAllNetworksAPI struct {
	*api.BaseAPI
}

// NewGetAllNetworks returns a new object of GetAllARecordsAPI.
func NewGetAllNetworks(fields []string) *GetAllNetworksAPI {
	this := new(GetAllNetworksAPI)
	var url string
	if len(fields) > 0 {
		url = "/wapi/v2.3.1/network?_return_fields=" + strings.Join(fields, ",")
	} else {
		url = "/wapi/v2.3.1/network"
	}

	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]Network))
	return this
}

// GetResponse casts the response object and
// returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllNetworksAPI) GetResponse() interface{} {
	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*[]Network)
	}
	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
