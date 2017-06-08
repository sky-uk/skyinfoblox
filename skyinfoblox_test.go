package skyinfoblox

import (
	"encoding/base64"
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var user = "infobloxUser"
var password = "infobloxPassword"
var ignoreSSL = true
var debug = false
var infobloxClient *InfobloxClient
var server *httptest.Server

const (
	unauthorizedStatusCode = http.StatusUnauthorized
	unauthorizedResponse   = `<!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN"><html><head><title>401 Authorization Required</title></head><body><h1>Authorization Required</h1><p>This server could not verify that youare authorized to access the documentrequested.  Either you supplied the wrongcredentials (e.g., bad password), or yourbrowser doesn't understand how to supplythe credentials required.</p></body></html>`
)

func hasHeader(req *http.Request, name string, value string) bool {
	return req.Header.Get(name) == value
}

func setup(statusCode int, responseBody string) {
	basicAuthHeaderValue := "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password))
	server = httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !hasHeader(r, "Authorization", basicAuthHeaderValue) {
				w.WriteHeader(unauthorizedStatusCode)
				fmt.Fprint(w, unauthorizedResponse)
				return
			}
			w.WriteHeader(statusCode)
			fmt.Fprintln(w, responseBody)
		}))
	infobloxClient = NewInfobloxClient(server.URL, user, password, ignoreSSL, debug)
}

func TestHappyCase(t *testing.T) {
	setup(200, "pong")
	infobloxClient = NewInfobloxClient(server.URL, user, password, ignoreSSL, debug)
	apiRequest := api.NewBaseAPI(http.MethodGet, "/", nil, nil)

	err := infobloxClient.Do(apiRequest)

	assert.Nil(t, err)
}

func TestBasicAuthFailure(t *testing.T) {
	setup(0, "")
	infobloxClient = NewInfobloxClient(server.URL, "invalidUser", "invalidPass", ignoreSSL, debug)

	apiRequest := api.NewBaseAPI(http.MethodGet, "/", nil, nil)
	infobloxClient.Do(apiRequest)

	assert.Equal(t, 401, apiRequest.StatusCode())
	assert.Equal(t, unauthorizedResponse, string(apiRequest.RawResponse()))

}

func TestIsJSON(t *testing.T) {
	assert.True(t, isJSON("application/json"))
	assert.True(t, isJSON("text/json"))
	assert.True(t, isJSON("text/json; charset=utf-8"))

	assert.False(t, isJSON("application/html"))
	assert.False(t, isJSON("text/html"))
	assert.False(t, isJSON("text/html; charset=utf-8"))
}
