package network

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDeleteNetworkMethod(t *testing.T) {
	client := NewDeleteNetwork("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	assert.Equal(t, http.MethodDelete, client.Method())
}

func TestDeleteNetworkEndpoint(t *testing.T) {
	client := NewDeleteNetwork("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	assert.Equal(t, "/wapi/v2.3.1/network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default", client.Endpoint())
}

func TestDeleteNetworkUnmarshalling(t *testing.T) {
	client := NewDeleteNetwork("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	client.SetResponseObject("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	resp := client.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
}
