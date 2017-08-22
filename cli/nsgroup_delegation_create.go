package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func createNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	createNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-create", flag.ExitOnError)
	createNSGroupDelegationFlags.String("name", "", "usage: -name nsgroup-delegation-name")
	createNSGroupDelegationFlags.String("comment", "", "usage: -comment 'A Comment'")
	RegisterCliCommand("nsgroup-delegation-create", createNSGroupDelegationFlags, createNSGroupDelegation)
}
