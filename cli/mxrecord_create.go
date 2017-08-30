package main

import (
	"github.com/sky-uk/skyinfoblox"
	"flag"
)

func createMXRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet){

}


func init(){
	createMXRecordFlags := flag.NewFlagSet("mxrecord-create", flag.ExitOnError)

	RegisterCliCommand("mxrecord-create", createMXRecordFlags, createMXRecord)
}