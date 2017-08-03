package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"net/http"
	"os"
)

func showAllAdminRoles(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllAdminRole := adminrole.NewGetAll()
	err := client.Do(showAllAdminRole)
	httpStatus := showAllAdminRole.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of admin roles. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	allAdminRoles := showAllAdminRole.ResponseObject().(*[]adminrole.AdminRole)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, adminRoleItem := range *allAdminRoles {
		row := map[string]interface{}{}
		row["Name"] = adminRoleItem.Name
		row["Reference"] = adminRoleItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllAdminRolesFlags := flag.NewFlagSet("admin-role-show-all", flag.ExitOnError)
	RegisterCliCommand("admin-role-show-all", showAllAdminRolesFlags, showAllAdminRoles)
}
