package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"net/http"
	"os"
)

func deleteNSGroupFwdStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	deleteNSGroupFwdStubAPI := nsgroupfwdstub.NewDelete(reference)
	err := client.Do(deleteNSGroupFwdStubAPI)
	httpStatus := deleteNSGroupFwdStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst deleting NS Group Forward/Stub reference %s. HTTP status: %d. Error: %+v\n", reference, httpStatus, string(deleteNSGroupFwdStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully deleted NS Group Forward/Stub reference %s\n", reference)
}

func init() {
	deleteNSGroupFwdStubFlags := flag.NewFlagSet("nsgroup-fwd-stub-delete", flag.ExitOnError)
	deleteNSGroupFwdStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REFERENCE")
	RegisterCliCommand("nsgroup-fwd-stub-delete", deleteNSGroupFwdStubFlags, deleteNSGroupFwdStub)
}
