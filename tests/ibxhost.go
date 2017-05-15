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

	color.Blue("\tHOST RECORD")
	fmt.Println("Create_host_record")
	if ret1, err := infobx.Create_host_record(hostname, "10.77.58.21", 160); err != nil {
		color.Red("Failed to host rec:\n %s", err)
	} else {
		color.Green("host created:\n %s", ret1)
	}

	fmt.Println("Update_host_record")
	err = infobx.Update_host_record(hostname, "10.77.58.32")
	if err != nil {
		color.Red("Failed to update:\n%s", err)
	} else {
		color.Green("Update OK")
	}

	fmt.Println("Get_host_record")
	ret2, err3 := infobx.Get_host_record(hostname)
	if err3 != nil {
		color.Red("Failed to Get host\n %s", err3)
	} else {
		color.Green("Host created:\n %s", ret2)
	}

	fmt.Println("Delete_host_record")
	err = infobx.Delete_host_record(hostname)
	if err != nil {
		color.Red("Failed to delete:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("Create_host_record - no ttl")
	if ret1, err := infobx.Create_host_record(hostname, "10.77.58.21", -6); err != nil {
		color.Red("Failed to host rec:\n %s", err)
	} else {
		color.Green("host created:\n %s", ret1)
	}
	fmt.Println("Delete_host_record")
	err = infobx.Delete_host_record(hostname)
	if err != nil {
		color.Red("Failed to delete:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("\tFAILURES")

	fmt.Println("Create_host_record")
	if ret1, err := infobx.Create_host_record(hostname, "10.77.58.333", 9); err != nil {
		color.Green("Failed to host rec:\n %s", err)
	} else {
		color.Red("host created:\n %s", ret1)
	}

	fmt.Println("Delete_host_record")
	err3 = infobx.Delete_host_record("noop.noop.ovp.bskyb.com")
	if err3 != nil {
		color.Green("Failed to delete host\n %s", err3)
	} else {
		color.Red("Host created:\n %s", ret2)
	}

	fmt.Println("Get_host_record")
	ret2, err3 = infobx.Get_host_record("noop.noop.ovp.bskyb.com")
	if err3 != nil {
		color.Red("Failed to Get host\n %s", err3)
	} else {
		color.Green("Got host:\n %s", ret2)
	}

	fmt.Println("DONE\n\n")
}
