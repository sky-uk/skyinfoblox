package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"os"
)

var zoneUpdateReferenceMessage = "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW"
var zoneUpdateDNSZone zoneauth.DNSZone

func zoneUpdate(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneUpdateDNSZone.Reference == "" {
		fmt.Println(zoneUpdateReferenceMessage)
		os.Exit(1)
	}

	updateZoneAuthAPI := zoneauth.NewUpdate(zoneUpdateDNSZone)
	err := client.Do(updateZoneAuthAPI)
	if err != nil {
		fmt.Println("Error updating zone " + zoneUpdateDNSZone.Reference + ": " + err.Error())
	}
	if updateZoneAuthAPI.StatusCode() == 200 {
		fmt.Println("Zone " + zoneUpdateDNSZone.FQDN + " successfully updated")
		if client.Debug {
			fmt.Println(updateZoneAuthAPI.GetResponse())
		}
	} else {
		fmt.Println("Error status code != 200 when updating reference " + zoneUpdateDNSZone.Reference)
		fmt.Println(updateZoneAuthAPI.GetResponse())
	}
}

func init() {
	zoneUpdateFlags := flag.NewFlagSet("zone-update", flag.ExitOnError)
	zoneUpdateFlags.StringVar(&zoneUpdateDNSZone.Comment, "comment", "", "usage: -comment 'My Comment'")
	zoneUpdateFlags.StringVar(&zoneUpdateDNSZone.Reference, "ref", "", zoneUpdateReferenceMessage)
	RegisterCliCommand("zone-update", zoneUpdateFlags, zoneUpdate)
}
