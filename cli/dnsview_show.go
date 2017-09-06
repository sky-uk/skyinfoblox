package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"net/http"
	"os"
)

func getDNSView(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	returnFields := []string{"name", "comment", "is_default"}

	if ref == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	getDNSViewAPI := dnsview.NewGet(ref, returnFields)
	getDNSViewErr := client.Do(getDNSViewAPI)
	if getDNSViewErr != nil {
		fmt.Println("Could not get the dns view")
	}
	if getDNSViewAPI.StatusCode() != http.StatusOK {
		fmt.Println(*getDNSViewAPI.ResponseObject().(*string))
	} else {
		dnsViewReturned := *getDNSViewAPI.ResponseObject().(*dnsview.DNSView)
		row := map[string]interface{}{}
		row["ref"] = dnsViewReturned.Reference
		row["name"] = dnsViewReturned.Name
		row["comment"] = dnsViewReturned.Comment
		row["is_default"] = *dnsViewReturned.IsDefault
		PrettyPrintSingle(row)
	}
}

func init() {
	getDNSViewFlags := flag.NewFlagSet("dns-view-show", flag.ExitOnError)
	getDNSViewFlags.String("ref", "", "usage: -ref 'Object reference of the DNS view.'")
	RegisterCliCommand("dns-view-show", getDNSViewFlags, getDNSView)
}
