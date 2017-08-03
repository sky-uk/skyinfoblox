package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"os"
)

func deletePermission(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	permissionRef := flagSet.Lookup("ref").Value.String()
	if permissionRef == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	deletePermission := permission.NewDelete(permissionRef)
	err := client.Do(deletePermission)
	httpStatus := deletePermission.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting permission. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted admin group %s\n", permissionRef)
}

func init() {
	deletePermissionFlags := flag.NewFlagSet("permission-delete", flag.ExitOnError)
	deletePermissionFlags.String("ref", "", "usage: -ref ")
	RegisterCliCommand("permission-delete", deletePermissionFlags, deletePermission)
}
