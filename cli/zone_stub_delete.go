package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
	"net/http"
)

func deleteZoneStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	deleteAPI := zonestub.NewDelete(reference)
	deleteErr := client.Do(deleteAPI)
	if deleteErr != nil {
		fmt.Println("Could not delete the record")
	}
	if deleteAPI.StatusCode() == http.StatusOK {
		fmt.Println("Zone Deleted")
	} else {
		fmt.Println(*deleteAPI.ResponseObject().(*string))
	}
}

func init() {
	deleteZoneStubFlags := flag.NewFlagSet("zone-delegated-create", flag.ExitOnError)
	deleteZoneStubFlags.String("ref", "", "usage: -ref 'Comment for the zone; maximum 256 characters'")
	RegisterCliCommand("zone-stub-delete", deleteZoneStubFlags, deleteZoneStub)
}
