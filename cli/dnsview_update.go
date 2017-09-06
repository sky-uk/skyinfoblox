package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"os"
)

func updateDNSView(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var dnsView dnsview.DNSView
	ref := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	returnFields := []string{"name", "comment", "is_default"}

	if ref == "" {
		fmt.Printf("\nError ref argument is required\n")
		os.Exit(1)
	}

	dnsView.Reference = ref
	dnsView.Name = name
	dnsView.Comment = comment

	updateDNSViewAPI := dnsview.NewUpdate(dnsView, returnFields)
	err := client.Do(updateDNSViewAPI)
	if err != nil {
		fmt.Println(fmt.Printf("could not update dns view %s", err.Error()))
	}
	if updateDNSViewAPI.StatusCode() == 200 {
		fmt.Println(updateDNSViewAPI.StatusCode())
		fmt.Println("DNS view object updated")
	} else {
		fmt.Println("DNS view object update failed with status code %s", updateDNSViewAPI.StatusCode())
	}
}

func init() {
	updateDNSViewFlags := flag.NewFlagSet("dns-view-update", flag.ExitOnError)
	updateDNSViewFlags.String("ref", "", "usage: -ref 'Object reference of the DNS view.'")
	updateDNSViewFlags.String("name", "", "usage: -name 'Name of the DNS view.'")
	updateDNSViewFlags.String("comment", "", "usage: -comment 'Comment for the DNS view; maximum 64 characters.'")
	RegisterCliCommand("dns-view-update", updateDNSViewFlags, updateDNSView)
}
