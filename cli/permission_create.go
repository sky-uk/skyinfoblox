package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"os"
)

func createPermission(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	perm := flagSet.Lookup("permission").Value.String()
	resourceType := flagSet.Lookup("resource_type").Value.String()
	object := flagSet.Lookup("object").Value.String()
	role := flagSet.Lookup("role").Value.String()
	group := flagSet.Lookup("group").Value.String()

	if object == "" && resourceType == "" {
		fmt.Printf("\nAt least one of object or resource_type is required.\n")
		os.Exit(1)
	}

	if (group != "" && role != "") || (group == "" && role == "") {
		fmt.Println("\nEither group or role is required but not both")
		os.Exit(1)
	}

	if perm == "" {
		fmt.Println("\npermission is a required argument")
		os.Exit(1)
	}

	newPermission := permission.Permission{
		Permission:   perm,
		ResourceType: resourceType,
		Object:       object,
		Role:         role,
		Group:        group,
	}

	createPermissionAPI := permission.NewCreate(newPermission)

	err := client.Do(createPermissionAPI)
	httpStatus := createPermissionAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst creating permission. HTTP status: %d. Error: %+v\n", httpStatus, err)
	}

	fmt.Println("Status Code: ", createPermissionAPI.StatusCode())
	fmt.Println("Permission created")
}

func init() {
	createFlags := flag.NewFlagSet("permission-create", flag.ExitOnError)
	createFlags.String("permission", "", "The type of permission. Valid values are: DENY, READ, WRITE.")
	createFlags.String("resource_type", "", "The type of resource this permission applies to. ")
	createFlags.String("object", "", "A reference to a WAPI object, which will be the object this permission applies to.")
	createFlags.String("role", "", "The name of the role this permission applies to.")
	createFlags.String("group", "", "The name of the admin group this permission applies to.")
	RegisterCliCommand("permission-create", createFlags, createPermission)
}
