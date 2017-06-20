package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"strconv"
)

func createTxtRecord(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet) {
	name := flagSet.Lookup("name").Value.String()
	text := flagSet.Lookup("text").Value.String()
	view := flagSet.Lookup("view").Value.String()
	TTLStr := flagSet.Lookup("ttl").Value.String()
	useTTLStr := flagSet.Lookup("use_ttl").Value.String()
	comment := flagSet.Lookup("comment").Value.String()

	if name == "" {
		createTXTRecordUsage()
		return
	}

	if TTLStr == "" {
		TTLStr = "0"
	}
	TTL, err := strconv.ParseUint(TTLStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing TTL value: ", TTLStr)
		return
	}

	if useTTLStr == "" {
		useTTLStr = "false"
	}
	useTTL, err := strconv.ParseBool(useTTLStr)
	if err != nil {
		fmt.Println("Error parsing use_ttl value: ", useTTLStr)
		return
	}

	rec := records.TXTRecord{
		Name:    name,
		Text:    text,
		View:    view,
		TTL:     uint(TTL),
		UseTTL:  &useTTL,
		Comment: comment,
	}
	createRecordsAPI := records.NewCreateTXTRecord(rec)

	err = client.Do(createRecordsAPI)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Status Code: ", createRecordsAPI.StatusCode())
	fmt.Printf("Response : %s\n", createRecordsAPI.GetResponse())
}

func init() {
	createTXTFlags := flag.NewFlagSet("record-txt-create", flag.ExitOnError)
	createTXTFlags.String("name", "", "The name for the record in FQDN format")
	createTXTFlags.String("text", "", "Text associated with the record. It can contain up to 255 bytes per substring, up to a total of 512 bytes.")
	createTXTFlags.String("view", "", "The name of the DNS View in which the record resides.")
	createTXTFlags.String("ttl", "", "the new network Classless Inter-Domain Routing")
	createTXTFlags.String("use_ttl", "", "Associated with TTL, if false TTL is not taken into consideration")
	createTXTFlags.String("comment", "", "Comment for the record; maximum 256 characters.")
	RegisterCliCommand("record-txt-create", createTXTFlags, createTxtRecord)
}

func createTXTRecordUsage() {
	fmt.Println("Parameters:")
	fmt.Println("-name")
	fmt.Println("-text")
	fmt.Println("-view")
	fmt.Println("-ttl")
	fmt.Println("-use_ttl")
	fmt.Println("-comment")
}
