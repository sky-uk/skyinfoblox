package records

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetSingleTXTRecordAPI base object.
type GetSingleTXTRecordAPI struct {
	*api.BaseAPI
}

// NewGetTXTRecord returns a new object of GetSingleTXTRecordAPI.
func NewGetTXTRecord(recordReference string, returnFields []string) *GetSingleTXTRecordAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		recordReference += returnFields
	}
	this := new(GetSingleTXTRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+recordReference, nil, new(TXTRecord))
	return this
}

// GetResponse returns ResponseObject of GetSingleTXTRecordAPI.
func (gs GetSingleTXTRecordAPI) GetResponse() TXTRecord {
	return *gs.ResponseObject().(*TXTRecord)
}
