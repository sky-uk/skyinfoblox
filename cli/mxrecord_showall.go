package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
)

func showAllMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	getAllMXRecordAPI := mxrecord.NewGetAll()
	createErr := client.Do(getAllMXRecordAPI)
	if createErr != nil {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", createErr.Error()))
	}
	if getAllMXRecordAPI.StatusCode() != http.StatusOK {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", *getAllMXRecordAPI.ResponseObject().(*string)))
	} else {
		if getAllMXRecordAPI.StatusCode() == http.StatusOK {
			headers := []string{"Ref", "Name", "Mail_Exchanger"}
			rows := []map[string]interface{}{}
			for _, object := range *getAllMXRecordAPI.ResponseObject().(*[]mxrecord.MxRecord) {
				row := map[string]interface{}{}
				row["Ref"] = object.Ref
				row["Name"] = object.Name
				row["Mail_Exchanger"] = object.MailExchanger
				rows = append(rows, row)
			}

			PrettyPrintMany(headers, rows)
		} else {
			fmt.Println("Status code: ", getAllMXRecordAPI.StatusCode())
			fmt.Printf("Response:\n%s\n ", *getAllMXRecordAPI.ResponseObject().(*string))
		}

	}
}

func init() {
	showAllMXRecordFlags := flag.NewFlagSet("mxrecord-showall", flag.ExitOnError)
	RegisterCliCommand("mxrecord-showall", showAllMXRecordFlags, showAllMXRecord)
}
