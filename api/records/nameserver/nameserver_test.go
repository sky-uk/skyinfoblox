package nameserver

import (
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var createNSRecordAPI, getAllNSRecordAPI, getNSRecordAPI, updateNSRecordAPI, deleteNSRecordAPI *api.BaseAPI
var returnFields []string
var reference string

func setupNameServerRecordTest() {

	reference = "record:ns/ZG5zLmJpbmRfbnMkLl9kZWZhdWx0LmNvbS5ic2t5Yi5zbHVwYWFzLi5ub25wcmRpYnhkbnMwNC5ic2t5Yi5jb20:ns1.example.com/example.com/default"
	nsServerAutoCreatePointerRecord := true
	zoneNameServer := ZoneNameServer{"192.168.0.1", &nsServerAutoCreatePointerRecord}
	zoneNameServers := make([]ZoneNameServer, 0)
	zoneNameServers = append(zoneNameServers, zoneNameServer)
	nameServerRecordObject := NSRecord{
		Reference:  reference,
		Name:       "example.com",
		Addresses:  zoneNameServers,
		NameServer: "ns1.example.com",
		View:       "default",
	}
	returnFields = []string{"name", "addresses", "nameserver", "view", "_ref"}

	adminGroupList := make([]NSRecord, 0)
	adminGroupList = append(adminGroupList, nameServerRecordObject)

	createNSRecordAPI = NewCreate(nameServerRecordObject)
	createNSRecordAPI.SetResponseObject(&reference)

	getAllNSRecordAPI = NewGetAll()
	getAllNSRecordAPI.SetResponseObject(&adminGroupList)

	getNSRecordAPI = NewGet(reference, returnFields)
	getNSRecordAPI.SetResponseObject(&nameServerRecordObject)

	updateNSRecordAPI = NewUpdate(nameServerRecordObject, returnFields)
	updateNSRecordAPI.SetResponseObject(&nameServerRecordObject)

	deleteNSRecordAPI = NewDelete(reference)
}

func TestNameServerRecordNewCreateMethod(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, http.MethodPost, createNSRecordAPI.Method())
}

func TestNameServerRecordNewCreateEndpoint(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, wapiVersion+nsEndpoint, createNSRecordAPI.Endpoint())
}

func TestNameServerRecordNewCreateResponse(t *testing.T) {
	setupNameServerRecordTest()
	response := *createNSRecordAPI.ResponseObject().(*string)

	assert.Equal(t, reference, response)
}

func TestNameServerRecordNewGetAllMethod(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, http.MethodGet, getAllNSRecordAPI.Method())
}

func TestNameServerRecordNewGetAllEndpoint(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, wapiVersion+nsEndpoint, getAllNSRecordAPI.Endpoint())
}

func TestNameServerRecordNewGetAllResponse(t *testing.T) {
	setupNameServerRecordTest()
	response := *getAllNSRecordAPI.ResponseObject().(*[]NSRecord)

	assert.Equal(t, "record:ns/ZG5zLmJpbmRfbnMkLl9kZWZhdWx0LmNvbS5ic2t5Yi5zbHVwYWFzLi5ub25wcmRpYnhkbnMwNC5ic2t5Yi5jb20:ns1.example.com/example.com/default", response[0].Reference)
	assert.Equal(t, "example.com", response[0].Name)
	assert.Equal(t, "default", response[0].View)
	assert.Equal(t, "ns1.example.com", response[0].NameServer)
	assert.Equal(t, "192.168.0.1", response[0].Addresses[0].Address)
	assert.Equal(t, true, *response[0].Addresses[0].AutoCreatePointerRecord)
}

func TestNameServerRecordNewGetMethod(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, http.MethodGet, getNSRecordAPI.Method())
}

func TestNameServerRecordNewGetEndpoint(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, wapiVersion+"/"+reference+"?_return_fields="+strings.Join(returnFields, ","), getNSRecordAPI.Endpoint())
}

func TestNameServerRecordNewGetResponse(t *testing.T) {
	setupNameServerRecordTest()
	response := *getNSRecordAPI.ResponseObject().(*NSRecord)

	assert.Equal(t, "example.com", response.Name)
	assert.Equal(t, "default", response.View)
	assert.Equal(t, "ns1.example.com", response.NameServer)
	assert.Equal(t, "192.168.0.1", response.Addresses[0].Address)
	assert.Equal(t, true, *response.Addresses[0].AutoCreatePointerRecord)
}

func TestNameServerRecordNewUpdateMethod(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, http.MethodPut, updateNSRecordAPI.Method())
}

func TestNameServerRecordNewUpdateEndpoint(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, wapiVersion+"/"+reference+"?_return_fields="+strings.Join(returnFields, ","), updateNSRecordAPI.Endpoint())
}

func TestNameServerRecordNewUpdateResponse(t *testing.T) {
	setupNameServerRecordTest()
	response := updateNSRecordAPI.ResponseObject().(*NSRecord)

	assert.Equal(t, "example.com", response.Name)
	assert.Equal(t, "default", response.View)
	assert.Equal(t, "ns1.example.com", response.NameServer)
	assert.Equal(t, "192.168.0.1", response.Addresses[0].Address)
	assert.Equal(t, true, *response.Addresses[0].AutoCreatePointerRecord)
}

func TestNameServerRecordNewDeleteMethod(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, http.MethodDelete, deleteNSRecordAPI.Method())
}

func TestNameServerRecordNewDeleteEndpoint(t *testing.T) {
	setupNameServerRecordTest()
	assert.Equal(t, wapiVersion+"/"+reference, deleteNSRecordAPI.Endpoint())
}
