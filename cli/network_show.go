package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/network"
	"net/http"
	"strings"
)

type arrayFlags []string

var fields arrayFlags

// GetNetwork : performs the binding logic
func GetNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	objRef := flagSet.Lookup("ref").Value.String()
	fields := flagSet.Lookup("fields").Value.String()
	var fieldArray []string
	fieldArray = strings.Split(fields, ",")
	fieldArray = append(fieldArray, "network", "network_view", "ipv4addr")
	getNetworkAPI := network.NewGetNetwork(objRef, fieldArray)

	err := client.Do(getNetworkAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		if getNetworkAPI.StatusCode() == http.StatusOK {
			object := getNetworkAPI.GetResponse().(network.Network)
			row := map[string]interface{}{}
			row["Network"] = object.Network
			row["View"] = object.NetworkView
			row["Ipv4addr"] = object.Ipv4addr
			PrettyPrintSingle(row)
		} else {
			fmt.Println("Status code: ", getNetworkAPI.StatusCode())
			fmt.Printf("Response:\n%+v\n ", getNetworkAPI.GetResponse().(api.RespError))
		}
	}
}

func init() {
	showFlags := flag.NewFlagSet("network-show", flag.ExitOnError)
	showFlags.String("ref", "", "the reference of the object to get")
	showFlags.Var(&fields, "fields", "other fields you like to get back...")
	RegisterCliCommand("network-show", showFlags, GetNetwork)
}

func (i *arrayFlags) Set(value string) error {

	//*i = append(*i, value)
	*i = strings.Split(value, ",")
	return nil
}

func (i *arrayFlags) String() string {
	return strings.Join(*i, ",")
}
