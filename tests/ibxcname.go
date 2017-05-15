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

	color.Blue("\tCNAME RECORD")
	fmt.Println("Create_cname_record")
	ret1, err2 := infobx.Create_cname_record("h1wup.test.ovp.bskyb.com", "testing.test.ovp.bskyb.com")
	if err2 != nil {
		color.Red("Failed to create cname record\n%s", err2)
	} else {
		color.Green("Cname created:\n%s", ret1)
	}

	fmt.Println("Get_cname_record")
	ret2, err := infobx.Get_cname_record("testing.test.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to Get host\n%s", err)
	} else {
		color.Green("Get cname data:\n%s", ret2)
	}

	fmt.Println("Delete_cname_record")
	err = infobx.Delete_cname_record("testing.test.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to delete cname:\n%s", err)
	} else {
		color.Green("Delete OK")
	}
	fmt.Println("\tFAILURES")
	fmt.Println("Get_cname_record")
	ret2, err = infobx.Get_cname_record("noop.noop.ovp.bskyb.com")
	if err != nil {
		color.Green("Failed to Get host\n%s", err)
	} else {
		color.Green("Get cname data:\n%s", ret2)
	}

	fmt.Println("Delete_cname_record")
	err = infobx.Delete_cname_record("noop.noop.ovp.bskyb.com")
	if err != nil {
		color.Green("Failed to delete cname:\n%s", err)
	} else {
		color.Red("Delete OK")
	}

	fmt.Println("DONE\n")
}
