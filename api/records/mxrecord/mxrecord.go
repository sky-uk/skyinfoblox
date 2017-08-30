package mxrecord

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

func NewCreate(mxRecord MxRecord) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodPost, fmt.Sprintf("%s/%s", WapiVersion, MXRecordEndpoint), mxRecord, new(string))
}

func NewGet(reference string, returnFields []string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s/%s", WapiVersion, MXRecordEndpoint, reference), nil, new(MxRecord))
}

func NewGetAll() *api.BaseAPI {
	return api.NewBaseAPI(http.MethodGet, fmt.Sprintf("%s/%s", WapiVersion, MXRecordEndpoint), nil, new([]MxRecord))

}

func NewUpdate(mxRecord MxRecord, returnFields []string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodPut, fmt.Sprintf("%s/%s/%s?return_fields=%s", WapiVersion, MXRecordEndpoint, mxRecord.Ref, strings.Join(returnFields, ",")), mxRecord, new(string))
}

func NewDelete(reference string) *api.BaseAPI {
	return api.NewBaseAPI(http.MethodDelete, fmt.Sprintf("%s/%s/%s", WapiVersion, MXRecordEndpoint, reference), nil, new(string))
}
