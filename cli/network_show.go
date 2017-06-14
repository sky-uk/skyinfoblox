package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
	"strings"
)

type arrayFlags []string

var fields arrayFlags

// GetNetwork : performs the binding logic
func GetNetwork(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	objRef := flagSet.Lookup("objref").Value.String()
	fields := flagSet.Lookup("fields").Value.String()
	var fieldArray []string
	fieldArray = strings.Split(fields, ",")
	getNetworkAPI := network.NewGetNetwork(objRef, fieldArray)

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
