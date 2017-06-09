package records

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
)

// DeleteRecordAPI base object.
type DeleteRecordAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new object of DeleteRecordAPI.
func NewDelete(recordReference string) *DeleteRecordAPI {
	this := new(DeleteRecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/wapi/v2.3.1/"+recordReference, nil, new(string))
	return this
}

// GetResponse returns ResponseObject of DeleteRecordAPI.
func (d DeleteRecordAPI) GetResponse() string {
	return *d.ResponseObject().(*string)
}
