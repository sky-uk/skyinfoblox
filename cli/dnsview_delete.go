package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"os"
)

func deleteDNSView(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()

	if ref == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	deleteDNSViewAPI := dnsview.NewDelete(ref)
	err := client.Do(deleteDNSViewAPI)
	if err != nil {
		fmt.Println(fmt.Printf("could not delete dns view %s", err.Error()))
	}
	if deleteDNSViewAPI.StatusCode() == 200 {
		fmt.Println("DNS view object deleted")
	}
	fmt.Println(deleteDNSViewAPI.StatusCode())
	fmt.Println(*deleteDNSViewAPI.ResponseObject().(*string))
}

func init() {
	deleteDNSViewFlags := flag.NewFlagSet("dns-view-delete", flag.ExitOnError)
	deleteDNSViewFlags.String("ref", "", "usage: -ref 'Object reference of the DNS view.'")
	RegisterCliCommand("dns-view-delete", deleteDNSViewFlags, deleteDNSView)
}
