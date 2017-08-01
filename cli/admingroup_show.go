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

func showAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	returnFields := []string{"name", "comment", "disable", "roles", "email_addresses", "superuser", "access_method"}

	adminGroupReference := flagSet.Lookup("ref").Value.String()
	if adminGroupReference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	adminGroupShowAPI := admingroup.NewGet(adminGroupReference, returnFields)
	err := client.Do(adminGroupShowAPI)
	httpStatus := adminGroupShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving admin group reference %s. HTTP status: %d. Error: %+v", adminGroupReference, httpStatus, err)
		os.Exit(2)
	}
	response := *adminGroupShowAPI.ResponseObject().(*admingroup.IBXAdminGroup)

	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Superuser"] = *response.SuperUser
	row["Disabled"] = *response.Disable
	row["Roles"] = strings.Join(response.Roles, ", ")
	row["Email Addresses"] = strings.Join(response.EmailAddresses, ", ")
	row["Access Method"] = strings.Join(response.AccessMethod, ", ")
	PrettyPrintSingle(row)

}

func init() {
	showAdminGroupFlags := flag.NewFlagSet("admin-group-show", flag.ExitOnError)
	showAdminGroupFlags.String("ref", "", "usage: -ref ")
	RegisterCliCommand("admin-group-show", showAdminGroupFlags, showAdminGroup)
}
