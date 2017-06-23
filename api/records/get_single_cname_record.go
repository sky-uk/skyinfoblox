package records

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// GetSingleCNAMERecordAPI base object.
type GetSingleCNAMERecordAPI struct {
	*api.BaseAPI
}

// NewGetCNAMERecord returns a new object of GetSingleCNAMERecordAPI.
func NewGetCNAMERecord(recordReference string, returnFields []string) *GetSingleCNAMERecordAPI {
	if returnFields != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFields, ",")
		recordReference += returnFields
	}
	this := new(GetSingleCNAMERecordAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/wapi/v2.3.1/"+recordReference, nil, new(CNAMERecord))
	return this
}

// GetResponse returns ResponseObject of GetSingleCNAMERecordAPI.
func (gs GetSingleCNAMERecordAPI) GetResponse() interface{} {
	if gs.StatusCode() == http.StatusOK {
		return *gs.ResponseObject().(*CNAMERecord)
	}

	var errStruct api.RespError
	err := json.Unmarshal(gs.RawResponse(), &errStruct)
	if err != nil {
		return nil
	}
	return errStruct
}
