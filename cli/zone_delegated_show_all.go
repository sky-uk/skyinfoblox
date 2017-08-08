package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
)

func showAllZoneDelegated(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	returnFields := []string{"comment", "delegate_to", "view", "fqdn"}
	showAPI := zonedelegated.NewGetAll(returnFields)
	showErr := client.Do(showAPI)
	if showErr != nil {
		fmt.Println("could not get zone")
	}
	if showAPI.StatusCode() == 200 {
		zoneShow := *showAPI.ResponseObject().(*[]zonedelegated.ZoneDelegated)
		for i := 0; i < len(zoneShow); i++ {
			fmt.Println(fmt.Sprintf("Reference: %s", zoneShow[i].Ref))
			fmt.Println(fmt.Sprintf("Comment: %s", zoneShow[i].Comment))
			fmt.Println(fmt.Sprintf("FQDN: %s", zoneShow[i].Fqdn))
			fmt.Println(zoneShow[i].DelegateTo)
		}

	} else {
		fmt.Println(showAPI.StatusCode())
		fmt.Println(*showAPI.ResponseObject().(*string))

	}
}

func init() {
	showAllZoneDelegatedFlags := flag.NewFlagSet("zone-delegated-showall", flag.ExitOnError)
	RegisterCliCommand("zone-delegated-showall", showAllZoneDelegatedFlags, showAllZoneDelegated)
}
