package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
)

func deleteMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	deleteMXRecordAPI := mxrecord.NewDelete(ref)
	createErr := client.Do(deleteMXRecordAPI)
	if createErr != nil {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", createErr.Error()))
	}
	if deleteMXRecordAPI.StatusCode() != http.StatusOK {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", *deleteMXRecordAPI.ResponseObject().(*string)))
	} else {
		fmt.Println("MXRecord Deleted")
		fmt.Println(*deleteMXRecordAPI.ResponseObject().(*string))

	}
}

func init() {
	deleteMXRecordFlags := flag.NewFlagSet("mxrecord-create", flag.ExitOnError)
	deleteMXRecordFlags.String("ref", "", "usage: -ref reference for the record to delete")
	RegisterCliCommand("mxrecord-delete", deleteMXRecordFlags, deleteMXRecord)
}
