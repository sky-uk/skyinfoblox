package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
)

func updateDHCPRange(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	startAddr := flagSet.Lookup("start").Value.String()
	endAddr := flagSet.Lookup("end").Value.String()
	network := flagSet.Lookup("network").Value.String()
	networkView := flagSet.Lookup("network-view").Value.String()
	memberFQDN := flagSet.Lookup("member-fqdn").Value.String()
	memberAddress := flagSet.Lookup("member-addr").Value.String()
	ref := flagSet.Lookup("ref").Value.String()
	allowRestart := true
	dhcpRange := dhcprange.DHCPRange{
		Ref:         ref,
		Network:     network,
		NetworkView: networkView,
		Start:       startAddr,
		End:         endAddr,
		Restart:     &allowRestart,
	}

	if memberFQDN != "" {
		dhcpMember := dhcprange.Member{ElementType: "dhcpmember", Name: memberFQDN, IPv4Address: memberAddress}
		dhcpRange.Member = dhcpMember
		dhcpRange.ServerAssociation = "MEMBER"
	}
	fmt.Println("DHCPRange object:", dhcpRange)
	updateDHCPRangeAPI := dhcprange.NewUpdateDHCPRange(dhcpRange)
	err := client.Do(updateDHCPRangeAPI)
	if err != nil {
		fmt.Println("Error creating network range" + err.Error())
	}
	if updateDHCPRangeAPI.StatusCode() == 200 {
		fmt.Println("Network range successfully created")
		if client.Debug {
			response := updateDHCPRangeAPI.GetResponse()
			fmt.Printf("%s", response)
		}
	} else {
		fmt.Printf("\nError status code was %d when attempting to creating range.\n ", updateDHCPRangeAPI.StatusCode())
		fmt.Printf("Response: %s\n", updateDHCPRangeAPI.GetResponse())
	}
}

func init() {
	updateFlags := flag.NewFlagSet("range-update", flag.ExitOnError)
	updateFlags.String("ref", "", "Ref for the record we want to update")
	updateFlags.String("start", "", "start address")
	updateFlags.String("end", "", "end address")
	updateFlags.String("network", "", "Network for the IP range")
	updateFlags.String("network-view", "default", "Network view for the IP Range, default is 'default'")
	updateFlags.String("member-fqdn", "", "FQDN of the member infoblox where this range needs to be enabled.")
	updateFlags.String("member-addr", "", "IPv4 address of the member infoblox where this range needs to be enabled.")
	RegisterCliCommand("range-update", updateFlags, updateDHCPRange)
}
