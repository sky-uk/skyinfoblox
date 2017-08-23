package zoneauth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteZoneAuthAPI *DeleteZoneAuthAPI
var zoneReference = "zone_auth/ZG5zLnpvbmUkLl9kZWZhdWx0LmNvbS5ic2t5Yi5vdnAucGFhcy50ZXN0aW5n:testing.paas.ovp.bskyb.com/default"

func setupDelete() {
	deleteZoneAuthAPI = NewDelete(zoneReference)
	deleteZoneAuthAPI.SetResponseObject(&zoneReference)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteZoneAuthAPI.Method())
}

func TestDeleteEndpoing(t *testing.T) {
	setupDelete()
	assert.Equal(t, fmt.Sprintf("%s/%s", wapiVersion, zoneReference), deleteZoneAuthAPI.Endpoint())
}

func TestDeleteResponseObject(t *testing.T) {
	setupDelete()
	getResponse := deleteZoneAuthAPI.GetResponse()
	assert.Equal(t, zoneReference, getResponse)
}
