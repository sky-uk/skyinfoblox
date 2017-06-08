package records

import (
	"github.com/sky-uk/go-infoblox/api"
	"net/http"
)

// CreateRecordAPI base object.
type CreateRecordAPI struct {
	*api.BaseAPI
}

// NewCreateRecord returns a new object of CreateRecordAPI.
func NewCreateRecord(recordType string, requestPayload GenericRecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/record:"+recordType, requestPayload, new(string))
	return this
}

// GetResponse returns ResponseObject of CreateRecordAPI.
func (c CreateRecordAPI) GetResponse() string {
	return *c.ResponseObject().(*string)
}
