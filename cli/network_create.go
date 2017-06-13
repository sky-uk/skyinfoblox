package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

func createNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ipAddr := flagSet.Lookup("ip_addr").Value.String()
	cidr := flagSet.Lookup("cidr").Value.String()
	net := network.Network{
		Network:     ipAddr + "/" + cidr,
		NetworkView: "default",
	}
	createNetworkAPI := network.NewCreateNetwork(net)

	err := client.Do(createNetworkAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	resp := createNetworkAPI.GetResponse()
	fmt.Println("Status code: ", createNetworkAPI.StatusCode())
	fmt.Println("Response string: ", resp)
}

func init() {
	createFlags := flag.NewFlagSet("network-create", flag.ExitOnError)
	createFlags.String("ip_addr", "", "the new network address")
	createFlags.String("cidr", "", "the new network Classless Inter-Domain Routing")
	RegisterCliCommand("network-create", createFlags, createNetwork)
}
