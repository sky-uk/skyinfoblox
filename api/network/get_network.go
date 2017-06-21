package network

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetNetworkAPI base object.
type GetNetworkAPI struct {
	*api.BaseAPI
}

// NewGetNetwork returns a new object of type GetNetworkAPI.
func NewGetNetwork(objRef string, returnFields []string) *GetNetworkAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		objRef += returnFields
	}
	this := new(GetNetworkAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+objRef, nil, new(Network))
	return this
}

// GetResponse casts the response object and
// returns the single network object
func (ga GetNetworkAPI) GetResponse() interface{} {
	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*Network)
	}
	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
