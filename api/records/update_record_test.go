package records

import (
//"testing"
//"github.com/stretchr/testify/assert"
//"net/http"
)
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func updateRecordSetup() *UpdateRecordAPI {
	recordReference := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default"
	updatePayload := GenericRecord{IPv4: "10.10.10.90"}
	return NewUpdateRecord(recordReference, updatePayload)
}

func TestUpdateRecordMethod(t *testing.T) {
	updateRecordAPI := updateRecordSetup()
	assert.Equal(t, http.MethodPut, updateRecordAPI.Method())
}

func TestUpdateRecordEndpoint(t *testing.T) {
	updateRecordAPI := updateRecordSetup()
	assert.Equal(t, fmt.Sprintf("%s/record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default", wapiVersion), updateRecordAPI.Endpoint())
}

func TestUpdateRecordUnmarshalling(t *testing.T) {
	updateRecordAPI := updateRecordSetup()
	responseString := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLmJza3liLnRlc3Qtb3ZwLGNyYWlnNHRlc3QsMTAuMTAuMTAuNA:craig4test.test-ovp.bskyb.com/default"
	updateRecordAPI.SetResponseObject(&responseString)
	assert.Equal(t, responseString, updateRecordAPI.GetResponse())
}
