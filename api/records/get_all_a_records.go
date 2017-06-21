package records

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllARecordsAPI base object.
type GetAllARecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllARecords returns a new object of GetAllARecordsAPI.
func NewGetAllARecords(fields []string) *GetAllARecordsAPI {
	var url string
	if len(fields) >= 1 {
		url = "/wapi/v2.3.1/record:a?_return_fields=" + strings.Join(fields, ",")
	} else {
		url = "/wapi/v2.3.1/record:a"
	}

	this := new(GetAllARecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, url, nil, new([]ARecord))
	return this
}

// GetResponse returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllARecordsAPI) GetResponse() interface{} {

	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*[]ARecord)
	}

	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
