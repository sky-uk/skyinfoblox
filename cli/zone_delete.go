package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"os"
)

var zoneDeleteReference string
var zoneDeleteUsageMessage = "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW"

func deleteZone(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneDeleteReference == "" {
		fmt.Println(zoneDeleteUsageMessage)
		os.Exit(1)
	}

	deleteZoneAuthAPI := zoneauth.NewDelete(zoneDeleteReference)
	err := client.Do(deleteZoneAuthAPI)
	if err != nil {
		fmt.Println("Error deleting zone reference: " + zoneDeleteReference + err.Error())
	}
	if deleteZoneAuthAPI.StatusCode() == 200 {
		fmt.Println("Successfully deleted zone reference: " + zoneDeleteReference)
		if client.Debug {
			fmt.Println(deleteZoneAuthAPI.GetResponse())
		}
	} else {
		fmt.Println("Error status code != 200 when deleting zone reference: " + zoneDeleteReference)
		fmt.Println(deleteZoneAuthAPI.ResponseObject())
	}
}

func init() {
	deleteZoneFlags := flag.NewFlagSet("deletezone", flag.ExitOnError)
	deleteZoneFlags.StringVar(&zoneDeleteReference, "ref", "", zoneDeleteUsageMessage)
	RegisterCliCommand("zone-delete", deleteZoneFlags, deleteZone)
}
