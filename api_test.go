package skyinfoblox

import (
	"errors"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func getClient() (*Client, error) {
	server, ok := os.LookupEnv("INFOBLOX_SERVER")
	if ok == false || server == "" {
		return nil, errors.New("INFOBLOX_SERVER env var not set")
	}

	username, ok := os.LookupEnv("INFOBLOX_USERNAME")
	if ok == false {
		return nil, errors.New("INFOBLOX_USERNAME env var not set")
	}

	password, ok := os.LookupEnv("INFOBLOX_PASSWORD")
	if ok == false {
		return nil, errors.New("INFOBLOX_PASSWORD env var not set")
	}

	params := Params{
		WapiVersion: "v2.6.1", // this is anyhow the default...
		URL:         server,
		User:        username,
		Password:    password,
		IgnoreSSL:   true,
		Debug:       true,
	}

	client := Connect(params)

	return client, nil
}

func TestFilterProfileKeys(t *testing.T) {
	adminuser := map[string]interface{}{
		"name":         "user1",
		"comment":      "this is a comment",
		"email":        "exampleuser@domain.internal.com",
		"admin_groups": []string{"APP-OVP-INFOBLOX-READONLY"},
		"password":     "c0a6264f0f128d94cd8ef26652e7d9fd",
	}
	validProfile := FilterProfileKeys(
		adminuser,
		[]string{"name", "comment"},
	)

	keys := make([]string, 0)
	for k := range validProfile {
		keys = append(keys, k)
	}
	assert.Equal(t, 2, len(keys))
	assert.Equal(t, "user1", validProfile["name"])
	assert.Equal(t, "this is a comment", validProfile["comment"])
}

func TestGetValidKeys(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal(err)
	}

	validKeysWhileReading := client.GetValidKeys("adminuser", []string{"r"})
	assert.Equal(t, 9, len(validKeysWhileReading))
	validKeysWhileWriting := client.GetValidKeys("adminuser", []string{"w"})
	assert.Equal(t, 10, len(validKeysWhileWriting))
	validKeysWhileUpdating := client.GetValidKeys("adminuser", []string{"u"})
	assert.Equal(t, 10, len(validKeysWhileUpdating))
}

func TestGetObjectSchema(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatal(err)
	}

	schema, err := client.GetObjectSchema("adminuser")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetObjectTypeFromRef(t *testing.T) {
	ref := "adminrole/b25lLnJvbGUkdGVzdDQ2Mw:test463"
	objType := GetObjectTypeFromRef(ref)
	assert.Equal(t, "adminrole", objType)
}

func TestFilterReturnFields(t *testing.T) {
	required := []string{"one", "two", "three"}
	allowed := []string{"one", "three"}
	filtered := FilterReturnFields(required, allowed)
	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, "one", filtered[0])
	assert.Equal(t, "three", filtered[1])

}

func TestAllAPI(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	client, err := getClient()
	if err != nil {
		t.Fatal(err)
	}

	// this API works with a defined struct...
	/*
		disable := true
		superUser := false


			adminGroup := model.IBXAdminGroup{
				AccessMethod:   []string{"API"},
				Comment:        "API Access only",
				Disable:        &disable,
				EmailAddresses: []string{"test@example-test.com"},
				Name:           "test",
				Roles:          []string{"test-role"},
				SuperUser:      &superUser,
			}
	*/

	// or with a generic map (that matches a given schema...)
	//adminGroup := make(map[string]interface{})
	//adminGroup["name"] = "test"
	adminRole := make(map[string]interface{})
	adminRole["name"] = "test" + strconv.Itoa(rand.Intn(1000))
	adminRole["comment"] = "An initial comment"

	// creating an object...
	refObj, err := client.Create("adminrole", adminRole)
	if err != nil {
		t.Fatal("Error creating an adminrole object")
	}
	assert.NotEmpty(t, refObj)
	t.Log("Object created, REFOBJ: ", refObj)

	//reading the object...
	role := make(map[string]interface{})
	err = client.Read(refObj, []string{"comment"}, &role)
	if err != nil {
		t.Fatal("Error reading object with ref: ", refObj)
	}
	t.Logf("Object (as map):\n%+v\n", role)

	//reading the object as struct...
	var roleObj model.AdminRole
	err = client.Read(refObj, []string{"comment"}, &roleObj)
	if err != nil {
		t.Fatal("Error reading object with ref: ", refObj)
	}
	t.Logf("Object (as struct):\n%+v\n", roleObj)

	//getting all roles...
	roles, err := client.ReadAll("adminrole")
	if err != nil {
		t.Fatal("Error reading all roles")
	}
	t.Logf("Objects:\n%+v\n", roles)

	//updating the object...
	adminRole["comment"] = "Object updated"
	updatedRefObj, err := client.Update(refObj, adminRole)
	if err != nil {
		t.Fatal("Error updating the object")
	}
	t.Logf("Object %s updated\n", updatedRefObj)

	// getting the updated object and chedking for the comment...
	err = client.Read(updatedRefObj, []string{"comment"}, &role)
	if err != nil {
		t.Fatal("Error reading object with ref: ", updatedRefObj)
	}
	t.Log("Updated object comment: ", role["comment"])
	assert.Equal(t, "Object updated", role["comment"])

	//deleting the object
	refObj, err = client.Delete(refObj)
	if err != nil {
		t.Fatal("Error creating an adminrole object")
	}
	assert.NotEmpty(t, refObj)
	t.Log("Object deleted, REFOBJ: ", refObj)

	// now creating and reading a user to check
	// control on attributes
	newUser := make(map[string]interface{})
	newUser["name"] = "user_" + strconv.Itoa(rand.Intn(1000))
	newUser["password"] = "foooooo" // at least 4 chars...
	newUser["comment"] = "test user for attributes check"
	newUser["admin_groups"] = []string{"APP-OVP-INFOBLOX-READONLY"}
	newUserRef, err := client.Create("adminuser", newUser)
	if err != nil {
		t.Fatal("Error creating an adminuser object")
	}
	// now we try to read the password (which is forbidden)
	user := make(map[string]interface{})
	err = client.Read(newUserRef, []string{"name", "password"}, &user)
	if err != nil {
		t.Fatal("Error reading an adminuser object")
	}
	assert.Equal(t, 2, len(user))
	assert.Equal(t, newUserRef, user["_ref"])
	assert.Equal(t, newUser["name"], user["name"])
}
