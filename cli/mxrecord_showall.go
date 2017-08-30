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
		fmt.Println("Avaliable MXRecords")
		fmt.Println(*getAllMXRecordAPI.ResponseObject().(*[]mxrecord.MxRecord))

	}
}

func init() {
	showAllMXRecordFlags := flag.NewFlagSet("mxrecord-showall", flag.ExitOnError)
	RegisterCliCommand("mxrecord-showall", showAllMXRecordFlags, showAllMXRecord)
}
