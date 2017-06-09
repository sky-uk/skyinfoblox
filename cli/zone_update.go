package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

var dnsZone zoneauth.DNSZone

func zoneUpdate(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	dnsZone.Comment = flagSet.Lookup("comment").Value.String()
	dnsZone.Reference = flagSet.Lookup("ref").Value.String()

	updateZoneAuthAPI := zoneauth.NewUpdate(dnsZone)
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
		spew.Dump("Error status code != 200 when updating reference " + dnsZone.Reference)
		spew.Dump(updateZoneAuthAPI.GetResponse())
	}
}

func init() {
	zoneUpdateFlags := flag.NewFlagSet("zone-update", flag.ExitOnError)
	zoneUpdateFlags.StringVar(&dnsZone.Comment, "comment", "", "usage: -comment 'My Comment'")
	zoneUpdateFlags.StringVar(&dnsZone.Reference, "ref", "", "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW")
	flag.Parse()
	RegisterCliCommand("zone-update", zoneUpdateFlags, zoneUpdate)
}
