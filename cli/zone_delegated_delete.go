package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func deleteZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	deleteZoneAPI := zonedelegated.NewDelete(reference)
	errDelete := client.Do(deleteZoneAPI)
	if errDelete != nil {
		fmt.Println("could not delete")
	}
	if deleteZoneAPI.StatusCode() == 200 {
		fmt.Print("Record deleted")
	} else {
		fmt.Println(deleteZoneAPI.StatusCode())
	}
}

func init() {
	deleteZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-delete", flag.ExitOnError)
	deleteZoneDelegatedFlags.String("ref", "", "usage: -ref 'reference for the zone'")
	RegisterCliCommand("zone-delegated-delete", deleteZoneDelegatedFlags, deleteZoneDelegated)
}
