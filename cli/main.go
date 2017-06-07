package main

import (
	"flag"
	"fmt"
	"github.com/sky-uk/skyinfoblox"
	"os"
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
	ibxServer string
	ibxPort   int
	debug     bool

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
	flag.StringVar(&ibxServer, "server", "localhost",
		"Infoblox API server hostname or address")
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
	for name := range commandMap {
		fmt.Fprintf(os.Stderr, "    %s\n", name)
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
