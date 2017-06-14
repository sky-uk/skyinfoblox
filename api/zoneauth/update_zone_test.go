package zoneauth

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateZoneAuthAPI *UpdateZoneAuthAPI
var updateDNSRef string
var updateDNSZone DNSZone

func setupTestUpdateZoneAuth() {
	updateDNSRef = "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default"
	updateDNSZone = DNSZone{Comment: "An updated comment", Reference: updateDNSRef}
	updateZoneAuthAPI = NewUpdate(updateDNSZone)
	updateZoneAuthAPI.SetResponseObject(&updateDNSZone)
}

func TestUpdateZoneAuthMethod(t *testing.T) {
	setupTestUpdateZoneAuth()
	assert.Equal(t, http.MethodPut, updateZoneAuthAPI.Method())
}

func TestUpdateZoneAuthEndpoint(t *testing.T) {
	setupTestUpdateZoneAuth()
	assert.Equal(t, "/wapi/v2.3.1/"+updateDNSRef+"?_return_fields=fqdn,view,comment", updateZoneAuthAPI.Endpoint())
}

func TestUpdateZoneAuthMarshalling(t *testing.T) {
	setupTestUpdateZoneAuth()
	expectedJSON := `{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default","comment":"An updated comment","scavenging_settings":{"scavenging_schedule":{}}}`
	jsonBytes, err := json.Marshal(updateZoneAuthAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestUpdateZoneAuthGetResponse(t *testing.T) {
	setupTestUpdateZoneAuth()
	getResponse := updateZoneAuthAPI.GetResponse()
	assert.Equal(t, updateDNSZone.Comment, getResponse.Comment)
	assert.Equal(t, updateDNSZone.Reference, getResponse.Reference)
}
