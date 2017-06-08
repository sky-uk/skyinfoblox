package zone_auth

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createZoneAuthAPI *CreateZoneAPI

func createSetup() {
	newZone := DNSZone{FQDN: "testing.paas.bskyb.com"}
	createZoneAuthAPI = NewCreate(newZone)
	createZoneAuthAPI.SetResponseObject("_ref\": \"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLnRlc3QucGFhcw:paas.test.ovp.sky.com/default\"")
}

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createZoneAuthAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/wapi/v2.3.1/zone_auth?_return_fields=fqdn", createZoneAuthAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedJSON := "{\"fqdn\":\"testing.paas.bskyb.com\"}"
	jsonBytes, err := json.Marshal(createZoneAuthAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	createSetup()
	getResponse := createZoneAuthAPI.GetResponse()
	assert.Equal(t, getResponse, "_ref\": \"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLnRlc3QucGFhcw:paas.test.ovp.sky.com/default\"")
}
