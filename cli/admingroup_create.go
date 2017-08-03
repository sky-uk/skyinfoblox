package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func createAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var createAdminGroupObject admingroup.IBXAdminGroup

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	superUser, superUserErr := strconv.ParseBool(flagSet.Lookup("super-user").Value.String())
	disable, disableErr := strconv.ParseBool(flagSet.Lookup("disable").Value.String())
	accessMethods := flagSet.Lookup("access-method").Value.String()
	emailAddresses := flagSet.Lookup("email-addresses").Value.String()
	roles := flagSet.Lookup("roles").Value.String()

	if name == "" {
		fmt.Printf("\nError name argument is required\n")
		os.Exit(1)
	}
	createAdminGroupObject.Name = name

	if comment != "" {
		createAdminGroupObject.Comment = comment
	}
	if superUserErr == nil {
		createAdminGroupObject.SuperUser = &superUser
	}
	if disableErr == nil {
		createAdminGroupObject.Disable = &disable
	}
	if accessMethods != "" {
		createAdminGroupObject.AccessMethod = strings.Split(accessMethods, ",")
	}
	if emailAddresses != "" {
		createAdminGroupObject.EmailAddresses = strings.Split(emailAddresses, ",")
	}
	if roles != "" {
		createAdminGroupObject.Roles = strings.Split(roles, ",")
	}

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
	createAdminGroupFlags.String("name", "", "usage: -name admin-group-name")
	createAdminGroupFlags.String("comment", "", "usage: -comment 'A comment'")
	createAdminGroupFlags.String("super-user", "", "usage: -super-user (true|false)")
	createAdminGroupFlags.String("disable", "", "usage: -disable (true|false)")
	createAdminGroupFlags.String("access-method", "GUI,API,TAXII", "usage: -access-method method (One or more of API, CLOUD_API, GUI, TAXII")
	createAdminGroupFlags.String("email-addresses", "", "usage: -email-addresses emailaddress@domain,emailaddress2@domain....")
	createAdminGroupFlags.String("roles", "", "usage: -roles role1,role2...")
	RegisterCliCommand("admin-group-create", createAdminGroupFlags, createAdminGroup)
}
