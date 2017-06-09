package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func deleteZone(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	zoneReference := flagSet.Lookup("ref").Value.String()

	deleteZoneAuthAPI := zoneauth.NewDelete(zoneReference)
	err := client.Do(deleteZoneAuthAPI)
	if err != nil {
		spew.Dump("Error deleting zone reference " + zoneReference + err.Error())
	}
	if deleteZoneAuthAPI.StatusCode() == 200 {
		spew.Dump("Successfully deleted zone reference " + zoneReference)
		if client.Debug {
			spew.Dump(deleteZoneAuthAPI.ResponseObject())
		}
	} else {
		spew.Dump("Error status code != 200 when deleting zone reference " + zoneReference)
		spew.Dump(deleteZoneAuthAPI.ResponseObject())
	}
}

func init() {
	var zoneReference string
	deleteZoneFlags := flag.NewFlagSet("deletezone", flag.ExitOnError)
	deleteZoneFlags.StringVar(&zoneReference, "ref", "", "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW")
	flag.Parse()
	RegisterCliCommand("zone-delete", deleteZoneFlags, deleteZone)
}
