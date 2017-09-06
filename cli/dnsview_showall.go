package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"net/http"
	"os"
)

func showAllDNSViews(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	showAllDNSViewsAPI := dnsview.NewGetAll()
	showAllDNSViewsErr := client.Do(showAllDNSViewsAPI)

	if showAllDNSViewsErr != nil {
		fmt.Println("Could not get the dns view")
		os.Exit(1)
	}
	if showAllDNSViewsAPI.StatusCode() != http.StatusOK {
		fmt.Println(*showAllDNSViewsAPI.ResponseObject().(*string))
		os.Exit(1)
	}

	allDNSViews := showAllDNSViewsAPI.ResponseObject().(*[]dnsview.DNSView)
	rows := []map[string]interface{}{}
	headers := []string{"Name", "Reference"}

	for _, singleDNSView := range *allDNSViews {
		row := map[string]interface{}{}
		row["Name"] = singleDNSView.Name
		row["Reference"] = singleDNSView.Reference
		rows = append(rows, row)
	}
	PrettyPrintMany(headers, rows)

}

func init() {
	showAllDNSViewsFlags := flag.NewFlagSet("dns-view-show-all", flag.ExitOnError)
	RegisterCliCommand("dns-view-show-all", showAllDNSViewsFlags, showAllDNSViews)
}
