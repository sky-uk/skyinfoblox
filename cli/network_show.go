package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

// GetNetwork : performs the binding logic
func GetNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	objRef := flagSet.Lookup("objref").Value.String()
	getNetworkAPI := network.NewGetNetwork(objRef)

	err := client.Do(getNetworkAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Status code: ", getNetworkAPI.StatusCode())
		fmt.Printf("Response:\n%+v\n ", getNetworkAPI.ResponseObject())
	}
}

func init() {
	showFlags := flag.NewFlagSet("network-show", flag.ExitOnError)
	showFlags.String("objref", "", "the reference of the object to get")
	RegisterCliCommand("network-show", showFlags, GetNetwork)
}
