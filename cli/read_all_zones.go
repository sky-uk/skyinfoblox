package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func readAllZones(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	readAllZoneAuthAPI := zoneauth.NewGetAll()
	err := client.Do(readAllZoneAuthAPI)
	if err != nil {
		spew.Dump("Error retrieving a list of all zones")
	}
	if readAllZoneAuthAPI.StatusCode() == 200 {
		allZoneReferences := readAllZoneAuthAPI.GetResponse()
		for _, zoneReference := range *allZoneReferences {
			spew.Dump("Zone FQDN: " + zoneReference.FQDN + " Reference: " + zoneReference.Reference)
		}
	} else {
		spew.Dump("Read All Zones return code != 200.")
	}
}

func init() {
	readAllZoneFlags := flag.NewFlagSet("read-all-zones", flag.ExitOnError)
	flag.Parse()
	RegisterCliCommand("read-all-zones", readAllZoneFlags, readAllZones)
}
