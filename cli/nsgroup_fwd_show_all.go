package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwd"
	"net/http"
	"os"
)

func showAllNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllNSGroupFwdAPI := nsgroupfwd.NewGetAll()
	err := client.Do(showAllNSGroupFwdAPI)
	httpStatus := showAllNSGroupFwdAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of Name Server Group Forwarding. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}

	allNSGroupFwd := showAllNSGroupFwdAPI.ResponseObject().(*[]nsgroupfwd.NSGroupFwd)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, nsGroupFwdItem := range *allNSGroupFwd {
		row := map[string]interface{}{}
		row["Name"] = nsGroupFwdItem.Name
		row["Reference"] = nsGroupFwdItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-show-all", flag.ExitOnError)
	RegisterCliCommand("nsgroup-fwd-show-all", showAllNSGroupFwdFlags, showAllNSGroupFwd)
}
