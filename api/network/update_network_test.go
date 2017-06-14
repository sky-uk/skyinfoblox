package network

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func updateNetworkSetup() *UpdateNetworkAPI {
	netToUpdate := Network{
		Ref:     "network/ZG5zLm5ldHdvcmskMTkyLjE2OC4xLjAvMjQvMA",
		Comment: "new comment",
	}
	return NewUpdateNetwork(netToUpdate)
}

func TestUpdateNetworkMethod(t *testing.T) {
	NetToUpdate := updateNetworkSetup()
	assert.Equal(t, http.MethodPut, NetToUpdate.Method())
}

func TestUpdateNetworkEndpoint(t *testing.T) {
	NetToUpdate := updateNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network/ZG5zLm5ldHdvcmskMTkyLjE2OC4xLjAvMjQvMA", NetToUpdate.Endpoint())
}

func TestUpdateNetworkUnmarshalling(t *testing.T) {
	NetTOUpdate := updateNetworkSetup()
	resp := "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default"
	NetTOUpdate.SetResponseObject(&resp)
	resp = NetTOUpdate.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
}

func TestUpdateNetworkUnmarshallingError(t *testing.T) {
	NetTOUpdate := updateNetworkSetup()
	errorStr := `
{
	"Error": "AdmConProtoError: Field is not allowed for update: network_view",
	"code": "Client.Ibap.Proto",
	"text": "Field is not allowed for update: network_view"
}`
	NetTOUpdate.SetResponseObject(&errorStr)
	resp := NetTOUpdate.GetResponse()
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
