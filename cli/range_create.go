package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
)

func createDHCPRange(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	startAddr := flagSet.Lookup("start").Value.String()
	endAddr := flagSet.Lookup("end").Value.String()
	network := flagSet.Lookup("network").Value.String()
	networkView := flagSet.Lookup("network-view").Value.String()
	memberFQDN := flagSet.Lookup("member-fqdn").Value.String()
	memberAddress := flagSet.Lookup("member-addr").Value.String()
	dhcpRange := dhcprange.DHCPRange{
		Network:     network,
		NetworkView: networkView,
		Start:       startAddr,
		End:         endAddr,
		Restart:     true,
	}

	if memberFQDN != "" {
		dhcpMember := dhcprange.Member{ElementType: "dhcpmember", Name: memberFQDN, IPv4Address: memberAddress}
		dhcpRange.Member = dhcpMember
		dhcpRange.ServerAssociation = "MEMBER"
	}
	fmt.Println("DHCPRange object:", dhcpRange)
	createDHCPRangeAPI := dhcprange.NewCreateDHCPRange(dhcpRange)
	err := client.Do(createDHCPRangeAPI)
	if err != nil {
		fmt.Println("Error creating network range" + err.Error())
	}
	if createDHCPRangeAPI.StatusCode() == 201 {
		fmt.Println("Network range successfully created")
		if client.Debug {
			response := createDHCPRangeAPI.GetResponse()
			fmt.Printf("%s", response)
		}
	} else {
		fmt.Printf("\nError status code was %d when attempting to creating range.\n ", createDHCPRangeAPI.StatusCode())
		fmt.Printf("Response: %s\n", createDHCPRangeAPI.GetResponse())
	}
}

func init() {
	createFlags := flag.NewFlagSet("range-create", flag.ExitOnError)
	createFlags.String("start", "", "start address")
	createFlags.String("end", "", "end address")
	createFlags.String("network", "", "Network for the IP range")
	createFlags.String("network-view", "default", "Network view for the IP Range, default is 'default'")
	createFlags.String("member-fqdn", "", "FQDN of the member infoblox where this range needs to be enabled.")
	createFlags.String("member-addr", "", "IPv4 address of the member infoblox where this range needs to be enabled.")
	RegisterCliCommand("range-create", createFlags, createDHCPRange)
}
