package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

func deleteNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	refObj := flagSet.Lookup("objref").Value.String()
	deleteNetworkAPI := network.NewDeleteNetwork(refObj)

	err := client.Do(deleteNetworkAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Status code: ", deleteNetworkAPI.StatusCode())
	fmt.Printf("Response: %+v\n", deleteNetworkAPI.GetResponse())
}

func init() {
	deleteFlags := flag.NewFlagSet("network-delete", flag.ExitOnError)
	deleteFlags.String("objref", "", "the network objec reference")
	RegisterCliCommand("network-delete", deleteFlags, deleteNetwork)
}
