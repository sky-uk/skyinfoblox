package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	//"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func updateZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var zoneDelegated zonedelegated.ZoneDelegated
	ref := flagSet.Lookup("ref").Value.String()
	returnFields := []string{"fqdn", "comment", "delegate_to"}
	getZoneDelegatedAPI := zonedelegated.NewGetZoneDelegated(ref, returnFields)
	getZoneErr := client.Do(getZoneDelegatedAPI)
	if getZoneErr != nil {
		fmt.Println(fmt.Sprintf("Error Getting the zone: %s", getZoneErr.Error()))
	}
	zoneDelegated = getZoneDelegatedAPI.ResponseObject().(zonedelegated.ZoneDelegated)
	if len(flagSet.Lookup("comment").Value.String()) > 0 {
		zoneDelegated.Comment = flagSet.Lookup("comment").Value.String()
	}

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
