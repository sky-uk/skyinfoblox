package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func showZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	showZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	showZoneDelegatedFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	showZoneDelegatedFlags.String("delegated-name", "", "usage: -delegated-name 'Delegated name server Name'")
	showZoneDelegatedFlags.String("delegated-address", "", "usage: -delegated-address 'Delegated server ip address'")
	showZoneDelegatedFlags.String("fqdn", "", "usage: -fqdn 'fqdn of the zone being delegated'")
	RegisterCliCommand("zone-delegated-create", showZoneDelegatedFlags, showZoneDelegated)
}
