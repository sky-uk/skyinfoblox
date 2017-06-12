package network

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createNetworkSetup() *CreateNetworkAPI {
	net := Network{
		Network:     "10.10.10.1/24",
		NetworkView: "default",
		Comment:     "Test network",
	}
	return NewCreateNetwork(net)
}

func TestCreateNetworkMethod(t *testing.T) {
	NewNetwork := createNetworkSetup()
	assert.Equal(t, http.MethodPost, NewNetwork.Method())
}

func TestCreateNetworkEndpoint(t *testing.T) {
	NewNetwork := createNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network?network=10.10.10.1/24", NewNetwork.Endpoint())
}

func TestCreateNetworkUnmarshalling(t *testing.T) {
	NewNetwork := createNetworkSetup()
	NewNetwork.SetResponseObject("network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
	resp := NewNetwork.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
}
