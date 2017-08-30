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

func updateNSGroupStub(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupStubObject nsgroupstub.NSGroupStub
	var memberServerObject common.MemberServer
	var memberServerList []common.MemberServer

	reference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	gridMemberFQDN := flagSet.Lookup("fqdn").Value.String()

	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}
	nsGroupStubObject.Reference = reference

	if name != "" {
		nsGroupStubObject.Name = name
	}
	if comment != "" {
		nsGroupStubObject.Comment = comment
	}
	if gridMemberFQDN != "" {
		memberServerObject.Name = gridMemberFQDN
		memberServerList = append(memberServerList, memberServerObject)
		nsGroupStubObject.StubMembers = memberServerList
	}

	updateNSGroupStubAPI := nsgroupstub.NewUpdate(nsGroupStubObject, nsgroupstub.RequestReturnFields)
	err := client.Do(updateNSGroupStubAPI)
	httpStatus := updateNSGroupStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating NS Group Stub %s. HTTP status: %d. Error: %+v\n", nsGroupStubObject.Name, httpStatus, string(updateNSGroupStubAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated NS Group Stub %s\n", nsGroupStubObject.Name)

}

func init() {
	updateNSGroupStubFlags := flag.NewFlagSet("nsgroup-stub-update", flag.ExitOnError)
	updateNSGroupStubFlags.String("ref", "", "usage: -ref INFOBLOX_OBJECT_REF")
	updateNSGroupStubFlags.String("name", "", "usage: -name nsgroup-stub-name")
	updateNSGroupStubFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupStubFlags.String("fqdn", "", "usage: -fqdn GRID-MEMBER-FQDN")
	RegisterCliCommand("nsgroup-stub-update", updateNSGroupStubFlags, updateNSGroupStub)
}
