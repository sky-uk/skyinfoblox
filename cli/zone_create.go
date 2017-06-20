package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"os"
)

var zoneCreateDNSZone zoneauth.DNSZone
var zoneCreateFQDNMessage = "usage: -fqdn mydomain.com"

func createZone(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneCreateDNSZone.FQDN == "" {
		fmt.Println(zoneCreateFQDNMessage)
		os.Exit(1)
	}

	createZoneAuthAPI := zoneauth.NewCreate(zoneCreateDNSZone)
	err := client.Do(createZoneAuthAPI)
	if err != nil {
		fmt.Println("Error creating new zone " + zoneCreateDNSZone.FQDN + ": " + err.Error())
	}
	if createZoneAuthAPI.StatusCode() == 201 {
		fmt.Println("Zone " + zoneCreateDNSZone.FQDN + " successfully created")
		if client.Debug {
			response := createZoneAuthAPI.GetResponse()
			fmt.Printf("%s", response)
		}
	} else {
		fmt.Printf("\nError status code was %d when attempting to creating zone %s.\n ", createZoneAuthAPI.StatusCode(), zoneCreateDNSZone.FQDN)
		fmt.Printf("Response: %s\n", createZoneAuthAPI.GetResponse())
	}
}

func init() {
	// Setting zoneUpdateDNSZone.SOADefaultTTL on creation is ignored by Infoblox.
	createZoneFlags := flag.NewFlagSet("zonecreate", flag.ExitOnError)
	createZoneFlags.StringVar(&zoneCreateDNSZone.FQDN, "fqdn", "", zoneCreateFQDNMessage)
	createZoneFlags.StringVar(&zoneCreateDNSZone.View, "view", "default", "usage: -view default")
	createZoneFlags.StringVar(&zoneCreateDNSZone.Comment, "comment", "", "usage: -comment 'My Comment'")
	createZoneFlags.StringVar(&zoneCreateDNSZone.ZoneFormat, "zone-format", "FORWARD", "usage: -zone-format (FORWARD|IPV4|IPV6)")
	RegisterCliCommand("zone-create", createZoneFlags, createZone)
}
