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

func createNSGroupFwdStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupFwdStubObject nsgroupfwdstub.NSGroupFwdStub
	var externalServerObject common.ExternalServer

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	externalServerName := flagSet.Lookup("ext-server-name").Value.String()
	externalServerIP := flagSet.Lookup("ext-server-ip").Value.String()

	if name == "" {
		fmt.Printf("\nError -name argument required\n")
		os.Exit(1)
	}
	nsGroupFwdStubObject.Name = name

	if comment != "" {
		nsGroupFwdStubObject.Comment = comment
	}
	if externalServerName == "" {
		fmt.Printf("\nError -external-server-name argument required\n")
		os.Exit(1)
	}
	externalServerObject.Name = externalServerName
	if externalServerIP == "" {
		fmt.Printf("\nError -external-server-ip argument required\n")
		os.Exit(1)
	}
	externalServerObject.Address = externalServerIP

	externalServerList := make([]common.ExternalServer, 0)
	externalServerList = append(externalServerList, externalServerObject)
	nsGroupFwdStubObject.ExternalServers = externalServerList

	createNSGroupFwdStubAPI := nsgroupfwdstub.NewCreate(nsGroupFwdStubObject)
	err := client.Do(createNSGroupFwdStubAPI)
	httpStatus := createNSGroupFwdStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating NS Group Forward/Stub %s. HTTP status: %d. Error: %+v\n", nsGroupFwdStubObject.Name, httpStatus, string(createNSGroupFwdStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS Group Forward/Stub %s\n", nsGroupFwdStubObject.Name)
}

func init() {
	createNSGroupFwdStubFlags := flag.NewFlagSet("nsgroup-fwd-stub-create", flag.ExitOnError)
	createNSGroupFwdStubFlags.String("name", "", "usage: -name nsgroup-fwd-stub-name")
	createNSGroupFwdStubFlags.String("comment", "", "usage: -comment 'A Comment'")
	createNSGroupFwdStubFlags.String("ext-server-name", "", "usage: -ext-server-name FQDN")
	createNSGroupFwdStubFlags.String("ext-server-ip", "", "usage: -ext-server-ip xxx.xxx.xxx.xxx")
	RegisterCliCommand("nsgroup-fwd-stub-create", createNSGroupFwdStubFlags, createNSGroupFwdStub)
}
