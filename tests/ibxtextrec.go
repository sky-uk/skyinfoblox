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

	fmt.Println("Create_text_record")
	resp, err := infobx.Create_text_record("name1.test.ovp.bskyb.com", "this is data")
	if err != nil {
		color.Red("Failed to crete text record:\n%s", err)
	} else {
		color.Green("Record created:\n%s", resp)
	}

	fmt.Println("Get_text_record")
	resp1, err := infobx.Get_text_record("name1.test.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to get text record:\n%s", err)
	} else {
		color.Green("Get text record:\n%s", resp1)
	}

	fmt.Println("Delete_text_record")
	err = infobx.Delete_text_record("name1.test.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to delete text record:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("DONE")

}
