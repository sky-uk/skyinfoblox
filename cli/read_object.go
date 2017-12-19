package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
	"strings"
)

func readObject(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

	ref := flagSet.Lookup("ref").Value.String()
	returnFields := flagSet.Lookup("return-fields").Value.String()

	var fieldsList []string
	if len(returnFields) > 0 {
		fieldsList = strings.Split(returnFields, ",")
	}

	if ref == "" {
		fmt.Printf("\n[ERROR] Error: object reference is required [Usage: -ref <object reference>]\n")
		os.Exit(1)
	}

	if debug == true {
		fmt.Println("[DEBUG] Reference to be read ", ref)
	}

	obj := make(map[string]interface{})
	err := client.Read(ref, fieldsList, &obj)
	if err != nil {
		fmt.Printf("[ERROR] Error reading reference %s, error: %s\n", ref, err)
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
