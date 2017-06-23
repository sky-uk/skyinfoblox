package records

import (
	"encoding/json"
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
func (ga GetAllCNAMERecordsAPI) GetResponse() interface{} {

	if ga.StatusCode() == http.StatusOK {
		return *ga.ResponseObject().(*[]CNAMERecord)
	}

	var errStruct api.RespError
	err := json.Unmarshal(ga.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
