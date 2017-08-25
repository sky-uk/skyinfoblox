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

func createNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupFwdObject nsgroupfwd.NSGroupFwd
	var nsGroupFwdServersObject common.ForwardingMemberServer
	var nsGroupFwdToObject common.ExternalServer

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	fwdServerName := flagSet.Lookup("fwd-server-name").Value.String()
	useOverrideFwders, useOverrideFwdersErr := strconv.ParseBool(flagSet.Lookup("use-override-fwders").Value.String())
	forwardersOnly, forwardersOnlyErr := strconv.ParseBool(flagSet.Lookup("forwarders-only").Value.String())
	externalFwdToName := flagSet.Lookup("ext-fwd-to-name").Value.String()
	externalFwdToIP := flagSet.Lookup("ext-fwd-to-ip").Value.String()

	if name == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	nsGroupFwdObject.Name = name

	if comment != "" {
		nsGroupFwdObject.Comment = comment
	}
	// Note we only support one fwdServerName in the cli
	if fwdServerName == "" {
		fmt.Printf("\nError fwd-server-name argument required\n")
		os.Exit(1)
	}
	nsGroupFwdServersObject.Name = fwdServerName

	if useOverrideFwdersErr == nil {
		nsGroupFwdServersObject.UseOverrideForwarders = &useOverrideFwders
	}
	if forwardersOnlyErr == nil {
		nsGroupFwdServersObject.ForwardersOnly = &forwardersOnly
	}
	if useOverrideFwders {
		// Note we only support one externalFwdToName and one externalFwdToIP in the cli. The name should correspond with the IP.
		if externalFwdToName == "" || externalFwdToIP == "" {
			fmt.Printf("\nError both -ext-fwd-to-name and -ext-fwd-to-ip options are required when specifying -use-override-fwders. The name and IP should be the same server.\n")
			os.Exit(1)
		}
		nsGroupFwdToObject.Name = externalFwdToName
		nsGroupFwdToObject.Address = externalFwdToIP
		nsGroupFwdToList := make([]common.ExternalServer, 0)
		nsGroupFwdToList = append(nsGroupFwdToList, nsGroupFwdToObject)
		nsGroupFwdServersObject.ForwardTo = nsGroupFwdToList
	}
	nsGroupFwdServersList := make([]common.ForwardingMemberServer, 0)
	nsGroupFwdServersList = append(nsGroupFwdServersList, nsGroupFwdServersObject)
	nsGroupFwdObject.ForwardingServers = nsGroupFwdServersList

	createNSGroupFwdAPI := nsgroupfwd.NewCreate(nsGroupFwdObject)
	err := client.Do(createNSGroupFwdAPI)
	httpStatus := createNSGroupFwdAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating NS Group Forwarding %s. HTTP status: %d. Error: %+v\n", nsGroupFwdObject.Name, httpStatus, string(createNSGroupFwdAPI.RawResponse()))
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS Group Forwarding %s\n", nsGroupFwdObject.Name)

}

func init() {
	createNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-create", flag.ExitOnError)
	createNSGroupFwdFlags.String("name", "", "usage: -name nsgroup-fwd-name")
	createNSGroupFwdFlags.String("comment", "", "usage: -comment 'A Comment'")
	createNSGroupFwdFlags.String("fwd-server-name", "", "usage: -fwd-server-name name-server (only supports one)")
	createNSGroupFwdFlags.String("use-override-fwders", "", "usage: -use-override-fwders (true|false)")
	createNSGroupFwdFlags.String("forwarders-only", "", "usage: -forwarders-only (true|false)")
	createNSGroupFwdFlags.String("ext-fwd-to-name", "", "usage: -ext-fwd-to-name name-server (only supports one)")
	createNSGroupFwdFlags.String("ext-fwd-to-ip", "", "usage: -ext-fwd-to-ip xxx.xxx.xxx.xxx (only supports one)")
	RegisterCliCommand("nsgroup-fwd-create", createNSGroupFwdFlags, createNSGroupFwd)
}
