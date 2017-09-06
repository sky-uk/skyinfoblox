package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
	"sort"
)

// ExecFunc executes the function for cli.
type ExecFunc func(client *skyinfoblox.InfobloxClient, flagSet *flag.FlagSet)

// Command struct - defines a cli command with flags and exec
type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

var (
	/*
	 * InfoBlox API server
	 */
	ibxServer   string
	ibxPort     int
	debug       bool
	wapiVersion string

	/*
	 * Authentication
	 */
	ibxUsername string
	ibxPassword string

	commandMap = make(map[string]Command, 0)
)

// RegisterCliCommand - allows additional cli commands to be registered.
func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = Command{flagSet, exec}
}

// InitFlags - initiall cli flags.
func InitFlags() {
	flag.StringVar(&ibxServer, "server", os.Getenv("IBX_SERVER"),
		"Infoblox API server hostname or address. (Env: IBX_SERVER)")
	flag.StringVar(&wapiVersion, "wapiVersion", "v2.6.1",
		"WAPI version (defaults to v2.6.1)")
	flag.IntVar(&ibxPort, "port", 443,
		"Infoblox API server port. Default:443")
	flag.StringVar(&ibxUsername, "username", os.Getenv("IBX_USERNAME"),
		"Authentication username (Env: IBX_USERNAME)")
	flag.StringVar(&ibxPassword, "password", os.Getenv("IBX_PASSWORD"),
		"Authentication password (Env: IBX_PASSWORD)")
	flag.BoolVar(&debug, "debug", false, "Debug output. Default:false")
}

func usage() {
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  Commands:\n")
	var keys []string
	for name := range commandMap {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(os.Stderr, "    %s\n", key)
	}
}

func main() {
	InitFlags()
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(2)
	}

	command := flag.Arg(0)
	cmd, inMap := commandMap[command]
	if !inMap {
		usage()
		os.Exit(2)
	}

	flagSet := cmd.flagSet
	flagSet.Parse(flag.Args()[1:])

	client := skyinfoblox.NewInfobloxClient(ibxServer, ibxUsername, ibxPassword, true, debug)

	cmd.exec(client, flagSet)
}
