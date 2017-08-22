package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func deleteNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	deleteNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-delete", flag.ExitOnError)
	deleteNSGroupDelegationFlags.String("ref", "", "usage: -ref")
	RegisterCliCommand("nsgroup-delegation-delete", deleteNSGroupDelegationFlags, deleteNSGroupDelegation)
}
