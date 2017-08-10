package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"net/http"
	"os"
)

var noZoneForwardRef = "usage: -ref zone_forward/XXXXXXXX:FQDN/VIEW"
var zoneToUpdate zoneforward.ZoneForward

func forwardZoneUpdate(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneToUpdate.Ref == "" {
		fmt.Println(noZoneForwardRef)
		os.Exit(1)
	}

	returnFields := []string{"address", "comment", "fqdn"}
	api := zoneforward.NewUpdate(zoneToUpdate, returnFields)
	err := client.Do(api)
	if err != nil {
		fmt.Println("Error updating forward zone " + zoneToUpdate.Ref + ": " + err.Error())
	}
	if api.StatusCode() == http.StatusOK {
		fmt.Println("Zone " + zoneToUpdate.Fqdn + " successfully updated")
		if client.Debug {
			response := *api.ResponseObject().(*zoneforward.ZoneForward)
			fmt.Printf("%v", response)
		}
	} else {
		fmt.Println("Error status code != 200 when updating reference " + zoneToUpdate.Ref)
		fmt.Printf("Error:\n%+v\n", string(api.RawResponse()))
	}
}

func init() {
	flags := flag.NewFlagSet("zoneforward-update", flag.ExitOnError)
	flags.StringVar(&zoneToUpdate.Comment, "comment", "", "usage: -comment 'My Comment'")
	flags.StringVar(&zoneToUpdate.Ref, "ref", "", noZoneForwardRef)
	RegisterCliCommand("zoneforward-update", flags, forwardZoneUpdate)
}
