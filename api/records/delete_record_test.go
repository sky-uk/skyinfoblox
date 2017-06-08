package records

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func deleteRecordSetup() *DeleteRecordAPI {
	recordReference := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default"
	return NewDelete(recordReference)
}

func TestDeleteRecordMethod(t *testing.T){
	deleteRecordAPI := deleteRecordSetup()
	assert.Equal(t, http.MethodDelete,deleteRecordAPI.Method())
}

func TestDeleteRecordEndpoint(t *testing.T){
	deleteRecordAPI := deleteRecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default", deleteRecordAPI.Endpoint())
}

func TestDeleteRecordUnmarshalling(t *testing.T) {
	deleteRecordAPI := deleteRecordSetup()
	responseString := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default"
	deleteRecordAPI.SetResponseObject(&responseString)
	assert.Equal(t, responseString, deleteRecordAPI.GetResponse())
}