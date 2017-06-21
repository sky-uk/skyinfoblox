package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func showAllZones(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllZoneAuthAPI := zoneauth.NewGetAllZones()
	err := client.Do(showAllZoneAuthAPI)
	if err != nil {
		fmt.Println("Error retrieving a list of all zones")
	}
	if showAllZoneAuthAPI.StatusCode() == 200 {
		allZoneReferences := showAllZoneAuthAPI.GetResponse().(zoneauth.DNSZoneReferences)
		rows := []map[string]interface{}{}
		headers := []string{"FQDN", "Reference"}

		for _, zoneReference := range allZoneReferences {
			row := map[string]interface{}{}
			row["FQDN"] = zoneReference.FQDN
			row["Reference"] = zoneReference.Reference
			rows = append(rows, row)
		}
		PrettyPrintMany(headers, rows)
	} else {
		fmt.Println("Error zone-show-all return code != 200.")
	}
}

func init() {
	readAllZoneFlags := flag.NewFlagSet("zone-show-all", flag.ExitOnError)
	RegisterCliCommand("zone-show-all", readAllZoneFlags, showAllZones)
}
