package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/nameserver"
	"net/http"
	"os"
)

func deleteNSRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	deleteNSAPI := nameserver.NewDelete(reference)
	err := client.Do(deleteNSAPI)
	httpStatus := deleteNSAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted %s\n", reference)
}

func init() {
	deleteNSFlags := flag.NewFlagSet("record-ns-delete", flag.ExitOnError)
	deleteNSFlags.String("ref", "", "usage: -ref reference-to-delete")
	RegisterCliCommand("record-ns-delete", deleteNSFlags, deleteNSRecord)
}
