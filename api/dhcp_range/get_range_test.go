package dhcprange

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getDHCPRangeSetup() *GetDHCPRangeAPI {
	objRef := "range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w"
	fields := []string{"end_addr", "start_addr", "network", "network_view", "member"}
	return NewGetDHCPRangeAPI(objRef, fields)
}

func TestGetDHCPRangeMethod(t *testing.T) {
	NewDHCPRangeReq := getDHCPRangeSetup()
	assert.Equal(t, http.MethodGet, NewDHCPRangeReq.Method())
}

func TestGetDHCPRangeEndpoint(t *testing.T) {
	GetDHCPRangeReq := getDHCPRangeSetup()
	assert.Equal(t, "/wapi/v2.3.1/range/ZG5zLm5ldHdvcmskMTAuMTAuMTAuMC8yNC8w?_return_fields=end_addr,start_addr,network,network_view,member", GetDHCPRangeReq.Endpoint())
}

func TestGetDHCPRangeUnmarshalling(t *testing.T) {
	GetDHCPRangeReq := getDHCPRangeSetup()
	responseObject := DHCPRange{
		Network:     "192.168.0.0/24",
		NetworkView: "default",
		Start:       "192.168.0.50",
		End:         "192.168.0.60",
		Restart:     true,
	}
	dhcpMember := Member{InternalType: "dhcpmember", Name: "test.example.com", Address: "192.168.0.10"}
	responseObject.Member = dhcpMember
	responseObject.ServerAssociation = "MEMBER"
	GetDHCPRangeReq.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetDHCPRangeReq.GetResponse())
}
