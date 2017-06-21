package zoneauth

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetSingleZoneAuthAPI type
type GetSingleZoneAuthAPI struct {
	*api.BaseAPI
}

// NewGetSingleZone : returns a zone's details.
func NewGetSingleZone(ref string, returnFieldList []string) *GetSingleZoneAuthAPI {
	if returnFieldList != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFieldList, ",")
		ref += returnFields
	}
	this := new(GetSingleZoneAuthAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+ref, nil, new(DNSZone))
	return this
}

// GetResponse : returns response obeject from GetSingleZone
func (gsz GetSingleZoneAuthAPI) GetResponse() interface{} {
	return *gsz.ResponseObject().(*DNSZone)

	var errStruct api.RespError
	err := json.Unmarshal(gsz.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
