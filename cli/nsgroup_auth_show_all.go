package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupauth"
	"net/http"
	"os"
)

func showAllNSGroupAuth(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllNSGroupAuthAPI := nsgroupauth.NewGetAll()
	err := client.Do(showAllNSGroupAuthAPI)
	httpStatus := showAllNSGroupAuthAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of Name Server Group Auths. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}

	allNSGroupAuth := showAllNSGroupAuthAPI.ResponseObject().(*[]nsgroupauth.NSGroupAuth)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, nsGroupAuthItem := range *allNSGroupAuth {
		row := map[string]interface{}{}
		row["Name"] = nsGroupAuthItem.Name
		row["Reference"] = nsGroupAuthItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllNSGroupAuthFlags := flag.NewFlagSet("nsgroup-auth-show-all", flag.ExitOnError)
	RegisterCliCommand("nsgroup-auth-show-all", showAllNSGroupAuthFlags, showAllNSGroupAuth)
}
