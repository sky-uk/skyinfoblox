package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
	"strconv"
)

func updateMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	mailExchanger := flagSet.Lookup("mail_exchanger").Value.String()
	preference := flagSet.Lookup("preference").Value.String()
	Preference, _ := strconv.Atoi(preference)
	view := flagSet.Lookup("view").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	var updateMxRecord mxrecord.MxRecord
	if name != "" {
		updateMxRecord.Name = name
	}

	if mailExchanger != "" {
		updateMxRecord.MailExchanger = mailExchanger
	}
	if view != "" {
		updateMxRecord.View = view
	}
	if preference != "" {
		updateMxRecord.Preference = uint(Preference)
	}
	if comment != "" {
		updateMxRecord.Comment = comment
	}
	updateMXRecordAPI := mxrecord.NewUpdate(reference, updateMxRecord)
	updateErr := client.Do(updateMXRecordAPI)
	if updateErr != nil {
		fmt.Println(fmt.Errorf("Error updating MXRecord: %s", updateErr.Error()))
	}
	if updateMXRecordAPI.StatusCode() != http.StatusOK {
		fmt.Println(fmt.Errorf("Error updating MXRecord: %s", *updateMXRecordAPI.ResponseObject().(*string)))
	} else {
		fmt.Println("MXRecord updated")
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
	updateMXRecordFlags.String("comment", "", "usage: -comment a comment")
	RegisterCliCommand("mxrecord-update", updateMXRecordFlags, updateMXRecord)
}
