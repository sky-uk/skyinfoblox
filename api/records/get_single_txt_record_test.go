package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getTXTRecordSetup() *GetSingleTXTRecordAPI {
	recordReference := "record:txt/ZG5zLmJpbmRfdHh0JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAuY3JhaWd0ZXN0MS4iY3JhaWciICJ0ZXN0IiAiMSI:craigtest1.test-ovp.bskyb.com/default"
	returnFields := []string{"name", "text"}
	return NewGetTXTRecord(recordReference, returnFields)
}

func TestGetTXTRecordMethod(t *testing.T) {
	GetSingleTXTRecordAPI := getTXTRecordSetup()
	assert.Equal(t, http.MethodGet, GetSingleTXTRecordAPI.Method())
}

func TestGetTXTRecordEndpoint(t *testing.T) {
	GetSingleTXTRecordAPI := getTXTRecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:txt/ZG5zLmJpbmRfdHh0JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAuY3JhaWd0ZXN0MS4iY3JhaWciICJ0ZXN0IiAiMSI:craigtest1.test-ovp.bskyb.com/default?_return_fields=name,text", GetSingleTXTRecordAPI.Endpoint())
}

func TestGetTXTRecordUnmarshalling(t *testing.T) {
	GetSingleTXTRecordAPI := getTXTRecordSetup()
	responseObject := TXTRecord{Ref: "record:txt/ZG5zLmJpbmRfdHh0JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAuY3JhaWd0ZXN0MS4iY3JhaWciICJ0ZXN0IiAiMSI:craigtest1.test-ovp.bskyb.com/default", Name: "craigtest1.test-ovp.bskyb.com", Text: "craig test 1"}
	GetSingleTXTRecordAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetSingleTXTRecordAPI.GetResponse())
}
