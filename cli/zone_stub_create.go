package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
)

func createZoneStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var stubZone zonestub.ZoneStub
	stubServers := common.ExternalServer{
		Address: flagSet.Lookup("stubaddress").Value.String(),
		Name:    flagSet.Lookup("stubname").Value.String(),
	}
	stubZone.View = "default"
	stubZone.Comment = flagSet.Lookup("comment").Value.String()
	stubZone.StubFrom = []common.ExternalServer{stubServers}
	stubZone.FQDN = flagSet.Lookup("fqdn").Value.String()
	createAPI := zonestub.NewCreate(stubZone)
	createErr := client.Do(createAPI)
	if createErr != nil {
		fmt.Println(createErr.Error())
	}
	if createAPI.StatusCode() != 200 {
		fmt.Println(createAPI.StatusCode())
		fmt.Println(*createAPI.ResponseObject().(*string))
	} else {
		fmt.Println("Create Stub created !!")
	}

}

func init() {
	createZoneStubFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	createZoneStubFlags.String("comment", "", "usage: -comment 'Comment for the zone; maximum 256 characters'")
	createZoneStubFlags.String("fqdn", "", "usage: -fqdn 'fqdn for the zone'")
	createZoneStubFlags.String("stubaddress", "", "usage: -stubaddress 'stub_to server ip address'")
	createZoneStubFlags.String("stubname", "", "usage: -stubname 'stub_to server name'")
	RegisterCliCommand("zone-stub-create", createZoneStubFlags, createZoneStub)
}
