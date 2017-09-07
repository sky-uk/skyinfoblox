package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
	"strings"
)

func readObject(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	ref := flagSet.Lookup("ref").Value.String()
	returnFields := flagSet.Lookup("return-fields").Value.String()

	var fieldsList []string
	if len(returnFields) > 0 {
		fieldsList = strings.Split(returnFields, ",")
	}

	if ref == "" {
		fmt.Printf("\nError: object reference is required [Usage: -ref <object reference>]\n")
		os.Exit(1)
	}

	if debug == true {
		fmt.Println("Reference to be read ", ref)
	}

	params := skyinfoblox.Params{
		WapiVersion: wapiVersion,
		URL:         ibxServer,
		User:        ibxUsername,
		Password:    ibxPassword,
		IgnoreSSL:   true,
		Debug:       debug,
	}

	ibxClient := skyinfoblox.Connect(params)

	obj := make(map[string]interface{})
	err := ibxClient.Read(ref, fieldsList, &obj)
	if err != nil {
		fmt.Printf("Error reading reference %s, error: %s\n", ref, err)
		os.Exit(1)
	}
	PrettyPrintSingle(obj)
}

func init() {
	readObjectFlags := flag.NewFlagSet("read-object", flag.ExitOnError)
	readObjectFlags.String("ref", "", "usage: -ref <object reference>")
	readObjectFlags.String("return-fields", "", "usage: -return-fields <a comma-separated list of fields>")
	RegisterCliCommand("read-object", readObjectFlags, readObject)
}
