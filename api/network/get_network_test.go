package network

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getNetworkSetup() *GetNetworkAPI {
	objRef := "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w"
	fields := []string{"comment"}
	return NewGetNetwork(objRef, fields)
}

func TestGetNetworkMethod(t *testing.T) {
	NewNetwork := getNetworkSetup()
	assert.Equal(t, http.MethodGet, NewNetwork.Method())
}

func TestGetNetworkEndpoint(t *testing.T) {
	GetNetworkAPI := getNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w?_return_fields=comment", GetNetworkAPI.Endpoint())
}

func TestGetNetworkUnmarshalling(t *testing.T) {
	GetNetworkAPI := getNetworkSetup()
	GetNetworkAPI.SetStatusCode(http.StatusOK)
	responseObject := Network{
		Ref:         "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default",
		Network:     "10.10.10.0/24",
		NetworkView: "default",
	}
	GetNetworkAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetNetworkAPI.GetResponse())
}
