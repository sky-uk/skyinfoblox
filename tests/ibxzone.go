package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	infobx, err := do_auth()
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	fmt.Println("Create_zone_auth")
	resp, err := infobx.Create_zone_auth("vp.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to create zone:\n%s", err)
	} else {
		color.Green("Zone created:\n%s", resp)
	}

	fmt.Println("Get_zone_auth")
	resp1, err := infobx.Get_zone_auth("vp.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to get zone:\n%s", err)
	} else {
		color.Green("Get zone auth:\n%s", resp1)
	}

	fmt.Println("Delete_zone_auth")
	err = infobx.Delete_zone_auth("vp.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to delete:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("DONE")

}
