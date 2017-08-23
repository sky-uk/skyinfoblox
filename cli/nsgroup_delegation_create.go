package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
	"net/http"
	"os"
)

func createNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupDelegationObject nsgroupdelegation.NSGroupDelegation
	var delegateToExternalServerObject common.ExternalServer
	var delegateToExternalServerList []common.ExternalServer

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	delegateToName := flagSet.Lookup("delegate-to-name").Value.String()
	delegateToIP := flagSet.Lookup("delegate-to-ip").Value.String()

	if name == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	nsGroupDelegationObject.Name = name

	if comment != "" {
		nsGroupDelegationObject.Comment = comment
	}

	if delegateToName == "" {
		fmt.Printf("\nError delegate-to-name argument required\n")
		os.Exit(1)
	}
	delegateToExternalServerObject.Name = delegateToName

	if delegateToIP == "" {
		fmt.Printf("\nError delegate-to-ip argument required\n")
		os.Exit(1)
	}
	delegateToExternalServerObject.Address = delegateToIP

	delegateToExternalServerList = append(delegateToExternalServerList, delegateToExternalServerObject)
	nsGroupDelegationObject.DelegateTo = delegateToExternalServerList

	createNSGroupDelegationAPI := nsgroupdelegation.NewCreate(nsGroupDelegationObject)
	err := client.Do(createNSGroupDelegationAPI)
	httpStatus := createNSGroupDelegationAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating NS Group Delegation %s. HTTP status: %d. Error: %+v\n", nsGroupDelegationObject.Name, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS Group Delegation %s\n", nsGroupDelegationObject.Name)
}

func init() {
	createNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-create", flag.ExitOnError)
	createNSGroupDelegationFlags.String("name", "", "usage: -name nsgroup-delegation-name")
	createNSGroupDelegationFlags.String("comment", "", "usage: -comment 'A Comment'")
	createNSGroupDelegationFlags.String("delegate-to-name", "", "usage: -delegate-to-name remote-host-name (CLI only supports one server)")
	createNSGroupDelegationFlags.String("delegate-to-ip", "", "usage: -delegate-to-ip remote-host-ip (CLI only supports one server)")
	RegisterCliCommand("nsgroup-delegation-create", createNSGroupDelegationFlags, createNSGroupDelegation)
}
