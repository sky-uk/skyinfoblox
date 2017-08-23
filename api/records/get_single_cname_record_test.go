package records

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getCNAMERecordSetup() *GetSingleCNAMERecordAPI {
	recordReference := "record:cname/ZG5zLmJpbmRfY25hbWUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cC5jcmFpZzV0ZXN0:craig5test.test-ovp.bskyb.com/default"
	returnFields := []string{"name", "canonical"}
	return NewGetCNAMERecord(recordReference, returnFields)
}

func TestGetCNAMERecordMethod(t *testing.T) {
	GetSingleCNAMERecordAPI := getCNAMERecordSetup()
	assert.Equal(t, http.MethodGet, GetSingleCNAMERecordAPI.Method())
}

func TestGetCNAMERecordEndpoint(t *testing.T) {
	GetSingleCNAMERecordAPI := getCNAMERecordSetup()
	assert.Equal(t, fmt.Sprintf("%s/record:cname/ZG5zLmJpbmRfY25hbWUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cC5jcmFpZzV0ZXN0:craig5test.test-ovp.bskyb.com/default?_return_fields=name,canonical", wapiVersion), GetSingleCNAMERecordAPI.Endpoint())
}

func TestGetCNAMERecordUnmarshalling(t *testing.T) {
	GetSingleCNAMERecordAPI := getCNAMERecordSetup()
	responseObject := CNAMERecord{Ref: "record:cname/ZG5zLmJpbmRfY25hbWUkLl9kZWZhdWx0LmNvbS5ic2t5Yi50ZXN0LW92cC5jcmFpZzV0ZXN0:craig5test.test-ovp.bskyb.com/default", Name: "craig5test.test-ovp.bskyb.com", Canonical: "test-ovp.bskyb.com"}
	GetSingleCNAMERecordAPI.SetResponseObject(&responseObject)
	assert.Equal(t, responseObject, GetSingleCNAMERecordAPI.GetResponse())
}
