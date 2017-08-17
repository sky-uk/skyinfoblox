package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupauth"
	"net/http"
	"os"
)

func showNSGroupAuth(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

	returnFields := []string{"comment", "external_primaries", "external_secondaries", "grid_primary", "grid_secondaries", "is_grid_default", "name", "use_external_primary"}

	reference := flagSet.Lookup("ref").Value.String()
	if reference == "" {
		fmt.Printf("\nError ref argument required\n")
		os.Exit(1)
	}

	nsGroupAuthShowAPI := nsgroupauth.NewGet(reference, returnFields)
	err := client.Do(nsGroupAuthShowAPI)
	httpStatus := nsGroupAuthShowAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		fmt.Printf("\nError whilst retrieving NS Group Auth reference %s. HTTP status: %d. Error: %+v", reference, httpStatus, err)
		os.Exit(1)
	}
	response := *nsGroupAuthShowAPI.ResponseObject().(*nsgroupauth.NSGroupAuth)
	row := map[string]interface{}{}
	row["Name"] = response.Name
	row["Comment"] = response.Comment
	row["Grid Default"] = *response.GridDefault
	row["Use External Primary"] = *response.UseExternalPrimary
	PrettyPrintSingle(row)
}

func init() {
	showNSGroupAuthFlags := flag.NewFlagSet("nsgroup-auth-show", flag.ExitOnError)
	showNSGroupAuthFlags.String("ref", "", "usage: -ref OBJECT_REF")
	RegisterCliCommand("nsgroup-auth-show", showNSGroupAuthFlags, showNSGroupAuth)
}
