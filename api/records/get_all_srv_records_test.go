package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getAllSRVRecordsSetup() *GetAllSRVRecordsAPI {
	returnFields := []string{"name", "port", "priority", "target", "weight"}
	return NewGetAllSRVRecords(returnFields)
}

func TestGetAllSRVRecordsMethod(t *testing.T) {
	GetAllSRVRecordsAPI := getAllSRVRecordsSetup()
	assert.Equal(t, http.MethodGet, GetAllSRVRecordsAPI.Method())
}

func TestGetAllSRVRecordsEndpoint(t *testing.T) {
	GetAllSRVRecordsAPI := getAllSRVRecordsSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:srv?_return_fields=name,port,priority,target,weight", GetAllSRVRecordsAPI.Endpoint())
}

func TestGetAllSRVRecordsUnmarshalling(t *testing.T) {
	GetAllSRVRecordsAPI := getAllSRVRecordsSetup()
	responseObject := []SRVRecord{{Ref: "record:srv/ZG5zLmJpbmRfc3J2JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAvY3JhaWc3dGVzdC8xMC81LzgwL2NyYWlnNHRlc3QudGVzdC1vdnAuYnNreWIuY29t:craig7test.test-ovp.bskyb.com/default", Name: "craig7test.test-ovp.bskyb.com", Port: 80, Priority: 10, Target: "craig4test.test-ovp.bskyb.com", Weight: 5}, {Ref: "record:srv/ZG5zLmJpbmRfc3J2JC5fZGVmYXVsdC5jb20uYnNreWIudGVzdC1vdnAvY3JhaWc4dGVzdC8xMC81LzgwL2NyYWlnNHRlc3QudGVzdC1vdnAuYnNreWIuY29t:craig8test.test-ovp.bskyb.com/default", Name: "craig8test.test-ovp.bskyb.com", Port: 80, Priority: 10, Target: "craig4test.test-ovp.bskyb.com", Weight: 5}}
	GetAllSRVRecordsAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetAllSRVRecordsAPI.GetResponse())
}
