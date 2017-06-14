package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
)

// UpdateNetwork : performs the binding logic
func UpdateNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	objRef := flagSet.Lookup("objref").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	net := network.Network{
		Ref:     objRef,
		Comment: comment,
	}

	UpdateNetAPI := network.NewUpdateNetwork(net)
	err := client.Do(UpdateNetAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Status code: ", UpdateNetAPI.StatusCode())
		fmt.Printf("Response:\n%+v\n ", UpdateNetAPI.GetResponse())
	}
}

func init() {
	updateFlags := flag.NewFlagSet("network-update", flag.ExitOnError)
	updateFlags.String("objref", "", "the reference of the object to update")
	updateFlags.String("comment", "", "the new comment")
	RegisterCliCommand("network-update", updateFlags, UpdateNetwork)
}
