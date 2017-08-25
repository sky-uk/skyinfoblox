package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwd"
	"net/http"
	"os"
	"strconv"
)

func updateNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	fwdServerName := flagSet.Lookup("fwd-server-name").Value.String()
	useOverrideFwders, useOverrideFwdersErr := strconv.ParseBool(flagSet.Lookup("use-override-fwders").Value.String())
	forwardersOnly, forwardersOnlyErr := strconv.ParseBool(flagSet.Lookup("forwarders-only").Value.String())
	externalFwdToName := flagSet.Lookup("ext-fwd-to-name").Value.String()
	externalFwdToIP := flagSet.Lookup("ext-fwd-to-ip").Value.String()

	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	// Getting the object from the API to populate all fields. When specifying some fields on the cmd line we need to send others at the same time to avoid silently overwriting existing values.
	nsGroupFwdGetAPI := nsgroupfwd.NewGet(reference, nsgroupfwd.RequestReturnFields)
	getErr := client.Do(nsGroupFwdGetAPI)
	httpStatus := nsGroupFwdGetAPI.StatusCode()
	if getErr != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Forwarding reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, string(nsGroupFwdGetAPI.RawResponse()))
		os.Exit(1)
	}

	// Note cli only supports one element in the ForwardingServers array and ForwardTo array
	nsGroupFwdObject := *nsGroupFwdGetAPI.ResponseObject().(*nsgroupfwd.NSGroupFwd)
	if name != "" {
		nsGroupFwdObject.Name = name
	}
	if comment != "" {
		nsGroupFwdObject.Comment = comment
	}

	if len(nsGroupFwdObject.ForwardingServers) == 0 {
		nsGroupFwdList := make([]common.ForwardingMemberServer, 1)
		nsGroupFwdObject.ForwardingServers = nsGroupFwdList
	}
	if fwdServerName != "" {
		nsGroupFwdObject.ForwardingServers[0].Name = fwdServerName
	}
	if useOverrideFwdersErr == nil {
		nsGroupFwdObject.ForwardingServers[0].UseOverrideForwarders = &useOverrideFwders
	}
	if forwardersOnlyErr == nil {
		nsGroupFwdObject.ForwardingServers[0].ForwardersOnly = &forwardersOnly
	}

	if len(nsGroupFwdObject.ForwardingServers[0].ForwardTo) == 0 {
		nsGroupForwardToList := make([]common.ExternalServer, 1)
		nsGroupFwdObject.ForwardingServers[0].ForwardTo = nsGroupForwardToList
	}
	if externalFwdToName != "" {
		nsGroupFwdObject.ForwardingServers[0].ForwardTo[0].Name = externalFwdToName
	}
	if externalFwdToIP != "" {
		nsGroupFwdObject.ForwardingServers[0].ForwardTo[0].Address = externalFwdToIP
	}

	updateNSGroupFwdAPI := nsgroupfwd.NewUpdate(nsGroupFwdObject, nsgroupfwd.RequestReturnFields)
	err := client.Do(updateNSGroupFwdAPI)
	httpStatus = updateNSGroupFwdAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating NS Group Forwarding %s. HTTP status: %d. Error: %+v\n", nsGroupFwdObject.Name, httpStatus, string(updateNSGroupFwdAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully updated NS Group Forwarding %s\n", nsGroupFwdObject.Name)
}

func init() {
	updateNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-update", flag.ExitOnError)
	updateNSGroupFwdFlags.String("name", "", "usage: -name nsgroup-name")
	updateNSGroupFwdFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupFwdFlags.String("fwd-server-name", "", "usage: -fwd-server-name grid-member-name (only supports one)")
	updateNSGroupFwdFlags.String("use-override-fwders", "", "usage: -use-override-fwders (true|false)")
	updateNSGroupFwdFlags.String("forwarders-only", "", "usage: -forwarders-only (true|false)")
	updateNSGroupFwdFlags.String("ext-fwd-to-name", "", "usage: -ext-fwd-to-name name-server (only supports one)")
	updateNSGroupFwdFlags.String("ext-fwd-to-ip", "", "usage: -ext-fwd-to-ip xxx.xxx.xxx.xxx (only supports one)")
	updateNSGroupFwdFlags.String("ref", "", "usage: -ref object-reference")
	RegisterCliCommand("nsgroup-fwd-update", updateNSGroupFwdFlags, updateNSGroupFwd)
}
