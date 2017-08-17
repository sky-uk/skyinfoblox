package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
	//"net/http"
)

func updateZoneStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var updateZone zonestub.ZoneStub
	if flagSet.Lookup("ref").Value.String() != "" {
		updateZone.Ref = flagSet.Lookup("ref").Value.String()
	}

	if flagSet.Lookup("comment").Value.String() != "" {
		updateZone.Comment = flagSet.Lookup("comment").Value.String()
	}
	stubServers := common.ExternalServer{
		Address: flagSet.Lookup("stubaddress").Value.String(),
		Name:    flagSet.Lookup("stubname").Value.String(),
	}

	if flagSet.Lookup("stubservername").Value.String() != "" && flagSet.Lookup("stubserveraddress").Value.String() != "" {
		updateZone.StubFrom = []common.ExternalServer{stubServers}
	}
	updateAPI := zonestub.NewUpdate(updateZone)
	updateErr := client.Do(updateAPI)
	if updateErr != nil {
		fmt.Println("could not update the Zone")
	}
	fmt.Println("zone updated with success")

}

func init() {
	updateZoneStubFlags := flag.NewFlagSet("zone-delegated-update", flag.ExitOnError)
	updateZoneStubFlags.String("ref", "", "usage: -ref 'ref for the zone'")
	updateZoneStubFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	updateZoneStubFlags.String("stubservername", "", "usage: -stubservername 'Comment for the zone; maximum 256 characters'")
	updateZoneStubFlags.String("stubserveraddress", "", "usage: -stubserveraddress 'Comment for the zone; maximum 256 characters'")
	RegisterCliCommand("zone-stub-update", updateZoneStubFlags, updateZoneStub)
}
