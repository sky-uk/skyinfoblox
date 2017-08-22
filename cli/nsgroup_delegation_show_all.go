package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
	"net/http"
	"os"
)

func showAllNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllNSGroupDelegationAPI := nsgroupdelegation.NewGetAll()
	err := client.Do(showAllNSGroupDelegationAPI)
	httpStatus := showAllNSGroupDelegationAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of Name Server Group Delegations. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}

	allNSGroupDelegation := showAllNSGroupDelegationAPI.ResponseObject().(*[]nsgroupdelegation.NSGroupDelegation)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, nsGroupDelegationItem := range *allNSGroupDelegation {
		row := map[string]interface{}{}
		row["Name"] = nsGroupDelegationItem.Name
		row["Reference"] = nsGroupDelegationItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func init() {
	showAllNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-show-all", flag.ExitOnError)
	RegisterCliCommand("nsgroup-delegation-show-all", showAllNSGroupDelegationFlags, showAllNSGroupDelegation)
}
