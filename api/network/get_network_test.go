package network

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getNetworkSetup() *GetNetworkAPI {
	net := Network{
		Ref:         "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default",
		Network:     "10.10.10.1",
		NetworkView: "default",
		Comment:     "Test network",
	}
	return NewGetNetwork(net)
}

func TestGetNetworkMethod(t *testing.T) {
	NewNetwork := getNetworkSetup()
	assert.Equal(t, http.MethodGet, NewNetwork.Method())
}

func TestGetNetworkEndpoint(t *testing.T) {
	GetNetworkAPI := getNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default", GetNetworkAPI.Endpoint())
}

func TestGetNetworkUnmarshalling(t *testing.T) {
	GetNetworkAPI := getNetworkSetup()
	responseObject := Network{
		Ref:         "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default",
		Network:     "10.10.10.0/24",
		NetworkView: "default",
	}
	GetNetworkAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, *GetNetworkAPI.GetResponse())
}
