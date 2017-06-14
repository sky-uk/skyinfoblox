package records

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// CreateRecordAPI base object.
type CreateRecordAPI struct {
	*api.BaseAPI
}

// CreateRecordAPI base object.
type CreateARecordAPI struct {
	*api.BaseAPI
}

// NewCreateRecord returns a new object of CreateRecordAPI.
func NewCreateRecord(recordType string, requestPayload GenericRecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/record:"+recordType, requestPayload, nil)
	return this
}

// NewCreateARecord - Creates a new A record
func NewCreateARecord(requestPayload ARecord) *CreateRecordAPI {
	this := new(CreateRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/wapi/v2.3.1/record:a", requestPayload, new(ARecord))
	return this
}

// GetResponse returns ResponseObject of CreateRecordAPI.
func (c CreateARecordAPI) GetResponse() ARecord {
	return c.ResponseObject().(ARecord)
}

// GetResponse returns ResponseObject of CreateRecordAPI.
func (c CreateRecordAPI) GetResponse() string {
	return c.ResponseObject().(string)
}
