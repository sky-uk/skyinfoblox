package main

import (
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
)

func findZone(fqdn string, client *skyinfoblox.InfobloxClient) string {

	//TODO probably want to search on FQDN and View?
	var foundZoneReference string
	readAllZoneAuthAPI := zoneauth.NewGetAllZones()
	err := client.Do(readAllZoneAuthAPI)
	if err != nil {
		spew.Dump("Error retrieving a list of all zones when searching for FQDN: " + fqdn)
	}
	if readAllZoneAuthAPI.StatusCode() == 200 {
		allZoneReferences := readAllZoneAuthAPI.GetResponse()
		for _, zoneReference := range *allZoneReferences {
			if zoneReference.FQDN == fqdn {
				foundZoneReference = zoneReference.Reference
				break
			}
		}
	} else {
		spew.Dump("Read All Zones return code != 200. Response: " + readAllZoneAuthAPI.ResponseObject().(string))
	}
	return foundZoneReference
}

func zoneShow(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	zoneReference := flagSet.Lookup("ref").Value.String()
	if zoneReference == "" {
		zoneReference = findZone(flagSet.Lookup("fqdn").Value.String(), client)
	}
	readZoneAuthAPI := zoneauth.NewGetSingleZone(zoneReference)
	err := client.Do(readZoneAuthAPI)
	if err != nil {
		spew.Dump("Error reading zone reference " + zoneReference + err.Error())
	}
	readZoneResponse := readZoneAuthAPI.GetResponse()
	if readZoneAuthAPI.StatusCode() == 200 {
		spew.Dump("Successfully read zone reference " + zoneReference)
		spew.Dump("FQDN: " + readZoneResponse.FQDN)
		spew.Dump("View: " + readZoneResponse.View)
		spew.Dump("Reference: " + readZoneResponse.Reference)
		spew.Dump("Comment: " + readZoneResponse.Comment)
	} else {
		spew.Dump("Error status code != 200 when reading zone reference " + zoneReference + " Error: " + err.Error())
	}
}

func init() {
	var dnsZone zoneauth.DNSZone
	zoneShowFlags := flag.NewFlagSet("zoneShow", flag.ExitOnError)
	zoneShowFlags.StringVar(&dnsZone.Reference, "ref", "", "usage: -ref zone_auth/XXXXXXXX:FQDN/VIEW")
	zoneShowFlags.StringVar(&dnsZone.FQDN, "fqdn", "", "usage: -fqdn mydomain.com")
	RegisterCliCommand("zone-show", zoneShowFlags, zoneShow)
}
