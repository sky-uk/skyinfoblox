package zoneauth

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllZoneAuthAPI *GetAllZoneAuthAPI

func setupReadAllZones() {
	getAllZoneAuthAPI = NewGetAllZones()
}

func TestGetAllMethod(t *testing.T) {
	setupReadAllZones()
	assert.Equal(t, http.MethodGet, getAllZoneAuthAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupReadAllZones()
	assert.Equal(t, "/wapi/v2.3.1/zone_auth?_return_fields=fqdn", getAllZoneAuthAPI.Endpoint())
}

func TestGetAllZonesUnmarshalling(t *testing.T) {
	setupReadAllZones()
	jsonContent := []byte(`[{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmFycGEuaW4tYWRkci4xMC4xMC4xMA:10.10.10.0%2F24/default","fqdn":"10.10.10.0/24"},{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default","fqdn":"np.ovp.sky.com"},{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cA:test-ovp.bskyb.com/default","fqdn":"test-ovp.bskyb.com"},{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi5vdnA:ovp.bskyb.com/default","fqdn":"ovp.bskyb.com"},{"_ref":"zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi5vdnAudGVzdA:test.ovp.bskyb.com/default","fqdn":"test.ovp.bskyb.com"}]`)
	jsonErr := json.Unmarshal(jsonContent, getAllZoneAuthAPI.ResponseObject())

	getAllZoneAuthAPIResponse := *getAllZoneAuthAPI.GetResponse()

	assert.Nil(t, jsonErr)
	assert.Len(t, getAllZoneAuthAPIResponse, 5)
	assert.Equal(t, "10.10.10.0/24", getAllZoneAuthAPIResponse[0].FQDN)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmFycGEuaW4tYWRkci4xMC4xMC4xMA:10.10.10.0%2F24/default", getAllZoneAuthAPIResponse[0].Reference)
	assert.Equal(t, "np.ovp.sky.com", getAllZoneAuthAPIResponse[1].FQDN)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5za3kub3ZwLm5w:np.ovp.sky.com/default", getAllZoneAuthAPIResponse[1].Reference)
	assert.Equal(t, "test-ovp.bskyb.com", getAllZoneAuthAPIResponse[2].FQDN)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cA:test-ovp.bskyb.com/default", getAllZoneAuthAPIResponse[2].Reference)
	assert.Equal(t, "ovp.bskyb.com", getAllZoneAuthAPIResponse[3].FQDN)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi5vdnA:ovp.bskyb.com/default", getAllZoneAuthAPIResponse[3].Reference)
	assert.Equal(t, "test.ovp.bskyb.com", getAllZoneAuthAPIResponse[4].FQDN)
	assert.Equal(t, "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi5vdnAudGVzdA:test.ovp.bskyb.com/default", getAllZoneAuthAPIResponse[4].Reference)
}
