package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/nameserver"
	"net/http"
	"os"
	"strings"
)

func createNSRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsRecord nameserver.NSRecord
	alwaysCreateAddressPTR := true

	nsRecord.Name = flagSet.Lookup("name").Value.String()
	nsRecord.NameServer = flagSet.Lookup("nameserver").Value.String()
	nsRecord.View = flagSet.Lookup("view").Value.String()

	if nsRecord.Name == "" {
		fmt.Printf("\nError: name argument required\n")
		os.Exit(1)
	}

	addressList := strings.Split(flagSet.Lookup("addresses").Value.String(), ",")
	addresses := make([]nameserver.ZoneNameServer, 0)
	for _, addressItem := range addressList {
		newAddress := nameserver.ZoneNameServer{addressItem, &alwaysCreateAddressPTR}
		addresses = append(addresses, newAddress)
	}
	nsRecord.Addresses = addresses

	createNSRecordAPI := nameserver.NewCreate(nsRecord)
	err := client.Do(createNSRecordAPI)
	httpStatus := createNSRecordAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst creating NS record %s. HTTP status: %d. Error: %+v\n", nsRecord.Name, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS record %s\n", nsRecord.Name)
}

func init() {
	createNSFlags := flag.NewFlagSet("record-ns-create", flag.ExitOnError)
	createNSFlags.String("name", "", "usage: -name zone-name")
	createNSFlags.String("nameserver", "", "usage: -nameserver name-server")
	createNSFlags.String("view", "default", "usage: -view the-view")
	createNSFlags.String("addresses", "", "usage: -addresses xxx.xxx.xxx.xxx,yyy.yyy.yyy.yyy (2x IP addresses")
	RegisterCliCommand("record-ns-create", createNSFlags, createNSRecord)
}
