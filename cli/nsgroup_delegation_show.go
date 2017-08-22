package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func showNSGroupDelegation(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	showNSGroupDelegationFlags := flag.NewFlagSet("nsgroup-delegation-show", flag.ExitOnError)
	showNSGroupDelegationFlags.String("ref", "", "usage: -ref OBJECT_REF")
	RegisterCliCommand("nsgroup-delegation-show", showNSGroupDelegationFlags, showNSGroupDelegation)
}
