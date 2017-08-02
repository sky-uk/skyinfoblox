package nameserver

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

// NewCreate : used to create a new name server record
func NewCreate(nsRecord NSRecord) *api.BaseAPI {
	createNSRecordAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+nsEndpoint, nsRecord, new(string))
	return createNSRecordAPI
}

// NewGetAll : used to get all name server records
func NewGetAll() *api.BaseAPI {
	getAllNSRecordsAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+nsEndpoint, nil, new(NSRecord))
	return getAllNSRecordsAPI
}

// NewGet : used to get a name server record
func NewGet(ref string, returnFieldList []string) *api.BaseAPI {

	if returnFieldList != nil {
		returnFields := "?_return_fields=" + strings.Join(returnFieldList, ",")
		ref += returnFields
	}
	getNSRecordAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+"/"+ref, nil, new(NSRecord))
	return getNSRecordAPI
}

// NewUpdate : used to update a name server record
func NewUpdate(nsRecord NSRecord, returnFields []string) *api.BaseAPI {

	var reference string
	if returnFields != nil {
		reference = nsRecord.Reference + "?_return_fields=" + strings.Join(returnFields, ",")
	} else {
		reference = nsRecord.Reference
	}
	updateNSRecordAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+"/"+reference, nsRecord, new(NSRecord))
	return updateNSRecordAPI
}

// NewDelete : used to delete a name server record
func NewDelete(reference string) *api.BaseAPI {
	deleteNSRecordAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+"/"+reference, nil, new(string))
	return deleteNSRecordAPI
}
