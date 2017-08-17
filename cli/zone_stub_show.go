package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
)

func showZoneStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	returnFields := []string{"stub_from", "fqdn", "comment"}
	getAPI := zonestub.NewGet(flagSet.Lookup("ref").Value.String(), returnFields)
	getErr := client.Do(getAPI)
	if getErr != nil {
		fmt.Println(getErr.Error())
	}

	if getAPI.StatusCode() != 200 {
		fmt.Println("Could not find the reference")
		fmt.Println(*getAPI.ResponseObject().(*string))
	} else {
		response := *getAPI.ResponseObject().(*zonestub.ZoneStub)
		row := map[string]interface{}{}
		row["FQDN"] = response.FQDN
		row["Comment"] = response.Comment
		row["StubFrom"] = response.StubFrom
		PrettyPrintSingle(row)
	}

}

func init() {
	showZoneStubFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	showZoneStubFlags.String("ref", "", "usage: -ref 'Comment for the zone; maximum 256 characters'")
	RegisterCliCommand("zone-stub-show", showZoneStubFlags, showZoneStub)
}
