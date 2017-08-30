package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
	"strconv"
)

func createMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	mailExchanger := flagSet.Lookup("mail_exchanger").Value.String()
	name := flagSet.Lookup("name").Value.String()
	preference := flagSet.Lookup("preference").Value.String()
	view := flagSet.Lookup("view").Value.String()
	var mxRecord mxrecord.MxRecord

	mxRecord.MailExchanger = mailExchanger
	mxRecord.Name = name
	Preference, _ := strconv.Atoi(preference)
	mxRecord.Preference = uint(Preference)
	if view != "" {
		mxRecord.View = view
	}
	createMXRecordAPI := mxrecord.NewCreate(mxRecord)
	createErr := client.Do(createMXRecordAPI)
	if createErr != nil {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", createErr.Error()))
	}
	if createMXRecordAPI.StatusCode() != http.StatusCreated {
		fmt.Println(fmt.Errorf("Error creating MXRecord: %s", *createMXRecordAPI.ResponseObject().(*string)))
	} else {
		fmt.Println("MXRecord created")
		fmt.Println(*createMXRecordAPI.ResponseObject().(*string))

	}
}

func init() {
	createMXRecordFlags := flag.NewFlagSet("mxrecord-create", flag.ExitOnError)
	createMXRecordFlags.String("mail_exchanger", "", "usage: -mail_exchanger mail exchanger")
	createMXRecordFlags.String("name", "", "usage: -name fqdn of the domain for this MXRecord")
	createMXRecordFlags.String("preference", "", "usage: -preference prefenrece for the MX record")
	createMXRecordFlags.String("view", "", "usage: -view view for the MX record")
	RegisterCliCommand("mxrecord-create", createMXRecordFlags, createMXRecord)
}
