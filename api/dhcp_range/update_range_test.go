package dhcprange

import (
	//	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func UpdateDHCPRangeSetup() *UpdateDHCPRangeAPI {
	dhcpRange := DHCPRange{
		Network:     "192.168.0.0/24",
		NetworkView: "default",
		Start:       "192.168.0.50",
		End:         "192.168.0.60",
		Restart:     true,
	}
	dhcpMember := Member{ElementType: "dhcpmember", Name: "test.example.com", IPv4Address: "192.168.0.10"}
	dhcpRange.Ref = "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default"
	dhcpRange.Member = dhcpMember
	dhcpRange.ServerAssociation = "MEMBER"
	newUpdateRangeAPI := NewUpdateDHCPRange(dhcpRange)
	responseString := "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default"
	newUpdateRangeAPI.SetResponseObject(&responseString)
	return newUpdateRangeAPI
}

func TestMethodDHCPRange(t *testing.T) {
	testRelay := UpdateDHCPRangeSetup()
	assert.Equal(t, http.MethodPut, testRelay.Method())

}

func TestEndpointDHCPRange(t *testing.T) {
	testRelay := UpdateDHCPRangeSetup()
	assert.Equal(t, "/wapi/v2.3.1/range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default", testRelay.Endpoint())

}

func TestNewUpdateDHCPRangeResponse(t *testing.T) {
	testRelay := UpdateDHCPRangeSetup()
	assert.Equal(t, "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w:192.168.0.0/24/default", testRelay.GetResponse())

}
