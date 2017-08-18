package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupauth"
	"net/http"
	"os"
	"strconv"
)

func createNSGroupAuth(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupAuthObject nsgroupauth.NSGroupAuth

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	gridDefault, gridDefaultErr := strconv.ParseBool(flagSet.Lookup("grid-default").Value.String())
	useExternalPrimary, useExternalPrimaryErr := strconv.ParseBool(flagSet.Lookup("use-external-primary").Value.String())

	if name == "" {
		fmt.Printf("\nError name argument required\n")
		os.Exit(1)
	}
	nsGroupAuthObject.Name = name

	if comment != "" {
		nsGroupAuthObject.Comment = comment
	}
	if gridDefaultErr == nil {
		nsGroupAuthObject.GridDefault = &gridDefault
	}
	if useExternalPrimaryErr == nil {
		nsGroupAuthObject.UseExternalPrimary = &useExternalPrimary
	}

	createNSGroupAuthAPI := nsgroupauth.NewCreate(nsGroupAuthObject)
	err := client.Do(createNSGroupAuthAPI)
	httpStatus := createNSGroupAuthAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError creating NS Group Auth %s. HTTP status: %d. Error: %+v\n", nsGroupAuthObject.Name, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully created NS Group Auth %s\n", nsGroupAuthObject.Name)

}

func init() {
	createNSGroupAuthFlags := flag.NewFlagSet("nsgroup-auth-create", flag.ExitOnError)
	createNSGroupAuthFlags.String("name", "", "usage: -name nsgroup-auth-name")
	createNSGroupAuthFlags.String("comment", "", "usage: -comment 'A Comment'")
	createNSGroupAuthFlags.String("grid-default", "", "usage: -grid-default (true|false)")
	createNSGroupAuthFlags.String("use-external-primary", "", "usage: -use-external-primary (true|false)")
	RegisterCliCommand("nsgroup-auth-create", createNSGroupAuthFlags, createNSGroupAuth)
}
