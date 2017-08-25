package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwd"
	"net/http"
	"os"
)

func showNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError -ref argument required\n")
		os.Exit(1)
	}

	nsGroupFwdShowAPI := nsgroupfwd.NewGet(reference, nsgroupfwd.RequestReturnFields)
	err := client.Do(nsGroupFwdShowAPI)
	httpStatus := nsGroupFwdShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Forwarding reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, string(nsGroupFwdShowAPI.RawResponse()))
		os.Exit(1)
	}
	// Note CLI only supports one item in each array
	response := *nsGroupFwdShowAPI.ResponseObject().(*nsgroupfwd.NSGroupFwd)
	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Forwarding Server Name"] = response.ForwardingServers[0].Name
	row["Use Override Forwarders"] = *response.ForwardingServers[0].UseOverrideForwarders
	row["Forwarders Only"] = *response.ForwardingServers[0].ForwardersOnly

	if len(response.ForwardingServers[0].ForwardTo) > 0 {
		row["Forward To Name"] = response.ForwardingServers[0].ForwardTo[0].Name
		row["Forward To Address"] = response.ForwardingServers[0].ForwardTo[0].Address
	}
	PrettyPrintSingle(row)

}

func init() {
	showNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-show", flag.ExitOnError)
	showNSGroupFwdFlags.String("ref", "", "usage: -ref object-reference")
	RegisterCliCommand("nsgroup-fwd-show", showNSGroupFwdFlags, showNSGroupFwd)
}
