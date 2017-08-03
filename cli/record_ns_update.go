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

func updateNSRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsRecord nameserver.NSRecord

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}
	nsRecord.Reference = reference

	nameServer := flagSet.Lookup("nameserver").Value.String()
	view := flagSet.Lookup("view").Value.String()
	addresses := flagSet.Lookup("addresses").Value.String()

	if nameServer != "" {
		nsRecord.NameServer = nameServer
	}
	if view != "" {
		nsRecord.View = view
	}
	if addresses != "" {
		alwaysCreateAddressPTR := true
		zoneNameServerObjects := make([]nameserver.ZoneNameServer, 0)
		addressList := strings.Split(addresses, ",")
		for _, address := range addressList {
			zoneNameServer := nameserver.ZoneNameServer{address, &alwaysCreateAddressPTR}
			zoneNameServerObjects = append(zoneNameServerObjects, zoneNameServer)
		}
		nsRecord.Addresses = zoneNameServerObjects
	}

	updateNSRecordAPI := nameserver.NewUpdate(nsRecord, nil)
	err := client.Do(updateNSRecordAPI)
	httpStatus := updateNSRecordAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst update NS record reference %s. HTTP Status: %d. Error: %+v\n", reference, httpStatus, reference)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated NS record reference %s\n", reference)
}

func init() {
	updateNSFlags := flag.NewFlagSet("record-ns-update", flag.ExitOnError)
	updateNSFlags.String("ref", "", "usage: -ref object-reference")
	updateNSFlags.String("nameserver", "", "usage: -nameserver name-server")
	updateNSFlags.String("view", "", "usage: -view the-view")
	updateNSFlags.String("addresses", "", "usage: -addresses xxx.xxx.xxx.xxx,yyy.yyy.yyy.yyy (2x IP addresses")
	RegisterCliCommand("record-ns-update", updateNSFlags, updateNSRecord)
}
