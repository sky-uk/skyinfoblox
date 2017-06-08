package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
)

type recordsListOptions struct {
	all bool
	txt bool
}

var (
	listOpts recordsListOptions
)

func recordsList(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	fields := []string{"name"}

	getAllARecordsAPI := records.NewGetAllARecords(fields)

	err := client.Do(getAllARecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getAllARecordsAPI.StatusCode() == 200 {
		if client.Debug {
			spew.Dump(getAllARecordsAPI.ResponseObject())
		}
	} else {
		fmt.Println("Status code: ", getAllARecordsAPI.StatusCode())
		fmt.Println("Response: ", getAllARecordsAPI.ResponseObject())
	}

	for _, obj := range *getAllARecordsAPI.GetResponse() {
		fmt.Println(obj)
	}

}

func init() {
	listFlags := flag.NewFlagSet("records-list", flag.ExitOnError)
	listFlags.BoolVar(&listOpts.all, "all", true, "List all records")
	listFlags.BoolVar(&listOpts.all, "txt", true, "List all records")
	RegisterCliCommand("records-list", listFlags, recordsList)
}
