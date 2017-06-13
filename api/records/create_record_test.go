package records

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func createGenericRecordSetup() *CreateRecordAPI {
	newARecord := GenericRecord{
		IPv4:    "1.1.1.1",
		Name:    "test.example.com",
		Comment: "test comment",
	}
	return NewCreateRecord("a", newARecord)
}

func createARecordSetup() *CreateRecordAPI {
	newARecord := ARecord{
		IPv4:    "1.1.1.1",
		Name:    "test",
		Comment: "test comment",
		Zone:    "example.com",
	}
	return NewCreateARecord(newARecord)
}

func TestCreateMethod(t *testing.T) {
	createRecordAPI := createGenericRecordSetup()
	assert.Equal(t, http.MethodPost, createRecordAPI.Method())
}

func TestNewCreateARecord(t *testing.T) {
	createARecordAPI := createARecordSetup()
	assert.Equal(t, http.MethodPost, createARecordAPI.Method())
}

func TestCreateARecordEndpoint(t *testing.T) {
	createARecordAPI := createARecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:a", createARecordAPI.Endpoint())

}
func TestCreateEndpoint(t *testing.T) {
	createRecordAPI := createGenericRecordSetup()
	assert.Equal(t, "/wapi/v2.3.1/record:a", createRecordAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createRecordAPI := createGenericRecordSetup()
	expectedJSON := "{\"name\":\"test.example.com\",\"comment\":\"test comment\",\"ipv4addr\":\"1.1.1.1\"}"
	jsonBytes, err := json.Marshal(createRecordAPI.RequestObject())
	assert.Nil(t, err)
	assert.Equal(t, expectedJSON, string(jsonBytes))
}

func TestCreateRecordUnmarshalling(t *testing.T) {
	createRecordAPI := createGenericRecordSetup()
	responseString := "record:a/ZG5zLmJpbmRfYSQuX2RlZmF1bHQuY29tLnNreS5vdnAubnAsdGVzdC55b3JnLDEwLjEwLjEwLjEw:yorg.test.np.ovp.sky.com/default"
	createRecordAPI.SetResponseObject(&responseString)
	assert.Equal(t, responseString, createRecordAPI.GetResponse())
}
