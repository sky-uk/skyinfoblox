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
	createZoneDelegated.Fqdn = flagSet.Lookup("fqdn").Value.String()
	createDelegation := zonedelegated.NewCreate(createZoneDelegated)
	err := client.Do(createDelegation)
	if err != nil {
		fmt.Println(fmt.Printf("could not create delegation %s", err.Error()))
	}
	if createDelegation.StatusCode() == 201 {
		fmt.Println("Zone Delegated object created")
	}
	fmt.Println(createDelegation.StatusCode())
	fmt.Println(*createDelegation.ResponseObject().(*string))

}

func init() {
	createZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	createZoneDelegatedFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	createZoneDelegatedFlags.String("delegated-name", "", "usage: -delegated-name 'Delegated name server Name'")
	createZoneDelegatedFlags.String("delegated-address", "", "usage: -delegated-address 'Delegated server ip address'")
	createZoneDelegatedFlags.String("fqdn", "", "usage: -fqdn 'fqdn of the zone being delegated'")
	RegisterCliCommand("zone-delegated-create", createZoneDelegatedFlags, createZoneDelegated)
}
