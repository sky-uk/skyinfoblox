package main

import (
	"flag"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

func createNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ipAddr := flagSet.Lookup("ip_addr").Value.String()
	cidr := flagSet.Lookup("cidr").Value.String()
	createNetworkAPI := network.NewCreateNetwork(ipAddr, cidr)

	err := client.Do(createNetworkAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if createNetworkAPI.StatusCode() == 200 {
		if client.Debug {
			spew.Dump(createNetworkAPI.ResponseObject())
		}
	} else {
		fmt.Println("Status code: ", createNetworkAPI.StatusCode())
		fmt.Println("Response: ", createNetworkAPI.ResponseObject())
	}
}

func init() {
	createFlags := flag.NewFlagSet("network-create", flag.ExitOnError)
	ipAddr := createFlags.String("ip_addr", "", "the new network address")
	cidr := createFlags.String("cidr", "", "the new network Classless Inter-Domain Routing")
	RegisterCliCommand("network-create", createFlags, createNetwork)
}
