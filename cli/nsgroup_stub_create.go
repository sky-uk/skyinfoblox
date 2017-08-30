package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"net/http"
	"os"
)

func createNSGroupStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupStubObject nsgroupstub.NSGroupStub
	var memberServerObject common.MemberServer
	var memberServerList []common.MemberServer

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	gridMemberFQDN := flagSet.Lookup("fqdn").Value.String()

	if name == "" {
		fmt.Printf("\nError -name argument required\n")
		os.Exit(1)
	}
	nsGroupStubObject.Name = name

	if comment != "" {
		nsGroupStubObject.Comment = comment
	}

	// The cli only supports one element in each array.
	if gridMemberFQDN == "" {
		fmt.Printf("\nError -fqdn argument required\n")
		os.Exit(1)
	}
	memberServerObject.Name = gridMemberFQDN
	memberServerList = append(memberServerList, memberServerObject)
	nsGroupStubObject.StubMembers = memberServerList

	createNSGroupStubAPI := nsgroupstub.NewCreate(nsGroupStubObject)
	err := client.Do(createNSGroupStubAPI)
	httpStatus := createNSGroupStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating NS Group Stub %s. HTTP status: %d. Error: %+v\n", nsGroupStubObject.Name, httpStatus, string(createNSGroupStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS Group Stub %s\n", nsGroupStubObject.Name)
}

func init() {
	createNSGroupStubFlags := flag.NewFlagSet("nsgroup-stub-create", flag.ExitOnError)
	createNSGroupStubFlags.String("name", "", "usage: -name nsgroup-stub-name")
	createNSGroupStubFlags.String("comment", "", "usage: -comment 'A Comment'")
	createNSGroupStubFlags.String("fqdn", "", "usage: -fqdn GRID-MEMBER-FQDN")
	RegisterCliCommand("nsgroup-stub-create", createNSGroupStubFlags, createNSGroupStub)
}
