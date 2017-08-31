package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"net/http"
	"os"
)

func updateNSGroupFwdStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupFwdStubObject nsgroupfwdstub.NSGroupFwdStub
	var externalServerObject common.ExternalServer

	reference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	externalServerName := flagSet.Lookup("ext-server-name").Value.String()
	externalServerIP := flagSet.Lookup("ext-server-ip").Value.String()

	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}
	nsGroupFwdStubObject.Reference = reference

	if name != "" {
		nsGroupFwdStubObject.Name = name
	}
	if comment != "" {
		nsGroupFwdStubObject.Comment = comment
	}
	if externalServerName != "" || externalServerIP != "" {
		if externalServerName == "" || externalServerIP == "" {
			fmt.Printf("\nError -ext-server-name and -ext-server-ip must be changed at the same time\n")
			os.Exit(1)
		}
		externalServerObject.Name = externalServerName
		externalServerObject.Address = externalServerIP
		externalServerList := make([]common.ExternalServer, 0)
		externalServerList = append(externalServerList, externalServerObject)
		nsGroupFwdStubObject.ExternalServers = externalServerList
	}

	updateNSGroupFwdStubAPI := nsgroupfwdstub.NewUpdate(nsGroupFwdStubObject, nsgroupfwdstub.RequestReturnFields)
	err := client.Do(updateNSGroupFwdStubAPI)
	httpStatus := updateNSGroupFwdStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating NS Group Forward/Stub %s. HTTP status: %d. Error: %+v\n", nsGroupFwdStubObject.Name, httpStatus, string(updateNSGroupFwdStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated NS Group Forward/Stub %s\n", nsGroupFwdStubObject.Name)
}

func init() {
	updateNSGroupFwdStubFlags := flag.NewFlagSet("nsgroup-fwd-stub-update", flag.ExitOnError)
	updateNSGroupFwdStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REFERENCE")
	updateNSGroupFwdStubFlags.String("name", "", "usage: -name nsgroup-fwd-stub-name")
	updateNSGroupFwdStubFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupFwdStubFlags.String("ext-server-name", "", "usage: -ext-server-name FQDN")
	updateNSGroupFwdStubFlags.String("ext-server-ip", "", "usage: -ext-server-ip xxx.xxx.xxx.xxx")
	RegisterCliCommand("nsgroup-fwd-stub-update", updateNSGroupFwdStubFlags, updateNSGroupFwdStub)
}
