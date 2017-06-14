package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

type networksListOptions struct {
	all bool
}

var (
	netsListOptions networksListOptions
)

func networksList(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	getAllNetworksAPI := network.NewGetAllNetworks()

	err := client.Do(getAllNetworksAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Status code: ", getAllNetworksAPI.StatusCode())
		fmt.Printf("Response:\n%+v\n ", getAllNetworksAPI.GetResponse())
	}
}

func init() {
	listFlags := flag.NewFlagSet("networks-show-all", flag.ExitOnError)
	listFlags.BoolVar(&netsListOptions.all, "all", true, "List all networks")
	RegisterCliCommand("networks-show-all", listFlags, networksList)
}
