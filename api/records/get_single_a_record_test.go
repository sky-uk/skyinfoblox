package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getARecordSetup() *GetSingleARecordAPI {
	recordReference := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnM3Rlc3QsMTAuMTAuMS44MA:craig3test.test-ovp.bskyb.com/default"
	returnFields := []string{"name", "ipv4addr"}
	return NewGetARecord(recordReference, returnFields)
}

func TestGetARecordMethod(t *testing.T) {
	GetSingleARecordAPI := getARecordSetup()
	assert.Equal(t, http.MethodGet, GetSingleARecordAPI.Method())
}

func TestGetARecordEndpoint(t *testing.T) {
	GetSingleARecordAPI := getARecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnM3Rlc3QsMTAuMTAuMS44MA:craig3test.test-ovp.bskyb.com/default?_return_fields=name,ipv4addr", GetSingleARecordAPI.Endpoint())
}

func TestGetARecordUnmarshalling(t *testing.T) {
	GetSingleARecordAPI := getARecordSetup()
	GetSingleARecordAPI.SetStatusCode(http.StatusOK)
	responseObject := ARecord{Ref: "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnM3Rlc3QsMTAuMTAuMS44MA:craig3test.test-ovp.bskyb.com/default", Name: "craig3test.test-ovp.bskyb.com", IPv4: "10.10.1.80"}
	GetSingleARecordAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetSingleARecordAPI.GetResponse())
}
