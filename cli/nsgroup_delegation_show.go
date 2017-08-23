package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
	"net/http"
	"os"
)

func showNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	nsGroupDelegationShowAPI := nsgroupdelegation.NewGet(reference, nsgroupdelegation.RequestReturnFields)
	err := client.Do(nsGroupDelegationShowAPI)
	httpStatus := nsGroupDelegationShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Delegation reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, err)
		os.Exit(1)
	}
	response := *nsGroupDelegationShowAPI.ResponseObject().(*nsgroupdelegation.NSGroupDelegation)
	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Delegate To Name"] = response.DelegateTo[0].Name
	row["Delegate To IP"] = response.DelegateTo[0].Address
	PrettyPrintSingle(row)
}

func init() {
	showNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-show", flag.ExitOnError)
	showNSGroupDelegationFlags.String("ref", "", "usage: -ref OBJECT_REF")
	RegisterCliCommand("nsgroup-delegation-show", showNSGroupDelegationFlags, showNSGroupDelegation)
}
