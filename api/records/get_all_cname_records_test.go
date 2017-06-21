package records

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getAllCNAMERecordsSetup() *GetAllCNAMERecordsAPI {
	returnFields := []string{"canonical", "name"}
	return NewGetAllCNAMERecords(returnFields)
}

func TestGetAllCNAMERecordsMethod(t *testing.T) {
	GetAllCNAMERecordsAPI := getAllCNAMERecordsSetup()
	assert.Equal(t, http.MethodGet, GetAllCNAMERecordsAPI.Method())
}

func TestGetAllCNAMERecordsEndpoint(t *testing.T) {
	GetAllCNAMERecordsAPI := getAllCNAMERecordsSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:cname?_return_fields=canonical,name", GetAllCNAMERecordsAPI.Endpoint())
}

func TestGetAllCNAMERecordsUnmarshalling(t *testing.T) {
	GetAllCNAMERecordsAPI := getAllCNAMERecordsSetup()
	GetAllCNAMERecordsAPI.SetStatusCode(http.StatusOK)
	responseObject := []CNAMERecord{{Ref: "record:cname/ZG5zLmJpbmRfY25hbWUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cC5jZHRlc3Q:cdtest.test-ovp.bskyb.com/default", Canonical: "test-ovp.bskyb.com", Name: "cdtest.test-ovp.bskyb.com"}, {Ref: "record:cname/ZG5zLmJpbmRfY25hbWUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cC5jcmFpZzV0ZXN0:craig5test.test-ovp.bskyb.com/default", Canonical: "test-ovp.bskyb.com", Name: "craig5test.test-ovp.bskyb.com"}}
	GetAllCNAMERecordsAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetAllCNAMERecordsAPI.GetResponse())
}
