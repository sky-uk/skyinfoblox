package records

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllTXTRecordsAPI base object.
type GetAllTXTRecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllTXTRecords returns a new object of GetAllTXTRecordsAPI.
func NewGetAllTXTRecords(fields []string) *GetAllTXTRecordsAPI {
	var url string
	if len(fields) >= 1 {
		url = "/wapi/v2.3.1/record:txt?_return_fields=" + strings.Join(fields, ",")
	} else {
		url = "/wapi/v2.3.1/record:txt"
	}
	this := new(GetAllTXTRecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]TXTRecord))
	return this
}

// GetResponse returns ResponseObject of GetAllTXTRecordsAPI.
func (ga GetAllTXTRecordsAPI) GetResponse() interface{} {
	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*[]TXTRecord)
	}

	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
