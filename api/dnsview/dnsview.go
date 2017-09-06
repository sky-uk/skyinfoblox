package dnsview

import (
	"github.com/sky-uk/skyinfoblox/api"
	"net/http"
	"strings"
)

//NewCreate : used to create a new DNSView object
func NewCreate(dnsView DNSView) *api.BaseAPI {
	createDnsViewAPI := api.NewBaseAPI(http.MethodPost, wapiVersion+dnsViewEndpoint, dnsView, new(string))
	return createDnsViewAPI
}

//NewUpdate : used to update an existing DNSView object
func NewUpdate(dnsView DNSView, returnFieldList []string) *api.BaseAPI {
	reference := dnsView.Reference + "?_return_fields=" + strings.Join(returnFieldList, ",")
	createDnsViewAPI := api.NewBaseAPI(http.MethodPut, wapiVersion+dnsViewEndpoint+reference, dnsView, new(DNSView))
	return createDnsViewAPI
}

//NewGetAll : used to retrieve all DNSView objects
func NewGetAll() *api.BaseAPI {
	getAllDnsViewAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+dnsViewEndpoint, nil, new([]DNSView))
	return getAllDnsViewAPI
}

//NewGet : used to retrieve a DNSView object
func NewGet(reference string, returnFieldList []string) *api.BaseAPI {
	reference += "?_return_fields=" + strings.Join(returnFieldList, ",")
	getDnsViewAPI := api.NewBaseAPI(http.MethodGet, wapiVersion+dnsViewEndpoint+reference, nil, new(DNSView))
	return getDnsViewAPI
}

//NewDelete : used to delete a DNSView object
func NewDelete(reference string) *api.BaseAPI {
	deleteDnsViewAPI := api.NewBaseAPI(http.MethodDelete, wapiVersion+dnsViewEndpoint+reference, nil, new(string))
	return deleteDnsViewAPI
}
