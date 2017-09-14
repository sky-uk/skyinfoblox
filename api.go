package skyinfoblox

import (
	"errors"
	"github.com/sky-uk/go-rest-api"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"log"
	"net/http"
	"strings"
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

func getProfileKeys(profile map[string]interface{}) []string {
	keys := []string{}

	for k := range profile {
		keys = append(keys, k)
	}
	return keys
}

// Create - creates an object
// returns an array with these fields:
// - the created object reference ("" in case of errors)
// - the error (nil in case of success)
func (client Client) Create(objType string, profile interface{}) (string, error) {
	var objRef string
	var errStruct common.ErrorStruct

	if profile, ok := profile.(map[string]interface{}); ok {
		client.FilterProfileAttrs(objType, profile, []string{"w"})
	}

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
		log.Printf("Error creating object %s, Error: %s, code: %s, text: %s\n",
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
func (client Client) Read(objRef string, returnFields []string, obj interface{}) error {
	var errStruct common.ErrorStruct

	queryStr := wapiEndpoint + client.version + "/" + objRef

	objType := GetObjectTypeFromRef(objRef)
	validKeys := client.GetValidKeys(objType, []string{"r"})
	fields := FilterReturnFields(returnFields, validKeys)

	if len(returnFields) > 0 {
		queryStr += "?_return_fields=" + strings.Join(fields, ",")
	}

	restAPI := rest.NewBaseAPI(
		http.MethodGet,
		queryStr,
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

	if newProfile, ok := newProfile.(map[string]interface{}); ok {
		objType := GetObjectTypeFromRef(objRef)
		client.FilterProfileAttrs(objType, newProfile, []string{"u"})
	}

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

/*
FilterProfileAttrs - filters out profile attributes
Workflow:
  - as the object can have nested structs, we need to proceed in this way:
  - we have the object type
  - we get the object schema
  - for each schema attr we get:
      - the type(s)
      - the authentication rules
      - the is_array boolean flag
  - we also have a global map of schemas for structs in the model package
  SO:
  - for each object attr:
    - for each attribute type
      - if type is not in global struct map
        - delete attribute from profile if not valid
      - else (it's a struct, we have hence metadata for it from model...):
        - if is_array:
          - for each array item (needs to have a pointer to the struct):
            - for each attr in item:
              - delete from item if not valid
------------------------------------------------------------------------------*/
func (client Client) FilterProfileAttrs(objType string, profile map[string]interface{}, filter []string) {

	log.Println("Filtering profile:\n", profile)

	structsAttrData := model.StructAttrs()

	schema, err := client.GetObjectSchema(objType)
	if err != nil {
		log.Printf("Error getting schema for object %s, error: %+v\n", objType, err)
	}
	fields := schema["fields"].([]interface{})
	for _, field := range fields {
		log.Println("Analizing field: ", field)
		fieldAsMap := field.(map[string]interface{})
		if profileItem, found := profile[fieldAsMap["name"].(string)]; found {
			for _, attrType := range fieldAsMap["type"].([]interface{}) {
				if structData, found := structsAttrData[attrType.(string)]; found {
					log.Println("Type is a struct", attrType)
					log.Println("Struct metadata:\n", structData)
					// attribute is a struct...
					if fieldAsMap["is_array"].(bool) == true {
						for _, item := range profileItem.([]map[string]interface{}) {
							filterStruct(item, structData.(map[string]model.SchemaAttr), filter)
						}
					}
				} else {
					// attribute is a scalar...
					filterAttr(profile, fieldAsMap, filter)
				}
			}
		} else {
			log.Printf("Attribute %s not defined in profile\n", fieldAsMap["name"].(string))
		}
	}
}

func filterStruct(item map[string]interface{}, attrData map[string]model.SchemaAttr, filter []string) {
	for attr := range item {
		valid := false
		for _, operation := range filter {
			if strings.Contains(attrData[attr].Supports, operation) {
				valid = true
				break
			}
		}
		if valid == false {
			delete(item, attr)
		}
	}
}

func filterAttr(item map[string]interface{}, attrData map[string]interface{}, filter []string) {
	valid := false
	for _, operation := range filter {
		if strings.Contains(attrData["supports"].(string), operation) {
			valid = true
			break
		}
	}
	if valid == false {
		delete(item, attrData["name"].(string))
	}
}

// GetValidKeys - retrieves the list of valid keys for the performed operation
// from the object schema
func (client Client) GetValidKeys(objType string, filter []string) []string {

	validKeys := []string{}
	schema, err := client.GetObjectSchema(objType)
	if err != nil {
		log.Printf("Error getting schema for object %s, error: %+v\n", objType, err)
	}
	fields := schema["fields"].([]interface{})
	for _, field := range fields {
		fieldAsMap := field.(map[string]interface{})
		for _, operation := range filter {
			if strings.Contains(fieldAsMap["supports"].(string), operation) {
				validKeys = append(validKeys, fieldAsMap["name"].(string))
				break
			}
		}
	}
	return validKeys
}

// FilterProfileKeys - filters the keys of the provided map, deleting the ones
// not contained in the valid keys list
func FilterProfileKeys(profile map[string]interface{}, validKeys []string) map[string]interface{} {

	outMap := make(map[string]interface{})
	for key, value := range profile {
		for _, validKey := range validKeys {
			if key == validKey {
				outMap[validKey] = value
				break
			}
		}

	}
	return outMap
}

// GetObjectSchema - retrieves the object schmea
func (client Client) GetObjectSchema(objType string) (map[string]interface{}, error) {

	var errStruct common.ErrorStruct
	schema := make(map[string]interface{})

	api := rest.NewBaseAPI(
		http.MethodGet,
		wapiEndpoint+client.version+"/"+objType+"?_schema",
		nil,
		&schema,
		&errStruct,
	)

	err := client.restClient.Do(api)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if errStruct.Error != "" {
		log.Printf("Error getting schema for object type %s, Error: %s, code: %s, text: %s",
			objType, errStruct.Error, errStruct.Code, errStruct.Text)
		return nil, errors.New(errStruct.Error)
	}

	return schema, nil
}

// GetObjectTypeFromRef - returns the object type given an object reference
// Object reference format:
// wapitype / refdata [ : name1 [ { / nameN }... ] ]
func GetObjectTypeFromRef(ref string) string {
	tokens := strings.Split(ref, "/")
	return tokens[0]
}

// FilterReturnFields - filters the list of required return fields based on
// the list of readable ones
func FilterReturnFields(required, allowed []string) []string {
	outList := []string{}
	for _, reqItem := range required {
		for _, allowedItem := range allowed {
			if reqItem == allowedItem {
				outList = append(outList, allowedItem)
				break
			}
		}
	}
	return outList
}
