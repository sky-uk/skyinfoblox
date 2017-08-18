package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupauth"
	"net/http"
	"os"
)

func deleteNSGroupAuth(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	deleteNSGroupAuthAPI := nsgroupauth.NewDelete(reference)
	err := client.Do(deleteNSGroupAuthAPI)
	httpStatus := deleteNSGroupAuthAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting NS Group Auth reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, err)
		os.Exit(1)
	}

	fmt.Printf("\nSuccessfully deleted NS Group Auth reference %s\n", reference)
}

func init() {
	deleteNSGroupAuthFlags := flag.NewFlagSet("nsgroup-auth-delete", flag.ExitOnError)
	deleteNSGroupAuthFlags.String("ref", "", "usage: -ref")
	RegisterCliCommand("nsgroup-auth-delete", deleteNSGroupAuthFlags, deleteNSGroupAuth)
}
