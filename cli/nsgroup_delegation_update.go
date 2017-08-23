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

func updateNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupDelegationObject nsgroupdelegation.NSGroupDelegation
	var delegateToExternalServerObject common.ExternalServer
	var delegateToExternalServerList []common.ExternalServer
	externalServerChange := false

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	delegateToName := flagSet.Lookup("delegate-to-name").Value.String()
	delegateToIP := flagSet.Lookup("delegate-to-ip").Value.String()
	reference := flagSet.Lookup("ref").Value.String()

	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}
	nsGroupDelegationObject.Reference = reference

	if name != "" {
		nsGroupDelegationObject.Name = name
	}
	if comment != "" {
		nsGroupDelegationObject.Comment = comment
	}
	if delegateToName != "" {
		delegateToExternalServerObject.Name = delegateToName
		externalServerChange = true
	}
	if delegateToIP != "" {
		delegateToExternalServerObject.Address = delegateToIP
		externalServerChange = true
	}
	if externalServerChange {
		if delegateToExternalServerObject.Name == "" || delegateToExternalServerObject.Address == "" {
			fmt.Printf("\nError both delegate-to-name and delegate-to-ip must be set at the same time or not at all\n")
			os.Exit(1)
		}
		delegateToExternalServerList = append(delegateToExternalServerList, delegateToExternalServerObject)
		nsGroupDelegationObject.DelegateTo = delegateToExternalServerList
	}
	updateNSGroupDelegationAPI := nsgroupdelegation.NewUpdate(nsGroupDelegationObject, nsgroupdelegation.RequestReturnFields)
	err := client.Do(updateNSGroupDelegationAPI)
	httpStatus := updateNSGroupDelegationAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating NS Group Delegation %s. HTTP status: %d. Error: %+v\n", nsGroupDelegationObject.Name, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated NS Group Delegation %s\n", nsGroupDelegationObject.Name)
}

func init() {
	updateNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-update", flag.ExitOnError)
	updateNSGroupDelegationFlags.String("name", "", "usage: -name nsgroup-delegation-name")
	updateNSGroupDelegationFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupDelegationFlags.String("delegate-to-name", "", "usage: -delegate-to-name remote-host-name (CLI only supports one server)")
	updateNSGroupDelegationFlags.String("delegate-to-ip", "", "usage: -delegate-to-ip remote-host-ip (CLI only supports one server)")
	updateNSGroupDelegationFlags.String("ref", "", "usage: -ref OBJECT_REFERENCE")
	RegisterCliCommand("nsgroup-delegation-update", updateNSGroupDelegationFlags, updateNSGroupDelegation)
}
