package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"strconv"
)

func createAdminRole(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	disable, _ := strconv.ParseBool(flagSet.Lookup("disable").Value.String())

	newRole := adminrole.AdminRole{
		Name:    name,
		Comment: comment,
		Disable: &disable,
	}

	createRoleAPI := adminrole.NewCreate(newRole)

	err := client.Do(createRoleAPI)
	if err != nil {
		fmt.Println("Could not create the role ")
	}
	fmt.Println("Status Code: ", createRoleAPI.StatusCode())
	fmt.Printf("Response : %s", *createRoleAPI.ResponseObject().(*string))
}

func init() {
	createFlags := flag.NewFlagSet("admin-role-create", flag.ExitOnError)
	createFlags.String("name", "", "The name of the admin role you are creating")
	createFlags.String("comment", "", "The descriptive comment of the Admin Role object.")
	createFlags.String("disable", "", "The disable flag for the admin role you are creating")
	RegisterCliCommand("admin-role-create", createFlags, createAdminRole)
}
