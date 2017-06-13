package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	if deleteNetworkAPI.StatusCode() == 200 {
		if client.Debug {
			spew.Dump(deleteNetworkAPI.ResponseObject())
		}
	} else {
		fmt.Println("Status code: ", deleteNetworkAPI.StatusCode())
		fmt.Println("Response: ", deleteNetworkAPI.ResponseObject())
	}
}

func init() {
	deleteFlags := flag.NewFlagSet("network-delete", flag.ExitOnError)
	deleteFlags.String("objref", "", "the network objec reference")
	RegisterCliCommand("network-delete", deleteFlags, deleteNetwork)
}
