package network

import (
	"fmt"
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
	assert.Equal(t, fmt.Sprintf("%s/network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default", wapiVersion), client.Endpoint())
}

func TestDeleteNetworkUnmarshalling(t *testing.T) {
	client := NewDeleteNetwork("network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	r := "network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default"
	client.SetResponseObject(&r)
	resp := client.GetResponse()
	assert.Equal(t, resp, "network/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
}
