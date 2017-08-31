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
		Ref:           "7aa35cba48a0ebf27c549465a60f3b24",
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
		mxRecordAPI := NewCreate(newMXRecord)
		response := fmt.Sprintf("%s/%s/randonstring:7aa35cba48a0ebf27c549465a60f3b24", wapiVersion, mxRecordEndpoint)
		fmt.Println(response)
		mxRecordAPI.SetResponseObject(response)
		return mxRecordAPI
	case "read":
		reference = "record:mx/7aa35cba48a0ebf27c549465a60f3b24"
		returnFields = []string{"name", "comment", "mail_exchanger"}
		mxRecordAPI := NewGet(reference, returnFields)
		mxRecordAPI.SetResponseObject(newMXRecord)
		return mxRecordAPI
	case "readall":
		mxRecordAPI := NewGetAll()
		mxRecordAPI.SetResponseObject(newMXRecord)
		return mxRecordAPI
	case "update":
		reference = "7aa35cba48a0ebf27c549465a60f3b24"
		returnFields = []string{"name", "comment", "mail_exchanger"}
		mxRecordAPI := NewUpdate(reference, newMXRecord)
		return mxRecordAPI
	case "delete":
		reference = "record:mx/7aa35cba48a0ebf27c549465a60f3b24"
		mxRecordAPI := NewDelete(reference)
		return mxRecordAPI

	default:
		return nil
	}
}

func TestCreateMXRecordEndpoint(t *testing.T) {
	createMxRecordAPI := setupMXRecord("create")
	assert.Equal(t, "/wapi/v2.6.1/record:mx", createMxRecordAPI.Endpoint())
}

func TestCreateMXRecordMethod(t *testing.T) {
	createMxRecordAPI := setupMXRecord("create")
	assert.Equal(t, http.MethodPost, createMxRecordAPI.Method())
}

func TestCreateMXRecordResponse(t *testing.T) {
	createMxRecordAPI := setupMXRecord("create")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/randonstring:7aa35cba48a0ebf27c549465a60f3b24", createMxRecordAPI.ResponseObject())
}

func TestReadMXRecordEndpoint(t *testing.T) {
	getMxRecordAPI := setupMXRecord("read")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/7aa35cba48a0ebf27c549465a60f3b24?_return_fields=name,comment,mail_exchanger", getMxRecordAPI.Endpoint())
}

func TestReadMXRecordMethod(t *testing.T) {
	getMxRecordAPI := setupMXRecord("read")
	assert.Equal(t, http.MethodGet, getMxRecordAPI.Method())
}

func TestReadMXRecordResponse(t *testing.T) {
	getMxRecordAPI := setupMXRecord("read")
	expectedResponse := MxRecord{
		Ref:               "7aa35cba48a0ebf27c549465a60f3b24",
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
	assert.Equal(t, expectedResponse, getMxRecordAPI.ResponseObject())
}

func TestUpdateMXRecordEndpoint(t *testing.T) {
	getMxRecordAPI := setupMXRecord("update")
	assert.Equal(t, http.MethodPut, getMxRecordAPI.Method())
}

func TestDeleteMXRecordEndpoint(t *testing.T) {
	getMxRecordAPI := setupMXRecord("delete")
	assert.Equal(t, "/wapi/v2.6.1/record:mx/7aa35cba48a0ebf27c549465a60f3b24", getMxRecordAPI.Endpoint())
}

func TestDeleteMXRecordMethod(t *testing.T) {
	getMxRecordAPI := setupMXRecord("delete")
	assert.Equal(t, http.MethodDelete, getMxRecordAPI.Method())
}

func TestGetAllMXRecordEndpoint(t *testing.T) {
	getMxRecordAPI := setupMXRecord("readall")
	assert.Equal(t, "/wapi/v2.6.1/record:mx", getMxRecordAPI.Endpoint())
}

func TestGetAllMXRecordMethod(t *testing.T) {
	getMxRecordAPI := setupMXRecord("readall")
	assert.Equal(t, http.MethodGet, getMxRecordAPI.Method())
}
