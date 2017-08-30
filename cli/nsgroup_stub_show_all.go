package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"net/http"
	"os"
)

func showAllNSGroupStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllNSGroupStubAPI := nsgroupstub.NewGetAll()
	err := client.Do(showAllNSGroupStubAPI)
	httpStatus := showAllNSGroupStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of Name Server Group Stub. HTTP status: %d. Error: %+v\n", httpStatus, string(showAllNSGroupStubAPI.RawResponse()))
		os.Exit(1)
	}

	allNSGroupStub := showAllNSGroupStubAPI.ResponseObject().(*[]nsgroupstub.NSGroupStub)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, nsGroupStubItem := range *allNSGroupStub {
		row := map[string]interface{}{}
		row["Name"] = nsGroupStubItem.Name
		row["Reference"] = nsGroupStubItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllNSGroupStubFlags := flag.NewFlagSet("nsgroup-stub-show-all", flag.ExitOnError)
	RegisterCliCommand("nsgroup-stub-show-all", showAllNSGroupStubFlags, showAllNSGroupStub)
}
