package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"os"
)

func showPermission(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	permissionRef := flagSet.Lookup("ref").Value.String()
	if permissionRef == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	showPermission := permission.NewGet(permissionRef)
	err := client.Do(showPermission)
	httpStatus := showPermission.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving permission. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	response := showPermission.ResponseObject().(*permission.Permission)

	row := map[string]interface{}{}
	row["Permission"] = response.Permission
	row["Resource Type"] = response.ResourceType
	row["Role"] = response.Role
	row["Reference"] = response.Ref
	PrettyPrintSingle(row)

}

func init() {
	showPermissionFlags := flag.NewFlagSet("permission-show", flag.ExitOnError)
	showPermissionFlags.String("ref", "", "usage: -ref ")
	RegisterCliCommand("permission-show", showPermissionFlags, showPermission)
}
