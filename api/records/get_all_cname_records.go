package records

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetAllCNAMERecordsAPI base object.
type GetAllCNAMERecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllCNAMERecords returns a new object of GetAllCNAMERecordsAPI.
func NewGetAllCNAMERecords(fields []string) *GetAllCNAMERecordsAPI {
	returnFields := ""
	if fields != nil {
		returnFields = "?_return_fields=" + strings.Join(fields, ",")
	}
	this := new(GetAllCNAMERecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/record:cname"+returnFields, nil, new([]CNAMERecord))
	return this
}

// GetResponse returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllCNAMERecordsAPI) GetResponse() []CNAMERecord {
	return *ga.ResponseObject().(*[]CNAMERecord)
}
