package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func createZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var createZoneDelegated zonedelegated.ZoneDelegated
	createZoneDelegated.View = "default"
	createZoneDelegated.ZoneFormat = "FORWARD"
	disable := false
	createZoneDelegated.Disable = &disable
	createZoneDelegated.DelegateTo = []common.ExternalServer{{Name: flagSet.Lookup("delegated-name").Value.String(), Address: flagSet.Lookup("delegated-address").Value.String()}}
	createZoneDelegated.Comment = flagSet.Lookup("comment").Value.String()
	createZoneDelegated.Address = flagSet.Lookup("address").Value.String()
	createDeletagion := zonedelegated.NewCreateZoneDelegated(createZoneDelegated)
	err := client.Do(createDeletagion)
	if err != nil {
		fmt.Println(fmt.Printf("could not create delegation %s", err.Error()))
	}
}

func init() {
	createZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	createZoneDelegatedFlags.String("address", "", "usage: -address 'The IP address of the server that is serving this zone'")
	createZoneDelegatedFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	createZoneDelegatedFlags.String("delegated-name", "", "usage: -delegated-name 'Delegated name server Name'")
	createZoneDelegatedFlags.String("delegated-address", "", "usage: -delegated-address 'Delegated server ip address'")
	RegisterCliCommand("zone-delegated-create", createZoneDelegatedFlags, createZoneDelegated)
}
