package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/network"
	"net/http"
)

type networksListOptions struct {
	all bool
}

var (
	netsListOptions networksListOptions
)

func networksList(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var fieldArray []string
	fieldArray = append(fieldArray, "network", "network_view", "ipv4addr")
	getAllNetworksAPI := network.NewGetAllNetworks(fieldArray)

	err := client.Do(getAllNetworksAPI)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		if getAllNetworksAPI.StatusCode() == http.StatusOK {
			headers := []string{"Ref", "Network", "View", "IP"}
			rows := []map[string]interface{}{}
			for _, obj := range getAllNetworksAPI.GetResponse().([]network.Network) {
				row := map[string]interface{}{}
				row["Ref"] = obj.Ref
				row["Network"] = obj.Network
				row["View"] = obj.NetworkView
				row["IP"] = obj.Ipv4addr
				rows = append(rows, row)
			}
			PrettyPrintMany(headers, rows)
		} else {
			fmt.Printf("Response:\n%+v\n ", getAllNetworksAPI.GetResponse().(api.RespError))
		}
	}
}

func init() {
	listFlags := flag.NewFlagSet("networks-show-all", flag.ExitOnError)
	listFlags.BoolVar(&netsListOptions.all, "all", true, "List all networks")
	RegisterCliCommand("networks-show-all", listFlags, networksList)
}
