package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
)

func showAllForwardZones(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	api := zoneforward.NewGetAll()
	err := client.Do(api)
	if err != nil {
		fmt.Println("Error retrieving a list of all zones")
	}
	if api.StatusCode() == 200 {
		allZoneReferences := *api.ResponseObject().(*[]zoneforward.ZoneForward)
		rows := []map[string]interface{}{}
		headers := []string{"FQDN", "Reference"}

		for _, zone := range allZoneReferences {
			row := map[string]interface{}{}
			row["FQDN"] = zone.Fqdn
			row["Reference"] = zone.Ref
			rows = append(rows, row)
		}
		PrettyPrintMany(headers, rows)
	} else {
		fmt.Println("Error zoneforward-show-all return code != 200.")
	}
}

func init() {
	readAllForwardZoneFlags := flag.NewFlagSet("zoneforward-show-all", flag.ExitOnError)
	RegisterCliCommand("zoneforward-show-all", readAllForwardZoneFlags, showAllForwardZones)
}
