package records

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// GetAllARecordsAPI base object.
type GetAllARecordsAPI struct {
	*api.BaseAPI
}

// NewGetAllARecords returns a new object of GetAllARecordsAPI.
func NewGetAllARecords() *GetAllARecordsAPI {
	this := new(GetAllARecordsAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/record:a", nil, new([]ARecord))
	return this
}

// GetResponse returns ResponseObject of GetAllARecordsAPI.
func (ga GetAllARecordsAPI) GetResponse() *[]ARecord {
	return ga.ResponseObject().(*[]ARecord)
}
