package skyinfoblox

import (
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestAllAPI(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	server, ok := os.LookupEnv("INFOBLOX_SERVER")
	if ok == false || server == "" {
		t.Skipf("INFOBLOX_SERVER env var not set")
	}

	username, ok := os.LookupEnv("INFOBLOX_USERNAME")
	if ok == false {
		t.Skip("INFOBLOX_USERNAME env var not set")
	}

	password, ok := os.LookupEnv("INFOBLOX_PASSWORD")
	if ok == false {
		t.Skip("INFOBLOX_PASSWORD env var not set")
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

	// creating an object...
	refObj, err := client.Create("adminrole", adminRole)
	if err != nil {
		t.Fatal("Error creating an adminrole object")
	}
	assert.NotEmpty(t, refObj)
	t.Log("Object created, REFOBJ: ", refObj)

	//reading the object...
	role := make(map[string]interface{})
	err = client.Read(refObj, &role)
	if err != nil {
		t.Fatal("Error reading object with ref: ", refObj)
	}
	t.Logf("Object (as map):\n%+v\n", role)

	//reading the object as struct...
	var roleObj model.AdminRole
	err = client.Read(refObj, &roleObj)
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
	err = client.Read(updatedRefObj, &role)
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
}
