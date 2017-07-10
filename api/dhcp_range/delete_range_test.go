package dhcprange

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDeleteDHCPRangeMethod(t *testing.T) {
	client := NewDeleteDHCPRange("range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	assert.Equal(t, http.MethodDelete, client.Method())
}

func TestDeleteDHCPRangeEndpoint(t *testing.T) {
	client := NewDeleteDHCPRange("range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	assert.Equal(t, "/wapi/v2.3.1/range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default", client.Endpoint())
}

func TestDeleteDHCPRangeUnmarshalling(t *testing.T) {
	client := NewDeleteDHCPRange("range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
	r := "range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default"
	client.SetResponseObject(&r)
	resp := client.GetResponse()
	assert.Equal(t, resp, "range/ZG5zLm5ldHdvcmskMTkyLjE2OC4wLjAvMjQvMA:192.168.0.0/24/default")
}
