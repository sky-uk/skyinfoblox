package skyinfoblox

import (
	"errors"
	"github.com/sky-uk/go-rest-api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"log"
	"net/http"
	"time"
)

const defaultWapiVersion = "v2.6.1"
const wapiEndpoint = "/wapi/"

//Client : the infoblox client
type Client struct {
	version    string
	restClient rest.Client
}

//Params : client connection parameters
type Params struct {
	URL         string
	User        string
	Password    string
	IgnoreSSL   bool
	Debug       bool
	Timeout     time.Duration
	WapiVersion string
}

// Connect - connects to the Infoblox server...
func Connect(params Params) *Client {

	client := new(Client)

	client.version = defaultWapiVersion
	if len(params.WapiVersion) != 0 {
		client.version = params.WapiVersion
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	client.restClient = rest.Client{
		URL:       params.URL,
		User:      params.User,
		Password:  params.Password,
		IgnoreSSL: params.IgnoreSSL,
		Debug:     params.Debug,
		Headers:   headers,
		Timeout:   params.Timeout,
	}

	return client
}

// Create - creates an object
// returns an array with these fields:
// - the created object reference ("" in case of errors)
// - the error (nil in case of success)
func (client Client) Create(objType string, profile interface{}) (string, error) {
	var objRef string
	var errStruct common.ErrorStruct

	restAPI := rest.NewBaseAPI(
		http.MethodPost,
		wapiEndpoint+client.version+"/"+objType,
		profile,
		&objRef,
		&errStruct,
	)

	err := client.restClient.Do(restAPI)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if errStruct.Error != "" {
		log.Printf("Error creating object %s, Error: %s, code: %s, text: %s",
			objType, errStruct.Error, errStruct.Code, errStruct.Text)
		return "", errors.New(errStruct.Error)
	}

	return objRef, nil
}

// Delete - deletes an object
// returns an array with these fields:
// - the deleted object reference ("" in case of errors)
// - the error (nil in case of success)
func (client Client) Delete(objRef string) (string, error) {
	var errStruct common.ErrorStruct

	restAPI := rest.NewBaseAPI(
		http.MethodDelete,
		wapiEndpoint+client.version+"/"+objRef,
		nil,
		&objRef,
		&errStruct,
	)

	err := client.restClient.Do(restAPI)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if errStruct.Error != "" {
		log.Printf("Error deleting object %s, Error: %s, code: %s, text: %s",
			objRef, errStruct.Error, errStruct.Code, errStruct.Text)
		return "", errors.New(errStruct.Error)
	}

	return objRef, nil
}

// Read - reads an object given its reference id
// The pointer to the object is passed as input param
// returns an error (nil in case of success)
func (client Client) Read(objRef string, obj interface{}) error {
	var errStruct common.ErrorStruct

	restAPI := rest.NewBaseAPI(
		http.MethodGet,
		wapiEndpoint+client.version+"/"+objRef,
		nil,
		&obj,
		&errStruct,
	)

	err := client.restClient.Do(restAPI)
	if err != nil {
		log.Println(err)
		return err
	}

	if errStruct.Error != "" {
		log.Printf("Error deleting object %s, Error: %s, code: %s, text: %s",
			objRef, errStruct.Error, errStruct.Code, errStruct.Text)
		return errors.New(errStruct.Error)
	}

	return nil
}

// ReadAll - reads all objects
func (client Client) ReadAll(objType string) ([]map[string]interface{}, error) {
	var errStruct common.ErrorStruct

	objs := make([]map[string]interface{}, 0)
	restAPI := rest.NewBaseAPI(
		http.MethodGet,
		wapiEndpoint+client.version+"/"+objType,
		nil,
		&objs,
		&errStruct,
	)

	err := client.restClient.Do(restAPI)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if errStruct.Error != "" {
		log.Printf("Error reading all objects of type %s, Error: %s, code: %s, text: %s",
			objType, errStruct.Error, errStruct.Code, errStruct.Text)
		return nil, errors.New(errStruct.Error)
	}

	return objs, nil
}

// Update - updates an object
// returns an array with these fields:
// - the updated object reference ("" in case of errors)
// - the error (nil in case of success)
func (client Client) Update(objRef string, newProfile interface{}) (string, error) {
	var errStruct common.ErrorStruct
	var updatedObjRef string

	restAPI := rest.NewBaseAPI(
		http.MethodPut,
		wapiEndpoint+client.version+"/"+objRef,
		newProfile,
		&updatedObjRef,
		&errStruct,
	)

	err := client.restClient.Do(restAPI)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if errStruct.Error != "" {
		log.Printf("Error updating object %s, Error: %s, code: %s, text: %s",
			objRef, errStruct.Error, errStruct.Code, errStruct.Text)
		return "", errors.New(errStruct.Error)
	}

	return updatedObjRef, nil
}
