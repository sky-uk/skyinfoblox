package network

import (
	"fmt"
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
	assert.Equal(t, fmt.Sprintf("%s/network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w?_return_fields=comment", wapiVersion), GetNetworkAPI.Endpoint())
}

func TestGetNetworkUnmarshalling(t *testing.T) {
	GetNetworkAPI := getNetworkSetup()
	responseObject := Network{
		Ref:         "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default",
		Network:     "10.10.10.0/24",
		NetworkView: "default",
	}
	GetNetworkAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetNetworkAPI.GetResponse())
}
