package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"net/http"
	"os"
)

func showNSGroupStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	nsGroupStubShowAPI := nsgroupstub.NewGet(reference, nsgroupstub.RequestReturnFields)
	err := client.Do(nsGroupStubShowAPI)
	httpStatus := nsGroupStubShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Stub reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, string(nsGroupStubShowAPI.RawResponse()))
		os.Exit(1)
	}
	// Note CLI only supports one item in each array
	response := *nsGroupStubShowAPI.ResponseObject().(*nsgroupstub.NSGroupStub)
	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Grid Member Name"] = response.StubMembers[0].Name

	PrettyPrintSingle(row)

}

func init() {
	showNSGroupStubFlags := flag.NewFlagSet("nsgroup-stub-show", flag.ExitOnError)
	showNSGroupStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REF")
	RegisterCliCommand("nsgroup-stub-show", showNSGroupStubFlags, showNSGroupStub)
}
