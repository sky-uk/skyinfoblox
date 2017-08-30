package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"fmt"
	"net/http"
)

func updateMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	mailExchanger := flagSet.Lookup("mail_exchanger").Value.String()
	preference := flagSet.Lookup("preference").Value.String()
	view := flagSet.Lookup("view").Value.String()
	returnFields := []string{"name", "mail_exchanger", "comment"}
	updateMxRecord := mxrecord.MxRecord{
		Ref: reference,
		Name: name,
		MailExchanger: mailExchanger,
		Preference: preference,
		View: view,
	}
	updateMXRecordAPI := mxrecord.NewUpdate(updateMxRecord,returnFields)
	updateErr := client.Do(updateMXRecordAPI)
	if updateErr != nil {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", updateErr.Error()))
	}
	if updateMXRecordAPI.StatusCode() != http.StatusCreated {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", *updateMXRecordAPI.ResponseObject().(*string)))
	} else {
		fmt.Println("MXRecord created")
		fmt.Println(*updateMXRecordAPI.ResponseObject().(*string))

	}

}

func init() {
	updateMXRecordFlags := flag.NewFlagSet("mxrecord-update", flag.ExitOnError)
	updateMXRecordFlags.String("ref", "", "usage: -ref reference for the record to delete")
	updateMXRecordFlags.String("name", "", "usage: -name name of the zone")
	updateMXRecordFlags.String("mail_exchanger", "", "usage: -name name of the zone")
	updateMXRecordFlags.String("preference", "", "usage: -preference preference of the mx record")
	updateMXRecordFlags.String("view", "", "usage: -view name of the view")
	RegisterCliCommand("mxrecord-update", updateMXRecordFlags, updateMXRecord)
}