package main

import (
	"errors"
	"github.com/fatih/color"
	"github.com/sky-uk/skyinfoblox"
	"os"
)

func do_auth() (*skyinfoblox.InfobloxClient, error) {
	config := true
	if os.Getenv("INFOBLOX_HOST") == "" {
		color.Red("Environment variable INFOBLOX_HOST not set and is required")
		config = false
	}
	if os.Getenv("INFOBLOX_USER") == "" {
		color.Red("Environment variable INFOBLOX_USER not set and is required")
		config = false
	}
	if os.Getenv("INFOBLOX_PASSWORD") == "" {
		color.Red("Environment variable INFOBLOX_PASSWORD not set and is required")
		config = false
	}
	if config {
		infobx := skyinfoblox.NewInfobloxClient(os.Getenv("INFOBLOX_HOST"), os.Getenv("INFOBLOX_USER"), os.Getenv("INFOBLOX_PASSWORD"), "", "", "", true, false, 0)
		return infobx, nil
	} else {
		return nil, errors.New("Unable to get configuration")
	}
}
