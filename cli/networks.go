package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	}
	if getAllNetworksAPI.StatusCode() == 200 {
		if client.Debug {
			spew.Dump(getAllNetworksAPI.ResponseObject())
		}
	} else {
		fmt.Println("Status code: ", getAllNetworksAPI.StatusCode())
		fmt.Println("Response: ", getAllNetworksAPI.ResponseObject())
	}

	for _, obj := range *getAllNetworksAPI.GetResponse() {
		fmt.Println(obj)
	}

}

func init() {
	listFlags := flag.NewFlagSet("networks-list", flag.ExitOnError)
	listFlags.BoolVar(&netsListOptions.all, "all", true, "List all networks")
	RegisterCliCommand("networks-list", listFlags, networksList)
}
