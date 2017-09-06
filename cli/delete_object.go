package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
)

func deleteObject(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	ref := flagSet.Lookup("ref").Value.String()

	if ref == "" {
		fmt.Printf("\nError: object reference is required [Usage: -ref <object reference>]\n")
		os.Exit(1)
	}

	if debug == true {
		fmt.Println("Reference to be deleted: ", ref)
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

	ref, err := ibxClient.Delete(ref)
	if err != nil {
		fmt.Printf("Error deleting reference %s, error: %s\n", ref, err)
		os.Exit(1)
	}

	fmt.Printf("\nSuccessfully delete object reference: %s\n", ref)
}

func init() {
	deleteObjectFlags := flag.NewFlagSet("delete-object", flag.ExitOnError)
	deleteObjectFlags.String("ref", "", "usage: -ref <object reference>")
	RegisterCliCommand("delete-object", deleteObjectFlags, deleteObject)
}
