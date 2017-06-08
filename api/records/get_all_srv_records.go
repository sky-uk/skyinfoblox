package records

import (
	"github.com/sky-uk/go-infoblox/api"
	"net/http"
	"strings"
)

// GetAllSRVRecordsAPI base object.
type GetAllSRVRecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllSRVRecords returns a new object of GetAllSRVRecordsAPI.
func NewGetAllSRVRecords(fields []string) *GetAllSRVRecordsAPI {
	returnFields := ""
	if fields != nil {
		returnFields = "?_return_fields=" + strings.Join(fields, ",")
	}
	this := new(GetAllSRVRecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/record:srv"+returnFields, nil, new([]SRVRecord))
	return this
}

// GetResponse returns ResponseObject of GetAllSRVRecordsAPI.
func (ga GetAllSRVRecordsAPI) GetResponse() *[]SRVRecord {
	return ga.ResponseObject().(*[]SRVRecord)
}
