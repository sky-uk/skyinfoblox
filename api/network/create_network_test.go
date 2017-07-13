package network

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createNetworkSetup() *CreateNetworkAPI {
	members := []Member{
		{ElementType: "dhcpmember", Name: "infoblox1.example.com", IPv4Address: "192.168.0.1"},
		{ElementType: "dhcpmember", Name: "infoblox2.example.com", IPv4Address: "192.168.0.2"},
	}
	net := Network{
		Network:     "10.10.10.1/24",
		NetworkView: "default",
		Comment:     "Test network",
		Members:     members,
	}
	newCreateNetworkAPI := NewCreateNetwork(net)
	responseString := "dummy response"
	newCreateNetworkAPI.SetResponseObject(&responseString)
	return newCreateNetworkAPI
}

func TestCreateNetworkMethod(t *testing.T) {
	NewNetwork := createNetworkSetup()
	assert.Equal(t, http.MethodPost, NewNetwork.Method())
}

func TestCreateNetworkEndpoint(t *testing.T) {
	NewNetwork := createNetworkSetup()
	assert.Equal(t, "/wapi/v2.3.1/network", NewNetwork.Endpoint())
}

func TestCreateNetworkUnmarshalling(t *testing.T) {
	NewNetwork := createNetworkSetup()
	response := "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default"
	NewNetwork.SetResponseObject(&response)
	resp := NewNetwork.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:10.10.10.0/24/default")
}

func TestCreateNetworkMarshalling(t *testing.T) {
	NewNetwork := createNetworkSetup()
	expectedJSON := "{\"_ref\":\"\",\"network\":\"10.10.10.1/24\",\"network_view\":\"default\",\"comment\":\"Test network\",\"members\":[{\"_struct\":\"dhcpmember\",\"ipv4addr\":\"192.168.0.1\",\"ipv6addr\":\"\",\"name\":\"infoblox1.example.com\"},{\"_struct\":\"dhcpmember\",\"ipv4addr\":\"192.168.0.2\",\"ipv6addr\":\"\",\"name\":\"infoblox2.example.com\"}]}"
	jsonBytes, err := json.Marshal(NewNetwork.RequestObject().(Network))
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestGetResponse(t *testing.T) {
	NewNetwork := createNetworkSetup()
	resp := NewNetwork.GetResponse()
	assert.Equal(t, resp, "dummy response")
}
