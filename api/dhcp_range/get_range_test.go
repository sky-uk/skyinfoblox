package dhcprange

import (
	"encoding/json"
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

func TestGetDHCPRangeMarshalling(t *testing.T) {
	allowRestart := true
	requestData := DHCPRange{
		Network:     "192.168.0.0/24",
		NetworkView: "default",
		Start:       "192.168.0.50",
		End:         "192.168.0.60",
		Restart:     &allowRestart,
	}
	dhcpMember := Member{ElementType: "dhcpmember", Name: "test.example.com", IPv4Address: "192.168.0.10"}
	requestData.Member = dhcpMember
	requestData.ServerAssociation = "MEMBER"

	expectedJSON := `{"_ref":"","start_addr":"192.168.0.50","end_addr":"192.168.0.60","network":"192.168.0.0/24","network_view":"default","restart_if_needed":true,"server_association_type":"MEMBER","member":{"_struct":"dhcpmember","ipv4addr":"192.168.0.10","name":"test.example.com"}}`
	jsonBytes, err := json.Marshal(requestData)
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))

}

func TestGetDHCPRangeUnmarshalling(t *testing.T) {
	responseBytes := []byte(`{"_ref":"","start_addr":"192.168.0.50","end_addr":"192.168.0.60","network":"192.168.0.0/24","network_view":"default","restart_if_needed":true,"server_association_type":"MEMBER","member":{"_struct":"dhcpmember","ipv4addr":"192.168.0.10","name":"test.example.com"}}`)
	responseObject := new(DHCPRange)

	err := json.Unmarshal(responseBytes, responseObject)
	assert.Nil(t, err)
	assert.NotNil(t, responseObject)
	assert.Equal(t, "192.168.0.0/24", responseObject.Network)
	assert.Equal(t, "default", responseObject.NetworkView)
	assert.Equal(t, "192.168.0.50", responseObject.Start)
	assert.Equal(t, "192.168.0.60", responseObject.End)

}
