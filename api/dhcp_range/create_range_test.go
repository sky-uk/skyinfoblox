package dhcprange

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createRangeSetup() *CreateDHCPRangeAPI {
	dhcpRange := DHCPRange{
		Network:     "192.168.0.0/24",
		NetworkView: "default",
		Start:       "192.168.0.50",
		End:         "192.168.0.60",
		Restart:     true,
	}
	dhcpMember := Member{ElementType: "dhcpmember", Name: "test.example.com", IPv4Address: "192.168.0.10"}
	dhcpRange.Member = dhcpMember
	dhcpRange.ServerAssociation = "MEMBER"
	newCreateDHCPRangeAPI := NewCreateDHCPRange(dhcpRange)
	responseString := "dummy response"
	newCreateDHCPRangeAPI.SetResponseObject(&responseString)
	return newCreateDHCPRangeAPI
}

func TestCreateNetworkMethod(t *testing.T) {
	NewDHCPRangeReq := createRangeSetup()
	assert.Equal(t, http.MethodPost, NewDHCPRangeReq.Method())
}

func TestCreateNetworkEndpoint(t *testing.T) {
	NewDHCPRangeReq := createRangeSetup()
	assert.Equal(t, "/wapi/v2.3.1/range", NewDHCPRangeReq.Endpoint())
}

func TestCreateNetworkUnmarshalling(t *testing.T) {
	NewDHCPRangeReq := createRangeSetup()
	response := "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default"
	NewDHCPRangeReq.SetResponseObject(&response)
	resp := NewDHCPRangeReq.GetResponse()
	assert.Equal(t, resp, "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default")
}

func TestGetResponse(t *testing.T) {
	NewDHCPRangeReq := createRangeSetup()
	resp := NewDHCPRangeReq.GetResponse()
	assert.Equal(t, resp, "dummy response")
}
