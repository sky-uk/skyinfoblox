package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
)

func deleteDHCPRange(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	refObj := flagSet.Lookup("objref").Value.String()
	deleteDHCPRangeAPI := dhcprange.NewDeleteDHCPRange(refObj)

	err := client.Do(deleteDHCPRangeAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Status code: ", deleteDHCPRangeAPI.StatusCode())
	fmt.Printf("Response: %+v\n", deleteDHCPRangeAPI.GetResponse())
}

func init() {
	deleteFlags := flag.NewFlagSet("range-delete", flag.ExitOnError)
	deleteFlags.String("objref", "", "the range object reference which will be deleted.")
	RegisterCliCommand("range-delete", deleteFlags, deleteDHCPRange)
}
