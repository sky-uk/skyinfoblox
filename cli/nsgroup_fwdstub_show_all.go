package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"net/http"
	"os"
)

func showAllNSGroupFwdStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllNSGroupFwdStubAPI := nsgroupfwdstub.NewGetAll()
	err := client.Do(showAllNSGroupFwdStubAPI)
	httpStatus := showAllNSGroupFwdStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of Name Server Group Forward/Stub. HTTP status: %d. Error: %+v\n", httpStatus, string(showAllNSGroupFwdStubAPI.RawResponse()))
		os.Exit(1)
	}

	allNSGroupFwdStub := showAllNSGroupFwdStubAPI.ResponseObject().(*[]nsgroupfwdstub.NSGroupFwdStub)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, nsGroupFwdStubItem := range *allNSGroupFwdStub {
		row := map[string]interface{}{}
		row["Name"] = nsGroupFwdStubItem.Name
		row["Reference"] = nsGroupFwdStubItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllNSGroupFwdStubFlags := flag.NewFlagSet("nsgroup-fwd-stub-show-all", flag.ExitOnError)
	RegisterCliCommand("nsgroup-fwd-stub-show-all", showAllNSGroupFwdStubFlags, showAllNSGroupFwdStub)
}
