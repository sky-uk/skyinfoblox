package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"net/http"
	"os"
	"strings"
)

var createAdminGroupObject admingroup.IBXAdminGroup
var createAdminGroupAccessMethods, createAdminGroupEmailAddresses, createAdminGroupRoles string
var createAdminGroupSuperUser, createAdminGroupDisable bool

func createAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if createAdminGroupObject.Name == "" {
		fmt.Printf("\nError name argument is required\n")
		os.Exit(1)
	}

	if createAdminGroupAccessMethods != "" {
		createAdminGroupObject.AccessMethod = strings.Split(createAdminGroupAccessMethods, ",")
	}
	if createAdminGroupEmailAddresses != "" {
		createAdminGroupObject.EmailAddresses = strings.Split(createAdminGroupEmailAddresses, ",")
	}
	if createAdminGroupRoles != "" {
		createAdminGroupObject.Roles = strings.Split(createAdminGroupRoles, ",")
	}
	createAdminGroupObject.Disable = &createAdminGroupDisable
	createAdminGroupObject.SuperUser = &createAdminGroupSuperUser

	createAdminGroupAPI := admingroup.NewCreate(createAdminGroupObject)
	err := client.Do(createAdminGroupAPI)
	httpStatus := createAdminGroupAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating admin group %s. HTTP status: %d. Error: %+v\n", createAdminGroupObject.Name, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created admin group %s\n", createAdminGroupObject.Name)

}

func init() {
	createAdminGroupFlags := flag.NewFlagSet("admin-group-create", flag.ExitOnError)
	createAdminGroupFlags.StringVar(&createAdminGroupObject.Name, "name", "", "usage: -name admin-group-name")
	createAdminGroupFlags.StringVar(&createAdminGroupObject.Comment, "comment", "", "usage: -comment 'A comment'")
	createAdminGroupFlags.BoolVar(&createAdminGroupSuperUser, "super-user", false, "usage: -super-user")
	createAdminGroupFlags.BoolVar(&createAdminGroupDisable, "disable", false, "usage: -disable")
	createAdminGroupFlags.StringVar(&createAdminGroupAccessMethods, "access-method", "GUI,API,TAXII", "usage: -access-method method (One or more of API, CLOUD_API, GUI, TAXII")
	createAdminGroupFlags.StringVar(&createAdminGroupEmailAddresses, "email-addresses", "", "usage: -email-addresses emailaddress@domain,emailaddress2@domain....")
	createAdminGroupFlags.StringVar(&createAdminGroupRoles, "roles", "", "usage: -roles role1,role2...")
	RegisterCliCommand("admin-group-create", createAdminGroupFlags, createAdminGroup)
}
