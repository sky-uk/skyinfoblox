package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"io/ioutil"
	"os"
)

func createAndReadObject(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

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

	retObj, err := client.CreateAndRead(objType, objAsMap)
	if err != nil {
		fmt.Printf("Error creating a %s object, error: %s\n", objType, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created object:\n%+v", retObj)
}

func init() {
	createAndReadObjectFlags := flag.NewFlagSet("create-and-read-object", flag.ExitOnError)
	createAndReadObjectFlags.String("type", "", "usage: -type <object type>")
	createAndReadObjectFlags.String("profile", "", "usage: -profile <a json-encoded file with object profile>")
	RegisterCliCommand("create-and-read-object", createAndReadObjectFlags, createAndReadObject)
}
