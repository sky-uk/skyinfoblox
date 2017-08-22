package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func updateNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	updateNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-update", flag.ExitOnError)
	updateNSGroupDelegationFlags.String("name", "", "usage: -name nsgroup-delegation-name")
	updateNSGroupDelegationFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupDelegationFlags.String("ref", "", "usage: -ref OBJECT_REFERENCE")
	RegisterCliCommand("nsgroup-delegation-update", updateNSGroupDelegationFlags, updateNSGroupDelegation)
}
