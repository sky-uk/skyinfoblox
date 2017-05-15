package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

const hostname = "h1wup.test.ovp.bskyb.com"

func main() {

	infobx, err := do_auth()
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	fmt.Println("Create_network")
	if ret1, err := infobx.Create_network("10.77.58.0/24"); err != nil {
		color.Red("Failed to netowkr:\n %s", err)
	} else {
		color.Green("network created:\n %s", ret1)
	}
	/*
		fmt.Println("Update_host_record")
		err := infobx.Update_host_record(hostname, "10.77.58.32")
		if err != nil {
			color.Red("Failed to update:\n%s", err)
		} else {
			color.Green("Update OK")
		}
	*/

	fmt.Println("Get_network")
	ret2, err := infobx.Get_network("10.77.58.0/24")
	if err != nil {
		color.Red("Failed to network\n %s", err)
	} else {
		color.Green("Network:\n %s", ret2)
	}

	fmt.Println("Delete_network")
	err = infobx.Delete_network("10.77.58.0/24")
	if err != nil {
		color.Red("Failed to delete:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("DONE")
}
