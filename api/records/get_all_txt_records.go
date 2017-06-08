package records

import (
	"github.com/sky-uk/go-infoblox/api"
	"net/http"
	"strings"
)

// GetAllTXTRecordsAPI base object.
type GetAllTXTRecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllTXTRecords returns a new object of GetAllTXTRecordsAPI.
func NewGetAllTXTRecords(fields []string) *GetAllTXTRecordsAPI {
	returnFields := ""
	if fields != nil {
		returnFields = "?_return_fields=" + strings.Join(fields, ",")
	}
	this := new(GetAllTXTRecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/record:txt"+returnFields, nil, new([]TXTRecord))
	return this
}

// GetResponse returns ResponseObject of GetAllTXTRecordsAPI.
func (ga GetAllTXTRecordsAPI) GetResponse() *[]TXTRecord {
	return ga.ResponseObject().(*[]TXTRecord)
}
