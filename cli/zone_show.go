package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"net/http"
	"os"
)

var zoneShowDNSZone zoneauth.DNSZone
var zoneShowFQDNMessage = "usage: -fqdn mydomain.com"
var zoneShowReferenceMessage = "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW"

func findZone(fqdn string, client *skyinfoblox.InfobloxClient) string {

	var foundZoneReference string
	readAllZoneAuthAPI := zoneauth.NewGetAllZones()
	err := client.Do(readAllZoneAuthAPI)
	if err != nil {
		fmt.Println("Error retrieving a list of all zones when searching for FQDN: " + fqdn)
	}
	if readAllZoneAuthAPI.StatusCode() == 200 {
		allZoneReferences := readAllZoneAuthAPI.GetResponse().(zoneauth.DNSZoneReferences)
		for _, zoneReference := range allZoneReferences {
			if zoneReference.FQDN == fqdn {
				foundZoneReference = zoneReference.Reference
				break
			}
		}
	} else {
		fmt.Printf("Read All Zones return code != 200. Response:\n%v\n", readAllZoneAuthAPI.GetResponse().(api.RespError))
	}
	return foundZoneReference
}

func zoneShow(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	if zoneShowDNSZone.Reference == "" && zoneShowDNSZone.FQDN == "" {
		fmt.Println(zoneShowFQDNMessage + " OR " + zoneShowReferenceMessage)
		os.Exit(1)
	}

	zoneReference := zoneShowDNSZone.Reference
	if zoneReference == "" {
		zoneReference = findZone(zoneShowDNSZone.FQDN, client)
	}
	if zoneReference == "" {
		return
	}
	returnFields := []string{"comment", "fqdn", "soa_default_ttl", "view"}
	readZoneAuthAPI := zoneauth.NewGetSingleZone(zoneReference, returnFields)
	err := client.Do(readZoneAuthAPI)
	if err != nil {
		fmt.Println("Error reading zone reference " + zoneReference + err.Error())
	}

	if readZoneAuthAPI.StatusCode() == http.StatusOK {
		readZoneResponse := readZoneAuthAPI.GetResponse().(zoneauth.DNSZone)
		row := map[string]interface{}{}
		row["FQDN"] = readZoneResponse.FQDN
		row["View"] = readZoneResponse.View
		row["Comment"] = readZoneResponse.Comment
		row["Reference"] = readZoneResponse.Reference
		row["SOA Default TTL"] = readZoneResponse.SOADefaultTTL
		PrettyPrintSingle(row)
	} else {
		fmt.Printf("Error status code != 200 when reading zone reference "+zoneReference+" Error:\n%v\n ",
			readZoneAuthAPI.GetResponse().(api.RespError))
	}
}

func init() {
	zoneShowFlags := flag.NewFlagSet("zoneShow", flag.ExitOnError)
	zoneShowFlags.StringVar(&zoneShowDNSZone.Reference, "ref", "", zoneShowReferenceMessage)
	zoneShowFlags.StringVar(&zoneShowDNSZone.FQDN, "fqdn", "", zoneShowFQDNMessage)
	RegisterCliCommand("zone-show", zoneShowFlags, zoneShow)
}
