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

func updateNSGroupAuth(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	var nsGroupAuthObject nsgroupauth.NSGroupAuth

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}
	nsGroupAuthObject.Reference = reference

	name := flagSet.Lookup("name").Value.String()
	comment := flagSet.Lookup("comment").Value.String()
	gridDefault, gridDefaultErr := strconv.ParseBool(flagSet.Lookup("grid-default").Value.String())
	useExternalPrimary, useExternalPrimaryErr := strconv.ParseBool(flagSet.Lookup("use-external-primary").Value.String())

	if name != "" {
		nsGroupAuthObject.Name = name
	}
	if comment != "" {
		nsGroupAuthObject.Comment = comment
	}
	if gridDefaultErr == nil {
		nsGroupAuthObject.GridDefault = &gridDefault
	}
	if useExternalPrimaryErr == nil {
		nsGroupAuthObject.UseExternalPrimary = &useExternalPrimary
	}

	updateNSGroupAuthAPI := nsgroupauth.NewUpdate(nsGroupAuthObject, nil)
	err := client.Do(updateNSGroupAuthAPI)
	httpStatus := updateNSGroupAuthAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError updating NS Group Auth reference %s. HTTP status: %d. Error %+v\n", reference, httpStatus, err)
		os.Exit(1)
	}
	fmt.Printf("\nSuccessfully update NS Group Auth reference %s\n", reference)
}

func init() {
	updateNSGroupAuthFlags := flag.NewFlagSet("nsgroup-auth-update", flag.ExitOnError)
	updateNSGroupAuthFlags.String("name", "", "usage: -name nsgroup-auth-name")
	updateNSGroupAuthFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupAuthFlags.String("grid-default", "", "usage: -grid-default (true|false)")
	updateNSGroupAuthFlags.String("use-external-primary", "", "usage: -use-external-primary (true|false)")
	updateNSGroupAuthFlags.String("ref", "", "usage: -ref OBJECT_REFERENCE")
	RegisterCliCommand("nsgroup-auth-update", updateNSGroupAuthFlags, updateNSGroupAuth)
}
