package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func updateZone(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var dnsZone zoneauth.DNSZone
	dnsZone.Comment = flagSet.Lookup("comment").Value.String()
	zoneReference := flagSet.Lookup("ref").Value.String()

	updateZoneAuthAPI := zoneauth.NewUpdate(dnsZone, zoneReference)
	err := client.Do(updateZoneAuthAPI)
	if err != nil {
		spew.Dump("Error updating zone " + dnsZone.Reference + ": " + err.Error())
	}
	if updateZoneAuthAPI.StatusCode() == 200 {
		spew.Dump("Zone " + dnsZone.FQDN + " successfully updated")
		if client.Debug {
			spew.Dump(updateZoneAuthAPI.GetResponse())
		}
	} else {
		spew.Dump("Error status code != 200 when updating reference " + zoneReference)
		spew.Dump(updateZoneAuthAPI.GetResponse())
	}
}

func init() {
	var dnsZone zoneauth.DNSZone
	updateZoneFlags := flag.NewFlagSet("updatezone", flag.ExitOnError)
	updateZoneFlags.StringVar(&dnsZone.Comment, "comment", "", "usage: -comment 'My Comment'")
	updateZoneFlags.StringVar(&dnsZone.Reference, "ref", "", "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW")
	flag.Parse()
	RegisterCliCommand("updatezone", updateZoneFlags, updateZone)
}
