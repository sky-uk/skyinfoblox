package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"os"
)

var updatedPermission permission.Permission

func updatePermission(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	perm := flagSet.Lookup("permission").Value.String()
	resourceType := flagSet.Lookup("resource_type").Value.String()
	object := flagSet.Lookup("object").Value.String()
	role := flagSet.Lookup("role").Value.String()
	group := flagSet.Lookup("group").Value.String()

	if ref == "" {
		fmt.Println("\nref argument is required\n")
		os.Exit(1)
	}

	updatedPermission.Reference = ref

	if group != "" && role != "" {
		fmt.Println("\nGroup and role cannot both be set\n")
		os.Exit(1)
	}

	if resourceType != "" {
		updatedPermission.ResourceType = resourceType
	}

	if object != "" {
		updatedPermission.Object = object
	}

	if role != "" {
		updatedPermission.Role = role
		updatedPermission.Group = ""
	}

	if group != "" {
		updatedPermission.Group = group
		updatedPermission.Role = ""
	}

	if perm != "" {
		updatedPermission.Permission = perm
	}

	updatePermissionAPI := permission.NewUpdate(ref, updatedPermission)

	err := client.Do(updatePermissionAPI)
	httpStatus := updatePermissionAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst updating permission. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}

	fmt.Println("Status Code: ", updatePermissionAPI.StatusCode())

	response := updatePermissionAPI.ResponseObject().(*permission.Permission)

	row := map[string]interface{}{}
	row["Permission"] = response.Permission
	row["Resource Type"] = response.ResourceType
	row["Role"] = response.Role
	row["Group"] = response.Group
	row["Reference"] = response.Reference
	PrettyPrintSingle(row)

}

func init() {
	createFlags := flag.NewFlagSet("permission-update", flag.ExitOnError)
	createFlags.String("ref", "", "A reference to the permission object you wish to update.")
	createFlags.String("permission", "", "The type of permission. Valid values are: DENY, READ, WRITE.")
	createFlags.String("resource_type", "", "The type of resource this permission applies to. ")
	createFlags.String("object", "", "A reference to a WAPI object, which will be the object this permission applies to.")
	createFlags.String("role", "", "The name of the role this permission applies to.")
	createFlags.String("group", "", "The name of the admin group this permission applies to.")
	RegisterCliCommand("permission-update", createFlags, updatePermission)
}
