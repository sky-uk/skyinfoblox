package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"os"
	"strings"
)

var updateAdminGroupObject admingroup.IBXAdminGroup
var updateAdminGroupAccessMethods, updateAdminGroupEmailAddresses, updateAdminGroupRoles string
var updateAdminGroupSuperUser, updateAdminGroupDisable bool

func updateAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if updateAdminGroupObject.Reference == "" {
		fmt.Printf("\nError ref argument is required\n")
		os.Exit(1)
	}

	if updateAdminGroupAccessMethods != "" {
		updateAdminGroupObject.AccessMethod = strings.Split(updateAdminGroupAccessMethods, ",")
	}
	if updateAdminGroupEmailAddresses != "" {
		updateAdminGroupObject.EmailAddresses = strings.Split(updateAdminGroupEmailAddresses, ",")
	}
	if updateAdminGroupRoles != "" {
		updateAdminGroupObject.Roles = strings.Split(updateAdminGroupRoles, ",")
	}
	updateAdminGroupObject.Disable = &updateAdminGroupDisable
	updateAdminGroupObject.SuperUser = &updateAdminGroupSuperUser

	updateAdminGroupAPI := admingroup.NewUpdate(updateAdminGroupObject, nil)
	err := client.Do(updateAdminGroupAPI)
	if err != nil {
		fmt.Printf("\nError updating admin group %s: %+v\n", updateAdminGroupObject.Name, err)
		os.Exit(2)
	}
	fmt.Printf("\nSuccessfully updated admin group %s\n", updateAdminGroupObject.Name)

}

func init() {
	updateAdminGroupFlags := flag.NewFlagSet("admin-group-update", flag.ExitOnError)
	updateAdminGroupFlags.StringVar(&updateAdminGroupObject.Reference, "ref", "", "usage: -ref")
	updateAdminGroupFlags.StringVar(&updateAdminGroupObject.Name, "name", "", "usage: -name admin-group-name")
	updateAdminGroupFlags.StringVar(&updateAdminGroupObject.Comment, "comment", "", "usage: -comment 'A comment'")
	updateAdminGroupFlags.BoolVar(&updateAdminGroupSuperUser, "super-user", false, "usage: -super-user")
	updateAdminGroupFlags.BoolVar(&updateAdminGroupDisable, "disable", false, "usage: -disable")
	updateAdminGroupFlags.StringVar(&updateAdminGroupAccessMethods, "access-method", "GUI,API,TAXII", "usage: -access-method method (One or more of API, CLOUD_API, GUI, TAXII")
	updateAdminGroupFlags.StringVar(&updateAdminGroupEmailAddresses, "email-addresses", "", "usage: -email-addresses emailaddress@domain,emailaddress2@domain....")
	updateAdminGroupFlags.StringVar(&updateAdminGroupRoles, "roles", "", "usage: -roles role1,role2...")
	RegisterCliCommand("admin-group-update", updateAdminGroupFlags, updateAdminGroup)
}
