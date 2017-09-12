package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"io/ioutil"
	"os"
)

func createObject(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

	objType := flagSet.Lookup("type").Value.String()
	objFile := flagSet.Lookup("profile").Value.String()

	if objType == "" {
		fmt.Printf("\nError: object type is required [Usage: -type <object type>]\n")
		os.Exit(1)
	}

	if objFile == "" {
		fmt.Printf("\nError: object profile is required [Usage: -profile <a json-encoded file>]\n")
		os.Exit(1)
	}

	// getting the object profile...
	jsonText, err := ioutil.ReadFile(objFile)
	if err != nil {
		panic(err)
	}

	objAsMap := make(map[string]interface{})
	err = json.Unmarshal(jsonText, &objAsMap)
	if err != nil {
		panic(err)
	}

	if debug == true {
		fmt.Printf("Object as Map:\n%+v\n", objAsMap)
	}

	ref, err := client.Create(objType, objAsMap)
	if err != nil {
		fmt.Printf("Error creating a %s object, error: %s\n", objType, err)

		os.Exit(1)
	}

	fmt.Printf("\nSuccessfully created object, REF: %s\n", ref)
}

func init() {
	createObjectFlags := flag.NewFlagSet("create-object", flag.ExitOnError)
	createObjectFlags.String("type", "", "usage: -type <object type>")
	createObjectFlags.String("profile", "", "usage: -profile <a json-encoded file with object profile>")
	RegisterCliCommand("create-object", createObjectFlags, createObject)
}
