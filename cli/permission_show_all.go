package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"os"
)

func showAllPermissions(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllPermissions := permission.NewGetAll()
	err := client.Do(showAllPermissions)
	httpStatus := showAllPermissions.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of permissions. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	allPermissions := showAllPermissions.ResponseObject().(*[]permission.Permission)
	rows := []map[string]interface{}{}
	headers := []string{"Permission", "Resource Type", "Role", "Reference"}

	for _, permission := range *allPermissions {
		row := map[string]interface{}{}
		row["Permission"] = permission.Permission
		row["Resource Type"] = permission.ResourceType
		row["Role"] = permission.Role
		row["Reference"] = permission.Ref
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	readAllPermissionFlags := flag.NewFlagSet("permission-show-all", flag.ExitOnError)
	RegisterCliCommand("permission-show-all", readAllPermissionFlags, showAllPermissions)
}
