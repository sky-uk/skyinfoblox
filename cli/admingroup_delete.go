package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"net/http"
	"os"
)

func deleteAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	adminGroupReference := flagSet.Lookup("ref").Value.String()
	if adminGroupReference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	deleteAdminGroupAPI := admingroup.NewDelete(adminGroupReference)
	err := client.Do(deleteAdminGroupAPI)
	httpStatus := deleteAdminGroupAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting admin group %s. HTTP status: %d. Error: %+v\n", adminGroupReference, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted admin group %s\n", adminGroupReference)
}

func init() {
	deleteAdminGroupFlags := flag.NewFlagSet("admin-group-delete", flag.ExitOnError)
	deleteAdminGroupFlags.String("ref", "", "usage: -ref ")
	RegisterCliCommand("admin-group-delete", deleteAdminGroupFlags, deleteAdminGroup)
}
