package records

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllSRVRecordsAPI base object.
type GetAllSRVRecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllSRVRecords returns a new object of GetAllSRVRecordsAPI.
func NewGetAllSRVRecords(fields []string) *GetAllSRVRecordsAPI {
	var url string
	if len(fields) >= 1 {
		url = "/wapi/v2.3.1/record:srv?_return_fields=" + strings.Join(fields, ",")
	} else {
		url = "/wapi/v2.3.1/record:srv"
	}

	this := new(GetAllSRVRecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]SRVRecord))
	return this
}

// GetResponse returns ResponseObject of GetAllSRVRecordsAPI.
func (ga GetAllSRVRecordsAPI) GetResponse() interface{} {
	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*[]SRVRecord)
	}

	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
