package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"os"
)

var zoneForwardDeleteReference string
var zoneForwardDeleteUsageMessage = "Usage: -ref zone_forward/XXXXXXXX:FQDN/VIEW"

func deleteZoneForward(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneForwardDeleteReference == "" {
		fmt.Println(zoneForwardDeleteUsageMessage)
		os.Exit(1)
	}

	api := zoneforward.NewDelete(zoneForwardDeleteReference)
	err := client.Do(api)
	if err != nil {
		fmt.Println("Error deleting forward zone reference: " + zoneForwardDeleteReference + err.Error())
	}
	if api.StatusCode() == 200 {
		fmt.Println("Successfully deleted forward zone reference: " + zoneForwardDeleteReference)
		if client.Debug {
			fmt.Println(api.ResponseObject().(string))
		}
	} else {
		fmt.Println("Error status code != 200 when deleting forward zone reference: " + zoneForwardDeleteReference)
		fmt.Println(api.ResponseObject().(string))
	}
}

func init() {
	deleteZoneForwardFlags := flag.NewFlagSet("zoneforward-delete", flag.ExitOnError)
	deleteZoneForwardFlags.StringVar(&zoneForwardDeleteReference, "ref", "", zoneForwardDeleteUsageMessage)
	RegisterCliCommand("zoneforward-delete", deleteZoneForwardFlags, deleteZoneForward)
}
