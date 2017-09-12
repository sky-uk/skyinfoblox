package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
	"sort"
)

func readAllObjects(client *skyinfoblox.Client, flagSet *flag.FlagSet) {

	objType := flagSet.Lookup("type").Value.String()

	if objType == "" {
		fmt.Printf("\nError: object type is required [Usage: -type <object type>]\n")
		os.Exit(1)
	}

	objs, err := client.ReadAll(objType)
	if err != nil {
		fmt.Printf("Error reading objects of type %s, error: %s\n", objType, err)
		os.Exit(1)
	}
	if len(objs) > 0 {
		keys := make([]string, len(objs[0]))
		i := 0
		for k := range objs[0] {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		PrettyPrintMany(keys, objs)
	} else {
		fmt.Println("No objects found!")
	}
}

func init() {
	readAllObjsFlags := flag.NewFlagSet("read-all-objects", flag.ExitOnError)
	readAllObjsFlags.String("type", "", "usage: -type <object type>")
	RegisterCliCommand("read-all-objects", readAllObjsFlags, readAllObjects)
}
