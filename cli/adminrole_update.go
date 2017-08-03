package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"net/http"
	"os"
	"strconv"
)

// UpdateAdminRole : performs the binding logic
func UpdateAdminRole(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	adminRoleReference := flagSet.Lookup("ref").Value.String()
	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	disable, _ := strconv.ParseBool(flagSet.Lookup("disable").Value.String())

	updatedAdminRole := adminrole.AdminRole{
		Name:    name,
		Comment: comment,
		Disable: &disable,
	}

	updateAdminRoleAPI := adminrole.NewUpdate(adminRoleReference, updatedAdminRole)
	err := client.Do(updateAdminRoleAPI)
	httpStatus := updateAdminRoleAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst updating admin roles. HTTP status: %d. Error: %+v\n", httpStatus, err)
		os.Exit(1)
	}
	fmt.Println("Status Code: ", updateAdminRoleAPI.StatusCode())

	response := *updateAdminRoleAPI.ResponseObject().(*adminrole.AdminRole)
	row := map[string]interface{}{}
	row["Reference"] = response.Reference
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Disabled"] = *response.Disable
	PrettyPrintSingle(row)
}

func init() {
	updateFlags := flag.NewFlagSet("admin-role-update", flag.ExitOnError)
	updateFlags.String("ref", "", "The reference of the object to update")
	updateFlags.String("name", "", "The new name of the admin role")
	updateFlags.String("comment", "", "The new comment")
	updateFlags.String("disable", "", "The disable flag.")
	RegisterCliCommand("admin-role-update", updateFlags, UpdateAdminRole)
}
