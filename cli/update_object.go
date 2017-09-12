package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"io/ioutil"
	"os"
)

func updateObject(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

	ref := flagSet.Lookup("ref").Value.String()
	objFile := flagSet.Lookup("profile").Value.String()

	if ref == "" {
		fmt.Printf("\nError: object reference is required [Usage: -ref <object reference>]\n")
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

	ref, err = client.Update(ref, objAsMap)
	if err != nil {
		fmt.Printf("Error updating object with reference %s, error: %s", ref, err)
		os.Exit(1)
	}

	fmt.Printf("\nSuccessfully updated object, reference: %s\n", ref)
}

func init() {
	updateObjFlags := flag.NewFlagSet("update-object", flag.ExitOnError)
	updateObjFlags.String("ref", "", "usage: -ref <reference to be updated>")
	updateObjFlags.String("profile", "", "usage: -profile <a json-encoded file with new object profile>")
	RegisterCliCommand("update-object", updateObjFlags, updateObject)
}
