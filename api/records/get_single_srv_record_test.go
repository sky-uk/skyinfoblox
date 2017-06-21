package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getSRVRecordSetup() *GetSingleSRVRecordAPI {
	recordReference := "record:srv/ZG5zLmJpbmRfc3J2JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAvY3JhaWc3dGVzdC8xMC81LzgwL2NyYWlnNHRlc3QudGVzdC1vdnAuYnNreWIuY29t:craig7test.test-ovp.bskyb.com/default"
	returnFields := []string{"name", "port", "priority", "target", "weight"}
	return NewGetSRVRecord(recordReference, returnFields)
}

func TestGetSRVRecordMethod(t *testing.T) {
	GetSingleSRVRecordAPI := getSRVRecordSetup()
	assert.Equal(t, http.MethodGet, GetSingleSRVRecordAPI.Method())
}

func TestGetSRVRecordEndpoint(t *testing.T) {
	GetSingleSRVRecordAPI := getSRVRecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:srv/ZG5zLmJpbmRfc3J2JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAvY3JhaWc3dGVzdC8xMC81LzgwL2NyYWlnNHRlc3QudGVzdC1vdnAuYnNreWIuY29t:craig7test.test-ovp.bskyb.com/default?_return_fields=name,port,priority,target,weight", GetSingleSRVRecordAPI.Endpoint())
}

func TestGetSRVRecordUnmarshalling(t *testing.T) {
	GetSingleSRVRecordAPI := getSRVRecordSetup()
	GetSingleSRVRecordAPI.SetStatusCode(http.StatusOK)
	responseObject := SRVRecord{Ref: "record:srv/ZG5zLmJpbmRfc3J2JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAvY3JhaWc3dGVzdC8xMC81LzgwL2NyYWlnNHRlc3QudGVzdC1vdnAuYnNreWIuY29t:craig7test.test-ovp.bskyb.com/default", Name: "craig7test.test-ovp.bskyb.com", Port: 80, Priority: 10, Target: "craig4test.test-ovp.bskyb.com", Weight: 5}
	GetSingleSRVRecordAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetSingleSRVRecordAPI.GetResponse())
}
