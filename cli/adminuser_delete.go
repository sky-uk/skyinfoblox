package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
)

func deleteAdminUser(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()

	updateUserAPI := adminuser.NewDeleteAdminUser(ref)
	err := client.Do(updateUserAPI)
	if err != nil {
		fmt.Println("Could not create the user ")
	}
	fmt.Println("Status Code: ", updateUserAPI.StatusCode())
	fmt.Printf("Response : %s", *updateUserAPI.ResponseObject().(*string))
}

func init() {
	createFlags := flag.NewFlagSet("adminuser-create", flag.ExitOnError)
	createFlags.String("ref", "", "user's unique ref")
	RegisterCliCommand("adminuser-delete", createFlags, deleteAdminUser)
}
