package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
)

func createDNSView(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	var dnsView dnsview.DNSView
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()

	dnsView.Name = name
	dnsView.Comment = comment

	createDnsViewAPI := dnsview.NewCreate(dnsView)
	err := client.Do(createDnsViewAPI)
	if err != nil {
		fmt.Println(fmt.Printf("could not create dns view %s", err.Error()))
	}
	if createDnsViewAPI.StatusCode() == 201 {
		fmt.Println("DNS view object created")
	}
	fmt.Println(createDnsViewAPI.StatusCode())
	fmt.Println(*createDnsViewAPI.ResponseObject().(*string))
}

func init() {
	createDnsViewFlags := flag.NewFlagSet("dns-view-create", flag.ExitOnError)
	createDnsViewFlags.String("name", "", "usage: -name 'Name of the DNS view.'")
	createDnsViewFlags.String("comment", "", "usage: -comment 'Comment for the DNS view; maximum 64 characters.'")
	RegisterCliCommand("dns-view-create", createDnsViewFlags, createDNSView)
}
