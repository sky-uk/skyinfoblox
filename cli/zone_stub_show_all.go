package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	//"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
)

func showAllZoneStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	returnFields := []string{"stub_from", "fqdn", "comment"}
	getAPI := zonestub.NewGetAll(returnFields)
	getErr := client.Do(getAPI)
	if getErr != nil {
		fmt.Println(getErr.Error())
	}

	if getAPI.StatusCode() != 200 {
		fmt.Println("Could not find the reference")
		fmt.Println(*getAPI.ResponseObject().(*string))
	} else {
		response := *getAPI.ResponseObject().(*[]zonestub.ZoneStub)
		for i := 0; i < len(response); i++ {
			fmt.Println(fmt.Sprintf("Reference: %s", response[i].Ref))
			fmt.Println(fmt.Sprintf("Comment: %s", response[i].Comment))
			fmt.Println(fmt.Sprintf("FQDN: %s", response[i].FQDN))
			fmt.Println(response[i].StubFrom)
		}
	}

}

func init() {
	showAllZoneStubFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	showAllZoneStubFlags.String("ref", "", "usage: -ref 'Comment for the zone; maximum 256 characters'")
	RegisterCliCommand("zone-stub-showall", showAllZoneStubFlags, showAllZoneStub)
}
