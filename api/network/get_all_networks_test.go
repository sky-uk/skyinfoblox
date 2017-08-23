package network

import (
	"fmt"
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
	assert.Equal(t, fmt.Sprintf("%s/network", wapiVersion), GetAllNetworksAPI.Endpoint())
}

func TestGetAllNetworksUnmarshalling(t *testing.T) {
	GetNetworkAPI := getAllNetworksSetup()
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
	assert.Equal(t, networks, GetNetworkAPI.GetResponse())
}
