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

func updateAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var adminGroup admingroup.IBXAdminGroup

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}
	adminGroup.Reference = reference

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	superUser, superUserErr := strconv.ParseBool(flagSet.Lookup("super-user").Value.String())
	disable, disableErr := strconv.ParseBool(flagSet.Lookup("disable").Value.String())
	accessMethods := flagSet.Lookup("access-method").Value.String()
	emailAddresses := flagSet.Lookup("email-addresses").Value.String()
	roles := flagSet.Lookup("roles").Value.String()

	if name != "" {
		adminGroup.Name = name
	}
	if comment != "" {
		adminGroup.Comment = comment
	}
	if superUserErr == nil {
		adminGroup.SuperUser = &superUser
	}
	if disableErr == nil {
		adminGroup.Disable = &disable
	}
	if accessMethods != "" {
		adminGroup.AccessMethod = strings.Split(accessMethods, ",")
	}
	if emailAddresses != "" {
		adminGroup.EmailAddresses = strings.Split(emailAddresses, ",")
	}
	if roles != "" {
		adminGroup.Roles = strings.Split(roles, ",")
	}

	updateAdminGroupAPI := admingroup.NewUpdate(adminGroup, nil)
	err := client.Do(updateAdminGroupAPI)
	httpStatus := updateAdminGroupAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating admin group %s. HTTP status: %d. Error: %+v\n", adminGroup.Reference, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated admin group %s\n", adminGroup.Reference)

}

func init() {
	updateAdminGroupFlags := flag.NewFlagSet("admin-group-update", flag.ExitOnError)
	updateAdminGroupFlags.String("ref", "", "usage: -ref")
	updateAdminGroupFlags.String("name", "", "usage: -name admin-group-name")
	updateAdminGroupFlags.String("comment", "", "usage: -comment 'A comment'")
	updateAdminGroupFlags.String("super-user", "", "usage: -super-user (true|false)")
	updateAdminGroupFlags.String("disable", "", "usage: -disable (true|false)")
	updateAdminGroupFlags.String("access-method", "GUI,API,TAXII", "usage: -access-method method (One or more of API, CLOUD_API, GUI, TAXII")
	updateAdminGroupFlags.String("email-addresses", "", "usage: -email-addresses emailaddress@domain,emailaddress2@domain....")
	updateAdminGroupFlags.String("roles", "", "usage: -roles role1,role2...")
	RegisterCliCommand("admin-group-update", updateAdminGroupFlags, updateAdminGroup)
}
