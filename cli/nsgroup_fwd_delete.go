package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwd"
	"net/http"
	"os"
)

func deleteNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	deleteNSGroupFwdAPI := nsgroupfwd.NewDelete(reference)
	err := client.Do(deleteNSGroupFwdAPI)
	httpStatus := deleteNSGroupFwdAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting NS Group Forwarding reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, string(deleteNSGroupFwdAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted NS Group Forwarding reference %s\n", reference)
}

func init() {
	deleteNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-delete", flag.ExitOnError)
	deleteNSGroupFwdFlags.String("ref", "", "usage: -ref object-reference")
	RegisterCliCommand("nsgroup-fwd-delete", deleteNSGroupFwdFlags, deleteNSGroupFwd)
}
