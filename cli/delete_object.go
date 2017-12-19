package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
)

func deleteObject(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

	ref := flagSet.Lookup("ref").Value.String()

	if ref == "" {
		fmt.Printf("\n[ERROR] Error: object reference is required [Usage: -ref <object reference>]\n")
		os.Exit(1)
	}

	if debug == true {
		fmt.Println("[DEBUG] Reference to be deleted: ", ref)
	}

	ref, err := client.Delete(ref)
	if err != nil {
		fmt.Printf("[ERROR] Error deleting reference %s, error: %s\n", ref, err)
		os.Exit(1)
	}

	fmt.Printf("\n[DEBUG] Successfully delete object reference: %s\n", ref)
}

func init() {
	deleteObjectFlags := flag.NewFlagSet("delete-object", flag.ExitOnError)
	deleteObjectFlags.String("ref", "", "usage: -ref <object reference>")
	RegisterCliCommand("delete-object", deleteObjectFlags, deleteObject)
}
