package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"io/ioutil"
	"net/http"
	"os"
)

var message = "usage: -config-file <a json-encoded> file with the zone forward profile>"
var templateFile string

func createZoneForward(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	zone := new(zoneforward.ZoneForward)

	if templateFile == "" {
		fmt.Println(message)
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(templateFile)
	if err != nil {
		fmt.Printf("Error reading template file, error:\n%+v\n", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(data, zone); err != nil {
		fmt.Printf("Error decoding zone template, error:\n%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Going to create forward zone:\n%+v\n", zone)

	api := zoneforward.NewCreate(*zone)
	err = client.Do(api)
	if err != nil {
		fmt.Println("Error creating new forward zone " + zone.Fqdn + ": " + err.Error())
	}

	if api.StatusCode() == http.StatusCreated {
		fmt.Println("Zone " + zone.Fqdn + " successfully created")
		if client.Debug {
			response := api.ResponseObject().(string)
			fmt.Printf("Object Reference: %s", response)
		}
	} else {
		fmt.Printf("\nError status code was %d when attempting to creating zone %s.\n ", api.StatusCode(), zone.Fqdn)
		fmt.Printf("Response: %+v\n", *api.ResponseObject().(*string))
	}
}

func init() {
	flags := flag.NewFlagSet("zoneforwardcreate", flag.ExitOnError)
	flags.StringVar(&templateFile, "config-file", "", message)
	RegisterCliCommand("zoneforward-create", flags, createZoneForward)
}
