package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
)

func deleteAdminRole(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()

	deleteAdminRoleAPI := adminrole.NewDelete(ref)
	err := client.Do(deleteAdminRoleAPI)
	if err != nil {
		fmt.Println("Could not create the user ")
	}
	fmt.Println("Status Code: ", deleteAdminRoleAPI.StatusCode())
	fmt.Printf("Response : %s", *deleteAdminRoleAPI.ResponseObject().(*string))
}

func init() {
	createFlags := flag.NewFlagSet("admin-role-delete", flag.ExitOnError)
	createFlags.String("ref", "", "Reference to the admin role object you wish to delete")
	RegisterCliCommand("admin-role-delete", createFlags, deleteAdminRole)
}
