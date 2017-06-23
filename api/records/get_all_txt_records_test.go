package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getAllTXTRecordsSetup() *GetAllTXTRecordsAPI {
	returnFields := []string{"name", "text"}
	return NewGetAllTXTRecords(returnFields)
}

func TestGetAllTXTRecordsMethod(t *testing.T) {
	GetAllTXTRecordsAPI := getAllTXTRecordsSetup()
	assert.Equal(t, http.MethodGet, GetAllTXTRecordsAPI.Method())
}

func TestGetAllTXTRecordsEndpoint(t *testing.T) {
	GetAllTXTRecordsAPI := getAllTXTRecordsSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:txt?_return_fields=name,text", GetAllTXTRecordsAPI.Endpoint())
}

func TestGetAllTXTRecordsUnmarshalling(t *testing.T) {
	GetAllTXTRecordsAPI := getAllTXTRecordsSetup()
	GetAllTXTRecordsAPI.SetStatusCode(http.StatusOK)
	responseObject := []TXTRecord{{Ref: "record:txt/ZG5zLmJpbmRfdHh0JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAuY3JhaWd0ZXN0MS4iY3JhaWciICJ0ZXN0IiAiMSI:craigtest1.test-ovp.bskyb.com/default", Name: "craigtest1.test-ovp.bskyb.com", Text: "craig test 1"}, {Ref: "record:txt/ZG5zLmJpbmRfdHh0JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAuY3JhaWcgdGVzdCAyLiJjcmFpZyIgInRlc3QiICIyIg:craig%20test%202.test-ovp.bskyb.com/default", Name: "craig test 2.test-ovp.bskyb.com", Text: "craig test 2"}}
	GetAllTXTRecordsAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetAllTXTRecordsAPI.GetResponse())
}
