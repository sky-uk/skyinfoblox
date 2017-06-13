package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func showAllZones(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllZoneAuthAPI := zoneauth.NewGetAllZones()
	err := client.Do(showAllZoneAuthAPI)
	if err != nil {
		spew.Dump("Error retrieving a list of all zones")
	}
	if showAllZoneAuthAPI.StatusCode() == 200 {
		allZoneReferences := showAllZoneAuthAPI.GetResponse()
		for _, zoneReference := range *allZoneReferences {
			spew.Dump("Zone FQDN: " + zoneReference.FQDN + " Reference: " + zoneReference.Reference)
		}
	} else {
		spew.Dump("Read All Zones return code != 200.")
	}
}

func init() {
	readAllZoneFlags := flag.NewFlagSet("zone-show-all", flag.ExitOnError)
	RegisterCliCommand("zone-show-all", readAllZoneFlags, showAllZones)
}
