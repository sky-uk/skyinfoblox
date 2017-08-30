package mxrecord

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var returnFields []string
var reference string

func setupMXRecord(action string) *api.BaseAPI {
	newMXRecord := MxRecord{
		Ref:           "SVC-APP-UNIT-TEST",
		Comment:       "This is a comment on an MXRecord",
		MailExchanger: "mail.example.com",
		Name:          "mydomain.com",
		Preference:    1,
		TTL:           1600,
		UseTTL:        true,
		View:          "default",
	}

	switch action {
	case "create":
		MXRecordAPI := NewCreate(newMXRecord)
		response := fmt.Sprintf("%s/%s/randonstring:SVC-APP-UNIT-TEST", WapiVersion, MXRecordEndpoint)
		fmt.Println(response)
		MXRecordAPI.SetResponseObject(response)
		return MXRecordAPI
	case "read":
		reference = "SVC-APP-UNIT-TEST"
		returnFields = []string{"name", "comment", "mail_exchanger"}
		MXRecordAPI := NewGet(newMXRecord.Ref, returnFields)
		MXRecordAPI.SetResponseObject(newMXRecord)
		return MXRecordAPI
	case "readall":
		MXRecordAPI := NewGetAll()
		MXRecordAPI.SetResponseObject(newMXRecord)
		return MXRecordAPI
	case "update":
		reference = "SVC-APP-UNIT-TEST"
		returnFields = []string{"name", "comment", "mail_exchanger"}
		MXRecordAPI := NewUpdate(newMXRecord, returnFields)
		return MXRecordAPI
	case "delete":
		reference = "SVC-APP-UNIT-TEST"
		MXRecordAPI := NewDelete(reference)
		return MXRecordAPI

	default:
		return nil
	}
}

func TestCreateMXRecordEndpoint(t *testing.T) {
	CreateMXRecordAPI := setupMXRecord("create")
	assert.Equal(t, "/wapi/v2.6.1/record:mx", CreateMXRecordAPI.Endpoint())
}

func TestCreateMXRecordMethod(t *testing.T) {
	CreateMXRecordAPI := setupMXRecord("create")
	assert.Equal(t, http.MethodPost, CreateMXRecordAPI.Method())
}

func TestCreateMXRecordResponse(t *testing.T) {
	CreateMXRecordAPI := setupMXRecord("create")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/randonstring:SVC-APP-UNIT-TEST", CreateMXRecordAPI.ResponseObject())
}

func TestReadMXRecordEndpoint(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("read")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/SVC-APP-UNIT-TEST", GetMXRecordAPI.Endpoint())
}

func TestReadMXRecordMethod(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("read")
	assert.Equal(t, http.MethodGet, GetMXRecordAPI.Method())
}

func TestReadMXRecordResponse(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("read")
	expectedResponse := MxRecord{
		Ref:               "SVC-APP-UNIT-TEST",
		Comment:           "This is a comment on an MXRecord",
		DDNSPrincipal:     "",
		DDNSProtected:     false,
		Disable:           false,
		ForbidReclamation: false,
		MailExchanger:     "mail.example.com",
		Name:              "mydomain.com",
		Preference:        1,
		TTL:               1600,
		UseTTL:            true,
		View:              "default"}
	assert.Equal(t, expectedResponse, GetMXRecordAPI.ResponseObject())
}

func TestUpdateMXRecordEndpoint(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("update")
	assert.Equal(t, http.MethodPut, GetMXRecordAPI.Method())
}

func TestDeleteMXRecordEndpoint(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("delete")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/SVC-APP-UNIT-TEST", GetMXRecordAPI.Endpoint())
}

func TestDeleteMXRecordMethod(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("delete")
	assert.Equal(t, http.MethodDelete, GetMXRecordAPI.Method())
}

func TestGetAllMXRecordEndpoint(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("readall")
	assert.Equal(t, "/wapi/v2.6.1/record:mx", GetMXRecordAPI.Endpoint())
}

func TestGetAllMXRecordMethod(t *testing.T) {
	GetMXRecordAPI := setupMXRecord("readall")
	assert.Equal(t, http.MethodGet, GetMXRecordAPI.Method())
}
