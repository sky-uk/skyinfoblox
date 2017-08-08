package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func updateZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var zoneDelegated zonedelegated.ZoneDelegated
	ref := flagSet.Lookup("ref").Value.String()
	returnFields := []string{"comment", "delegate_to"}
	getZoneDelegatedAPI := zonedelegated.NewGetZoneDelegated(ref, returnFields)
	getZoneErr := client.Do(getZoneDelegatedAPI)
	if getZoneErr != nil {
		fmt.Println(fmt.Sprintf("Error Getting the zone: %s", getZoneErr.Error()))
	}
	zoneDelegated = *getZoneDelegatedAPI.ResponseObject().(*zonedelegated.ZoneDelegated)
	if flagSet.Lookup("comment").Value.String() != "" {
		zoneDelegated.Comment = flagSet.Lookup("comment").Value.String()
	}
	if flagSet.Lookup("delegated-name").Value.String() != "" && flagSet.Lookup("delegated-address").Value.String() != "" {
		delegation := common.ExternalServer{
			Name:    flagSet.Lookup("delegated-name").Value.String(),
			Address: flagSet.Lookup("delegated-address").Value.String(),
		}
		zoneDelegated.DelegateTo = []common.ExternalServer{delegation}
	}

	updateZoneAPI := zonedelegated.NewUpdateZoneDelegated(ref, zoneDelegated)
	updateErr := client.Do(updateZoneAPI)
	if updateErr != nil {
		fmt.Println("could not update")
	}
	fmt.Println(updateZoneAPI.StatusCode())
	fmt.Println(*updateZoneAPI.ResponseObject().(*string))

}

func init() {
	updateZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-update", flag.ExitOnError)
	updateZoneDelegatedFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	updateZoneDelegatedFlags.String("delegated-name", "", "usage: -delegated-name 'Delegated name server Name'")
	updateZoneDelegatedFlags.String("delegated-address", "", "usage: -delegated-address 'Delegated server ip address'")
	updateZoneDelegatedFlags.String("fqdn", "", "usage: -fqdn 'fqdn of the zone being delegated'")
	updateZoneDelegatedFlags.String("ref", "", "usage: -ref 'reference of the zone being updated'")
	RegisterCliCommand("zone-delegated-update", updateZoneDelegatedFlags, updateZoneDelegated)
}
