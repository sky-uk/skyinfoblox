package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
)

func getAdminUser(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	fieldList := []string{"name", "email", "comment", "admin_groups"}
	getUserAPI := adminuser.NewGetAdminUser(ref, fieldList)
	getUserErr := client.Do(getUserAPI)
	if getUserErr != nil {
		fmt.Println("Could not get the user")
	}
	fmt.Println("Status Code: ", getUserAPI.StatusCode())
	fmt.Printf("Response : %s", *getUserAPI.ResponseObject().(*adminuser.AdminUser))

}

func init() {
	createFlags := flag.NewFlagSet("adminuser-show", flag.ExitOnError)
	createFlags.String("ref", "", "user's unique ref")
	RegisterCliCommand("adminuser-show", createFlags, getAdminUser)

}
