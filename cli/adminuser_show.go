package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
	"net/http"
)

func getAdminUser(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	fieldList := []string{"name", "email", "comment", "admin_groups"}
	getUserAPI := adminuser.NewGetAdminUser(ref, fieldList)
	getUserErr := client.Do(getUserAPI)
	if getUserErr != nil {
		fmt.Println("Could not get the user")
	}
	if getUserAPI.StatusCode() != http.StatusOK {
		fmt.Println(*getUserAPI.ResponseObject().(*string))
	} else {
		userToShow := *getUserAPI.ResponseObject().(*adminuser.AdminUser)
		row := map[string]interface{}{}
		row["ref"] = userToShow.Ref
		row["username"] = userToShow.Name
		row["comment"] = userToShow.Comment
		row["admin_grups"] = userToShow.Groups
		row["email"] = userToShow.Email
		PrettyPrintSingle(row)
	}
}

func init() {
	createFlags := flag.NewFlagSet("adminuser-show", flag.ExitOnError)
	createFlags.String("ref", "", "user's unique ref")
	RegisterCliCommand("adminuser-show", createFlags, getAdminUser)

}
