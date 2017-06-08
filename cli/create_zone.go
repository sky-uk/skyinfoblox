package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func createZone(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var dnsZone zoneauth.DNSZone
	dnsZone.FQDN = flagSet.Lookup("fqdn").Value.String()
	dnsZone.View = flagSet.Lookup("view").Value.String()
	dnsZone.Comment = flagSet.Lookup("comment").Value.String()

	createZoneAuthAPI := zoneauth.NewCreate(dnsZone)
	err := client.Do(createZoneAuthAPI)
	if err != nil {
		spew.Dump("Error creating new zone " + dnsZone.FQDN + ": " + err.Error())
	}
	if createZoneAuthAPI.StatusCode() == 201 {
		spew.Dump("Zone " + dnsZone.FQDN + " successfully created")
		if client.Debug {
			spew.Dump(createZoneAuthAPI.GetResponse())
		}
	} else {
		spew.Dump("Error status code != 201 when creating zone " + dnsZone.FQDN)
		spew.Dump(createZoneAuthAPI.GetResponse())
	}
}

func init() {
	var dnsZone zoneauth.DNSZone
	createZoneFlags := flag.NewFlagSet("createzone", flag.ExitOnError)
	createZoneFlags.StringVar(&dnsZone.FQDN, "fqdn", "", "usage: -fqdn mydomain.com")
	createZoneFlags.StringVar(&dnsZone.View, "view", "default", "usage: -view default")
	createZoneFlags.StringVar(&dnsZone.Comment, "comment", "", "usage: -comment 'My Comment'")
	flag.Parse()
	RegisterCliCommand("createzone", createZoneFlags, createZone)
}
