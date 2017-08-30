package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
)

func showMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	returnFields := []string{"name", "mail_exchanger", "comment"}
	getMXRecordAPI := mxrecord.NewGet(reference, returnFields)
	createErr := client.Do(getMXRecordAPI)
	if createErr != nil {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", createErr.Error()))
	}
	if getMXRecordAPI.StatusCode() != http.StatusOK {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", *getMXRecordAPI.ResponseObject().(*string)))
	} else {
		fmt.Println("Avaliable MXRecords")
		fmt.Println(*getMXRecordAPI.ResponseObject().(*mxrecord.MxRecord))

	}

}

func init() {
	showMXRecordFlags := flag.NewFlagSet("mxrecord-show", flag.ExitOnError)
	showMXRecordFlags.String("ref", "", "usage: -ref reference for the record to delete")
	RegisterCliCommand("mxrecord-show", showMXRecordFlags, showMXRecord)
}
