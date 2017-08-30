package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"net/http"
	"os"
)

func deleteNSGroupStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	deleteNSGroupStubAPI := nsgroupstub.NewDelete(reference)
	err := client.Do(deleteNSGroupStubAPI)
	httpStatus := deleteNSGroupStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting NS Group Stub reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, string(deleteNSGroupStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted NS Group Stub reference %s\n", reference)
}

func init() {
	deleteNSGroupStubFlags := flag.NewFlagSet("nsgroup-stub-delete", flag.ExitOnError)
	deleteNSGroupStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REF")
	RegisterCliCommand("nsgroup-stub-delete", deleteNSGroupStubFlags, deleteNSGroupStub)
}
