package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
	"strconv"
)

func createAdminUser(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	userName := flagSet.Lookup("username").Value.String()
	email := flagSet.Lookup("email").Value.String()
	disable, _ := strconv.ParseBool(flagSet.Lookup("disable").Value.String())
	comment := flagSet.Lookup("comment").Value.String()
	password := flagSet.Lookup("password").Value.String()

	newUser := adminuser.AdminUser{
		Name:     userName,
		Email:    email,
		Disable:  &disable,
		Comment:  comment,
		Password: password,
		Groups:   []string{"test"},
	}
	createUserAPI := adminuser.NewCreateAdminUser(newUser)
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
	RegisterCliCommand("adminuser-create", createFlags, createAdminUser)
}
