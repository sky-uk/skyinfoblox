package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
	"net/http"
	"os"
)

func deleteNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	deleteNSGroupDelegationAPI := nsgroupdelegation.NewDelete(reference)
	err := client.Do(deleteNSGroupDelegationAPI)
	httpStatus := deleteNSGroupDelegationAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting NS Group Delegation reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted NS Group Delegation reference %s\n", reference)
}

func init() {
	deleteNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-delete", flag.ExitOnError)
	deleteNSGroupDelegationFlags.String("ref", "", "usage: -ref")
	RegisterCliCommand("nsgroup-delegation-delete", deleteNSGroupDelegationFlags, deleteNSGroupDelegation)
}
