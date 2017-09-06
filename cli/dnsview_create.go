package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"os"
)

func createDNSView(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var dnsView dnsview.DNSView
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()

	if name == "" {
		fmt.Printf("\nError name argument is required\n")
		os.Exit(1)
	}

	dnsView.Name = name
	dnsView.Comment = comment

	createDNSViewAPI := dnsview.NewCreate(dnsView)
	err := client.Do(createDNSViewAPI)
	if err != nil {
		fmt.Println(fmt.Printf("could not create dns view %s", err.Error()))
	}
	if createDNSViewAPI.StatusCode() == 201 {
		fmt.Println("DNS view object created")
	}
	fmt.Println(createDNSViewAPI.StatusCode())
	fmt.Println(*createDNSViewAPI.ResponseObject().(*string))
}

func init() {
	createDNSViewFlags := flag.NewFlagSet("dns-view-create", flag.ExitOnError)
	createDNSViewFlags.String("name", "", "usage: -name 'Name of the DNS view.'")
	createDNSViewFlags.String("comment", "", "usage: -comment 'Comment for the DNS view; maximum 64 characters.'")
	RegisterCliCommand("dns-view-create", createDNSViewFlags, createDNSView)
}
