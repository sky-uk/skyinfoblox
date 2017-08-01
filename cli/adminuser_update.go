package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
	"strconv"
)

func updateAdminUser(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	ref := flagSet.Lookup("ref").Value.String()
	userName := flagSet.Lookup("username").Value.String()
	email := flagSet.Lookup("email").Value.String()
	disable, _ := strconv.ParseBool(flagSet.Lookup("disable").Value.String())
	comment := flagSet.Lookup("comment").Value.String()
	password := flagSet.Lookup("password").Value.String()

	updateUser := adminuser.AdminUser{
		Ref:      ref,
		Name:     userName,
		Email:    email,
		Disable:  &disable,
		Comment:  comment,
		Password: password,
		Groups:   []string{"test"},
	}
	createUserAPI := adminuser.NewUpdateAdminUser(updateUser)
	err := client.Do(createUserAPI)
	if err != nil {
		fmt.Println("Could not create the user ")
	}
	fmt.Println("Status Code: ", createUserAPI.StatusCode())
	fmt.Printf("Response : %s", *createUserAPI.ResponseObject().(*string))
}

func init() {
	createFlags := flag.NewFlagSet("adminuser-create", flag.ExitOnError)
	createFlags.String("username", "", "the name of user you are creating")
	createFlags.String("email", "", "email address")
	createFlags.String("disable", "", "disable the user")
	createFlags.String("comment", "", "a comment for the user")
	createFlags.String("password", "", "user's password")
	createFlags.String("ref", "", "user's unique ref")
	RegisterCliCommand("adminuser-create", createFlags, updateAdminUser)
}
