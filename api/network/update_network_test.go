package network

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func updateNetworkSetup() *UpdateNetworkAPI {
	return NewUpdateNetwork("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4xLjAvMjQvMA", "comment=foo")
}

func TestUpdateNetworkMethod(t *testing.T) {
	NewNetwork := updateNetworkSetup()
	assert.Equal(t, http.MethodPut, NewNetwork.Method())
}

func TestUpdateNetworkEndpoint(t *testing.T) {
	NewNetwork := updateNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network/ZG5zLm5ldHdvcmskMTkyLjE2OC4xLjAvMjQvMA", NewNetwork.Endpoint())
}

func TestUpdateNetworkUnmarshalling(t *testing.T) {
	NewNetwork := updateNetworkSetup()
	NewNetwork.SetResponseObject("network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
	resp := NewNetwork.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
}

func TestUpdateNetworkUnmarshallingError(t *testing.T) {
	NewNetwork := updateNetworkSetup()
	NewNetwork.SetResponseObject(`
{
	"Error": "AdmConProtoError: Field is not allowed for update: network_view",
	"code": "Client.Ibap.Proto",
	"text": "Field is not allowed for update: network_view"
}`)
	resp := NewNetwork.GetResponse()
	errorObj := new(skyinfoblox.RespError)
	errStr := []byte(resp)
	JSONerr := json.Unmarshal(errStr, errorObj)
	if JSONerr != nil {
		t.Error("Error decoding response: ", JSONerr)
	} else {
		t.Log("Got proper error code :", errorObj.Code)
		assert.Equal(t, errorObj.Code, "Client.Ibap.Proto")
	}
}
