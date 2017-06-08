package records

import (
	"github.com/sky-uk/go-infoblox/api"
	"net/http"
	"strings"
)

// GetRecordAPI base object.
type GetRecordAPI struct {
	*api.BaseAPI
}


// NewGetRecord returns a new object of GetRecordAPI.
func NewGetRecord(recordReference string, returnFields []string) *GetRecordAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		recordReference += returnFields
	}
	this := new(GetRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+recordReference, nil, new(GenericRecord))
	return this
}

// GetResponse returns ResponseObject of GetRecordAPI.
func (gs GetRecordAPI) GetResponse() *GenericRecord {
	return gs.ResponseObject().(*GenericRecord)
}