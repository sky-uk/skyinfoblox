package zoneauth

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createZoneAuthAPI *CreateZoneAuthAPI

func createSetup() {
	newZone := DNSZone{FQDN: "testing.paas.bskyb.com"}
	createZoneAuthAPI = NewCreate(newZone)
	responseObject := `"zone_auth/ZG5zLnpvbmUkLjEuY29tLnNreS5vdnAudGVzdC5wYWFzMg:paas2.test.ovp.sky.com/dev"`
	createZoneAuthAPI.SetResponseObject(&responseObject)
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createZoneAuthAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, fmt.Sprintf("%s/zone_auth", wapiVersion), createZoneAuthAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := `{"fqdn":"testing.paas.bskyb.com","scavenging_settings":{"scavenging_schedule":{}}}`
	jsonBytes, err := json.Marshal(createZoneAuthAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	getResponse := createZoneAuthAPI.GetResponse()
	assert.Equal(t, `"zone_auth/ZG5zLnpvbmUkLjEuY29tLnNreS5vdnAudGVzdC5wYWFzMg:paas2.test.ovp.sky.com/dev"`, getResponse)
}
