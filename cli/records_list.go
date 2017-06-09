package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
)

type recordsListOptions struct {
	all bool
	txt bool
	srv bool
}

var (
	listOpts recordsListOptions
)

func recordsList(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if client.Debug {
		fmt.Println("listOptions:", listOpts)
	}

	if listOpts.all {
		listAllRecords(client)
	} else if listOpts.txt {
		listTXTRecords(client)
	} else if listOpts.srv {
		listSRVRecords(client)
	} else {
		fmt.Println("No record type specified.")
	}
}

func listAllRecords(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing All 'a' type of Records")
	}
	fields := []string{"name", "ipv4addr"}
	getAllARecordsAPI := records.NewGetAllARecords(fields)

	err := client.Do(getAllARecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getAllARecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getAllARecordsAPI.StatusCode())
		fmt.Println("Response: ", getAllARecordsAPI.ResponseObject())
	}

	headers := []string{"Name", "IPv4", "Ref"}
	rows := []map[string]interface{}{}
	for _, obj := range getAllARecordsAPI.GetResponse() {
		row := map[string]interface{}{}
		row["Name"] = obj.Name
		row["IPv4"] = obj.IPv4
		row["Ref"] = obj.Ref
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func listTXTRecords(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing All 'txt' type of Records")
	}
	getTXTRecordsAPI := records.NewGetAllTXTRecords([]string{})

	err := client.Do(getTXTRecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getTXTRecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getTXTRecordsAPI.StatusCode())
		fmt.Println("Response: ", getTXTRecordsAPI.ResponseObject())
	}

	headers := []string{"Name", "Text", "Ref"}
	rows := []map[string]interface{}{}
	for _, obj := range getTXTRecordsAPI.GetResponse() {
		row := map[string]interface{}{}
		row["Name"] = obj.Name
		row["Text"] = obj.Text
		row["Ref"] = obj.Ref
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func listSRVRecords(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing All 'srv' type of Records")
	}
	getSRVRecordsAPI := records.NewGetAllSRVRecords([]string{})

	err := client.Do(getSRVRecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getSRVRecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getSRVRecordsAPI.StatusCode())
		fmt.Println("Response: ", getSRVRecordsAPI.ResponseObject())
	}

	headers := []string{"Name", "Port", "Target", "Ref"}
	rows := []map[string]interface{}{}
	for _, obj := range getSRVRecordsAPI.GetResponse() {
		row := map[string]interface{}{}
		row["Name"] = obj.Name
		row["Port"] = obj.Port
		row["Target"] = obj.Target
		row["Ref"] = obj.Ref
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)
}

func init() {
	listFlags := flag.NewFlagSet("records-list", flag.ExitOnError)
	listFlags.BoolVar(&listOpts.all, "a", false, "List all a records")
	listFlags.BoolVar(&listOpts.txt, "txt", false, "List all txt records")
	listFlags.BoolVar(&listOpts.srv, "srv", false, "List all srv records")
	RegisterCliCommand("records-list", listFlags, recordsList)
}
