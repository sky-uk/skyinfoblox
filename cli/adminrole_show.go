package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"net/http"
	"os"
)

func showAdminRole(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	adminRoleReference := flagSet.Lookup("ref").Value.String()
	if adminRoleReference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	adminRoleShowAPI := adminrole.NewGet(adminRoleReference)
	err := client.Do(adminRoleShowAPI)
	httpStatus := adminRoleShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving admin role reference %s. HTTP status: %d. Error: %+v", adminRoleReference, httpStatus, err)
		os.Exit(2)
	}
	response := *adminRoleShowAPI.ResponseObject().(*adminrole.AdminRole)

	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Disabled"] = *response.Disable
	PrettyPrintSingle(row)
}

func init() {
	showAdminRoleFlags := flag.NewFlagSet("admin-role-show", flag.ExitOnError)
	showAdminRoleFlags.String("ref", "", "usage: -ref ")
	RegisterCliCommand("admin-role-show", showAdminRoleFlags, showAdminRole)
}
