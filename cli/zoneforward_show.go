package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"net/http"
	"os"
)

var selectedZone zoneforward.ZoneForward
var noFqdnMsg = "usage: -fqdn mydomain.com"
var noRefMsg = "usage: -ref zone_forward_/XXXXXXXX:FQDN/VIEW"

func findzoneforward(fqdn string, client *skyinfoblox.InfobloxClient) string {

	var foundzone string
	api := zoneforward.NewGetAll()
	err := client.Do(api)
	if err != nil {
		fmt.Println("Error retrieving a list of all zoneforwards when searching for FQDN: " + fqdn)
	}
	if api.StatusCode() == http.StatusOK {
		zones := *api.ResponseObject().(*[]zoneforward.ZoneForward)
		for _, zone := range zones {
			if zone.Fqdn == fqdn {
				foundzone = zone.Ref
				break
			}
		}
	} else {
		fmt.Println("Read All zoneforwards return code != 200. Response: " + api.ResponseObject().(string))
	}
	return foundzone
}

func zoneforwardShow(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if selectedZone.Ref == "" && selectedZone.Fqdn == "" {
		fmt.Println(noFqdnMsg + " OR " + noRefMsg)
		os.Exit(1)
	}

	zone := selectedZone.Ref
	if zone == "" {
		zone = findzoneforward(selectedZone.Fqdn, client)
	}
	returnFields := []string{
		"address",
		"comment",
		"disable",
		"display_domain",
		"dns_fqdn",
		"forward_to",
		"forwarders_only",
		"forwarding_servers",
		"fqdn",
		"locked_by",
		"mask_prefix",
		"parent",
		"prefix",
		"using_srg_associations",
		"view",
		"zone_format",
	}
	readzoneforwardAPI := zoneforward.NewGet(zone, returnFields)
	err := client.Do(readzoneforwardAPI)
	if err != nil {
		fmt.Println("Error reading zoneforward reference " + zone + err.Error())
	}
	retZone := *readzoneforwardAPI.ResponseObject().(*zoneforward.ZoneForward)
	if readzoneforwardAPI.StatusCode() == 200 {
		row := map[string]interface{}{}
		row["Reference"] = retZone.Ref
		row["Address"] = retZone.Address
		row["Comment"] = retZone.Comment
		row["Disable"] = retZone.Disable
		row["DisplayDomain"] = retZone.DisplayDomain
		row["DNS FQDN"] = retZone.DNSFqdn
		row["FQDN"] = retZone.Fqdn
		row["Forwarding to"] = len(retZone.ForwardTo)
		for idx, z := range retZone.ForwardTo {
			row[fmt.Sprintf("Forwarded Server %d", idx)] = z.Name
		}
		row["View"] = retZone.View
		PrettyPrintSingle(row)
	} else {
		fmt.Println("Error status code != 200 when reading zoneforward reference " + zone + " Error: " + err.Error())
	}
}

func init() {
	zoneforwardShowFlags := flag.NewFlagSet("zoneforwardshow", flag.ExitOnError)
	zoneforwardShowFlags.StringVar(&selectedZone.Ref, "ref", "", noRefMsg)
	zoneforwardShowFlags.StringVar(&selectedZone.Fqdn, "fqdn", "", noFqdnMsg)
	RegisterCliCommand("zoneforward-show", zoneforwardShowFlags, zoneforwardShow)
}
