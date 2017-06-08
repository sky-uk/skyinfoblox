package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getAllARecordsSetup() *GetAllARecordsAPI {
	returnFields := []string{"ipv4addr", "name"}
	return NewGetAllARecords(returnFields)
}

func TestGetAllARecordsMethod(t *testing.T) {
	GetAllARecordsAPI := getAllARecordsSetup()
	assert.Equal(t, http.MethodGet, GetAllARecordsAPI.Method())
}

func TestGetAllARecordsEndpoint(t *testing.T) {
	GetAllARecordsAPI := getAllARecordsSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:a?_return_fields=ipv4addr,name", GetAllARecordsAPI.Endpoint())
}

func TestGetAllARecordsUnmarshalling(t *testing.T) {
	GetAllARecordsAPI := getAllARecordsSetup()
	responseObject := []ARecord{{Ref: "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlndGVzdCwxMC4xMC4xMC4x:craigtest.test-ovp.bskyb.com/default", IPv4: "10.10.10.1", Name: "craigtest.test-ovp.bskyb.com"}, {Ref: "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnMnRlc3QsMTAuMTAuMTAuMg:craig2test.test-ovp.bskyb.com/default", IPv4: "10.10.10.2", Name: "craig2test.test-ovp.bskyb.com"}}
	//responseString := []string{"record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnMnRlc3QsMTAuMTAuMTAuMg:craig2test.test-ovp.bskyb.com/default 10.10.10.2 craig2test.test-ovp.bskyb.com","record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlndGVzdCwxMC4xMC4xMC4x:craigtest.test-ovp.bskyb.com/default 10.10.10.1 craigtest.test-ovp.bskyb.com"}
	GetAllARecordsAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetAllARecordsAPI.GetResponse())
}
