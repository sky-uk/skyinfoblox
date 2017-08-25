package main

import (
	"flag"
	"github.com/sky-uk/skyinfoblox"
)

func updateNSGroupFwd(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {

}

func init() {
	updateNSGroupFwdFlags := flag.NewFlagSet("nsgroup-fwd-update", flag.ExitOnError)
	updateNSGroupFwdFlags.String("name", "", "usage: -name nsgroup-name")
	updateNSGroupFwdFlags.String("comment", "", "usage: -comment 'A Comment'")
	updateNSGroupFwdFlags.String("ref", "", "usage: -ref object-reference")
	RegisterCliCommand("nsgroup-fwd-update", updateNSGroupFwdFlags, updateNSGroupFwd)
}
