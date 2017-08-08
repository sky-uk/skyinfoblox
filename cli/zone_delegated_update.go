package main


import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)


func updateZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	updateZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	updateZoneDelegatedFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	updateZoneDelegatedFlags.String("delegated-name", "", "usage: -delegated-name 'Delegated name server Name'")
	updateZoneDelegatedFlags.String("delegated-address", "", "usage: -delegated-address 'Delegated server ip address'")
	updateZoneDelegatedFlags.String("fqdn", "", "usage: -fqdn 'fqdn of the zone being delegated'")
	updateZoneDelegatedFlags.String("ref", "", "usage: -ref 'reference of the zone being updated'")
	RegisterCliCommand("zone-delegated-create", updateZoneDelegatedFlags, updateZoneDelegated)
}
