package network

import (
	"encoding/json"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getAllNetworksSetup() *GetAllNetworksAPI {
	var fields []string
	return NewGetAllNetworks(fields)
}

func TestGetAllNetworksMethod(t *testing.T) {
	allNetworks := getAllNetworksSetup()
	assert.Equal(t, http.MethodGet, allNetworks.Method())
}

func TestGetAllNetworksEndpoint(t *testing.T) {
	GetAllNetworksAPI := getAllNetworksSetup()
	assert.Equal(t, "/wapi/v2.3.1/network", GetAllNetworksAPI.Endpoint())
}

func TestGetAllNetworksUnmarshalling(t *testing.T) {
	GetNetworkAPI := getAllNetworksSetup()
	GetNetworkAPI.SetStatusCode(http.StatusOK)
	net1 := Network{
		Ref:         "network/foo:10.10.10.0/24/default",
		Network:     "10.10.10.0/24",
		NetworkView: "default",
	}
	net2 := Network{
		Ref:         "network/bar:10.10.11.0/24/default",
		Network:     "10.10.11.0/24",
		NetworkView: "default",
	}

	networks := []Network{net1, net2}

	GetNetworkAPI.SetResponseObject(&networks)
	assert.Equal(t, networks, GetNetworkAPI.GetResponse().([]Network))
}

func TestGetAllNetworksUnmarshallingError(t *testing.T) {
	GetNetworkAPI := getAllNetworksSetup()
	GetNetworkAPI.SetStatusCode(http.StatusNotFound)

	errorString := `
{
  "Error": "AdmConProtoError: Unknown object type (fo)",
  "code": "Client.Ibap.Proto",
  "text": "Unknown object type (fo)"
}`
	GetNetworkAPI.SetRawResponse([]byte(errorString))
	var errStruct api.RespError
	err := json.Unmarshal(GetNetworkAPI.RawResponse(), &errStruct)
	if err == nil {
		assert.Equal(t, errStruct, GetNetworkAPI.GetResponse().(api.RespError), "Got same error structure back...")
	}
}
