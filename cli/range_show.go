package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"net/http"
	"strings"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
)

// GetDHCPRange gets a single DHCP Range
func GetDHCPRange(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	objRef := flagSet.Lookup("ref").Value.String()
	fields := flagSet.Lookup("fields").Value.String()
	var fieldArray []string
	fieldArray = strings.Split(fields, ",")
	fieldArray = append(fieldArray, "end_addr", "start_addr", "network", "network_view", "member")
	getDHCPRangeAPI := dhcprange.NewGetDHCPRangeAPI(objRef, fieldArray)

	err := client.Do(getDHCPRangeAPI)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		if getDHCPRangeAPI.StatusCode() == http.StatusOK {
			object := getDHCPRangeAPI.GetResponse()
			row := map[string]interface{}{}
			row["Network"] = object.Network
			row["Network View"] = object.NetworkView
			row["Start"] = object.Start
			row["End"] = object.End
			row["Member Address"] = object.Member.Address
			row["Member Name"] = object.Member.Name
			PrettyPrintSingle(row)
		} else {
			fmt.Println("Status code: ", getDHCPRangeAPI.StatusCode())
			fmt.Printf("Response:\n%s\n ", getDHCPRangeAPI.ResponseObject())
		}
	}
}

func init() {
	showFlags := flag.NewFlagSet("range-show", flag.ExitOnError)
	showFlags.String("ref", "", "the reference of the object to get")
	showFlags.Var(&fields, "fields", "other fields you like to get back...")
	RegisterCliCommand("range-show", showFlags, GetDHCPRange)
}