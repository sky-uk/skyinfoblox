package zoneauth

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getSingleZoneAuthAPI *GetSingleZoneAuthAPI
var testZoneAuthReference = "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmFycGEuaW4tYWRkci4xMC4xMC4xMA:10.10.10.0%2F24/default"

func setupGetSingleZoneAuth() {
	returnFields := []string{"comment", "fqdn", "view"}
	getSingleZoneAuthAPI = NewGetSingleZone(testZoneAuthReference, returnFields)
}

func TestGetZoneAuthResponseMethod(t *testing.T) {
	setupGetSingleZoneAuth()
	assert.Equal(t, http.MethodGet, getSingleZoneAuthAPI.Method())
}

func TestGetZoneAuthEndpoint(t *testing.T) {
	setupGetSingleZoneAuth()
	assert.Equal(t, "/wapi/v2.3.1/"+testZoneAuthReference+"?_return_fields=comment,fqdn,view", getSingleZoneAuthAPI.Endpoint())
}

func TestGetZoneAuthResponse(t *testing.T) {
	setupGetSingleZoneAuth()
	getSingleZoneAuthAPI.SetStatusCode(http.StatusOK)
	jsonContent := []byte(`{"_ref": "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default","fqdn": "np.ovp.sky.com","view": "default","comment": "A Comment"}`)
	jsonErr := json.Unmarshal(jsonContent, getSingleZoneAuthAPI.ResponseObject())

	assert.Nil(t, jsonErr)
	response := getSingleZoneAuthAPI.GetResponse().(DNSZone)
	assert.Equal(t, "np.ovp.sky.com", response.FQDN)
	assert.Equal(t, "default", response.View)
	assert.Equal(t, "A Comment", response.Comment)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default", response.Reference)
}
