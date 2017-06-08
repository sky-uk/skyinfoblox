package network

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

var getAllNetworksAPI *GetAllNetworksAPI

func setupGetAll() {
	getAllNetworksAPI = NewGetAllNetworks()
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllNetworksAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/wapi/v2.3.1/network", getAllNetworksAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()

	jsonContent, err := ioutil.ReadFile("../../httpTemplates/get_all_networks.json")
	if err != nil {
		t.Fatalf("Unable to read HTTP template: %v", err)
	}

	jsonErr := json.Unmarshal([]byte(jsonContent), getAllNetworksAPI.ResponseObject())
	assert.Nil(t, jsonErr)
	t.Logf("Response Object: %+v", getAllNetworksAPI.ResponseObject())

	/*----
	fmt.Println(getAllNetworksAPI.GetResponse().Children[0].Name)
	assert.Nil(t, jsonErr)
	assert.Len(t, getAllNetworksAPI.GetResponse().Children, 5)
	assert.Equal(t, "PaaSExampleHTTPvirtualserver", getAllNetworksAPI.GetResponse().Children[0].Name)
	assert.Equal(
		t,
		"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver",
		getAllNetworksAPI.GetResponse().Children[0].Href,
	)
	assert.Equal(
		t,
		"PaaSExampleHTTPvirtualserver1",
		getAllNetworksAPI.GetResponse().Children[1].Name,
	)
	assert.Equal(
		t,
		"/api/tm/3.8/config/active/virtual_servers/PaaSExampleHTTPvirtualserver1",
		getAllNetworksAPI.GetResponse().Children[1].Href,
	)
	-----*/
}

func TestGetAllNetworks_GetResponse(t *testing.T) {
	setupGetAll()
	assert.IsType(t, getAllNetworksAPI.ResponseObject(), &[]Network{})
}
