package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"net/http"
	"os"
)

func showAllAdminGroup(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllAdminGroup := admingroup.NewGetAll()
	err := client.Do(showAllAdminGroup)
	httpStatus := showAllAdminGroup.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving a list of admin groups. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	allAdminGroups := showAllAdminGroup.ResponseObject().(*[]admingroup.IBXAdminGroupReference)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, adminGroupItem := range *allAdminGroups {
		row := map[string]interface{}{}
		row["Name"] = adminGroupItem.AdminGroupName
		row["Reference"] = adminGroupItem.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	showAllAdminGroupFlags := flag.NewFlagSet("admin-group-show-all", flag.ExitOnError)
	RegisterCliCommand("admin-group-show-all", showAllAdminGroupFlags, showAllAdminGroup)
}
