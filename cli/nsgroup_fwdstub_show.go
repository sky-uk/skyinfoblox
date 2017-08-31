package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"net/http"
	"os"
)

func showNSGroupFwdStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	nsGroupFwdStubShowAPI := nsgroupfwdstub.NewGet(reference, nsgroupfwdstub.RequestReturnFields)
	err := client.Do(nsGroupFwdStubShowAPI)
	httpStatus := nsGroupFwdStubShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Forward/Stub reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, string(nsGroupFwdStubShowAPI.RawResponse()))
		os.Exit(1)
	}
	// Note CLI only supports one item in each array
	response := *nsGroupFwdStubShowAPI.ResponseObject().(*nsgroupfwdstub.NSGroupFwdStub)
	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["External Server Name"] = response.ExternalServers[0].Name
	row["External Server IP"] = response.ExternalServers[0].Address
	PrettyPrintSingle(row)

}

func init() {
	showNSGroupFwdStubFlags := flag.NewFlagSet("nsgroup-fwd-stub-show", flag.ExitOnError)
	showNSGroupFwdStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REFERENCE")
	RegisterCliCommand("nsgroup-fwd-stub-show", showNSGroupFwdStubFlags, showNSGroupFwdStub)
}
