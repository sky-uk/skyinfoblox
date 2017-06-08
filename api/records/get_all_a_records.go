package records

import (
	"github.com/sky-uk/go-infoblox/api"
	"net/http"
	"strings"
)

// GetAllARecordsAPI base object.
type GetAllARecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllARecords returns a new object of GetAllARecordsAPI.
func NewGetAllARecords(fields []string) *GetAllARecordsAPI {
	returnFields := ""
	if fields != nil {
		returnFields = "?_return_fields=" + strings.Join(fields, ",")
	}
	this := new(GetAllARecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/record:a"+returnFields, nil, new([]ARecord))
	return this
}

// GetResponse returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllARecordsAPI) GetResponse() *[]ARecord {
	return ga.ResponseObject().(*[]ARecord)
}
