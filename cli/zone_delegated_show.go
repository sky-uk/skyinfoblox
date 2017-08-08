package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func showZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	reference := flagSet.Lookup("ref").Value.String()
	returnFields := []string{"comment", "delegate_to", "view", "fqdn"}
	showAPI := zonedelegated.NewGet(reference, returnFields)
	showErr := client.Do(showAPI)
	if showErr != nil {
		fmt.Println("could not get zone")
	}
	if showAPI.StatusCode() == 200 {
		zoneShow := showAPI.ResponseObject().(*zonedelegated.ZoneDelegated)
		fmt.Println(zoneShow.Ref)
		fmt.Println(zoneShow.Comment)
		fmt.Println(zoneShow.Fqdn)
		fmt.Println(zoneShow.DelegateTo)
	} else {
		fmt.Println(showAPI.StatusCode())
		fmt.Println(showAPI.ResponseObject().(*string))

	}
}

func init() {
	showZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-show", flag.ExitOnError)
	showZoneDelegatedFlags.String("ref", "", "usage: -ref 'reference for the zone'")
	RegisterCliCommand("zone-delegated-show", showZoneDelegatedFlags, showZoneDelegated)
}
