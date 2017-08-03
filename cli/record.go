package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"github.com/sky-uk/skyinfoblox/api/records/nameserver"
	"net/http"
	"os"
)

var recordRef string
var recordType string

func listRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if recordRef == "" {
		fmt.Println("Error: No record reference specified.")
		return
	}

	switch recordType {
	case "a":
		listARecord(client)
	case "txt":
		listTXTRecord(client)
	case "srv":
		listSRVRecord(client)
	case "ns":
		listNSRecord(client)
	default:
		fmt.Println("Error: No or wrong record type specified.")
	}
}

func listARecord(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing single 'a' type record of reference:", recordRef)
	}
	fields := []string{"name", "ipv4addr"}
	getSingleARecordsAPI := records.NewGetARecord(recordRef, fields)

	err := client.Do(getSingleARecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getSingleARecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getSingleARecordsAPI.StatusCode())
		fmt.Println("Response: ", getSingleARecordsAPI.ResponseObject())
	}

	record := getSingleARecordsAPI.GetResponse()
	row := map[string]interface{}{}
	row["Name"] = record.Name
	row["IPv4"] = record.IPv4
	row["Ref"] = record.Ref
	PrettyPrintSingle(row)

}

func listTXTRecord(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing single 'txt' type record of reference:", recordRef)
	}
	fields := []string{"name", "text"}
	getSingleTXTRecordsAPI := records.NewGetTXTRecord(recordRef, fields)

	err := client.Do(getSingleTXTRecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getSingleTXTRecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getSingleTXTRecordsAPI.StatusCode())
		fmt.Println("Response: ", getSingleTXTRecordsAPI.ResponseObject())
	}

	record := getSingleTXTRecordsAPI.GetResponse()
	row := map[string]interface{}{}
	row["Name"] = record.Name
	row["Text"] = record.Text
	row["Ref"] = record.Ref
	PrettyPrintSingle(row)
}

func listSRVRecord(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing single 'srv' type record of reference:", recordRef)
	}
	fields := []string{"name", "port", "target"}
	getSingleSRVRecordsAPI := records.NewGetSRVRecord(recordRef, fields)

	err := client.Do(getSingleSRVRecordsAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if getSingleSRVRecordsAPI.StatusCode() != 200 {
		fmt.Println("Status code: ", getSingleSRVRecordsAPI.StatusCode())
		fmt.Println("Response: ", getSingleSRVRecordsAPI.ResponseObject())
	}

	record := getSingleSRVRecordsAPI.GetResponse()
	row := map[string]interface{}{}
	row["Name"] = record.Name
	row["Port"] = record.Port
	row["Target"] = record.Target
	row["Priority"] = record.Priority
	row["Ref"] = record.Ref
	PrettyPrintSingle(row)
}

func listNSRecord(client *skyinfoblox.InfobloxClient) {
	if client.Debug {
		fmt.Println("Listing single 'ns' type record of reference:", recordRef)
	}
	fields := []string{"name", "nameserver", "view", "addresses"}
	getNSRecordAPI := nameserver.NewGet(recordRef, fields)

	err := client.Do(getNSRecordAPI)
	httpStatus := getNSRecordAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS record %s. HTTP status: %d. Error: %+v\n", recordRef, httpStatus, err)
		os.Exit(1)
	}

	record := *getNSRecordAPI.ResponseObject().(*nameserver.NSRecord)
	row := map[string]interface{}{}
	row["Name"] = record.Name
	row["Reference"] = record.Reference
	row["Name Server"] = record.NameServer
	row["View"] = record.View

	var addressString string
	addresses := record.Addresses
	for idx, address := range addresses {
		if idx == 0 {
			addressString = fmt.Sprintf("[%s, %t]", address.Address, *address.AutoCreatePointerRecord)
		} else {
			addressString = fmt.Sprintf("%s, [%s, %t]", addressString, address.Address, *address.AutoCreatePointerRecord)
		}
	}
	row["Addresses [Address, PTR]"] = addressString
	PrettyPrintSingle(row)
}

func init() {
	listFlags := flag.NewFlagSet("record", flag.ExitOnError)
	listFlags.StringVar(&recordRef, "ref", "", "Reference of the record to get")
	listFlags.StringVar(&recordType, "type", "", "Type of the record to get. i.e  txt,a,srv")
	RegisterCliCommand("record", listFlags, listRecord)
}
