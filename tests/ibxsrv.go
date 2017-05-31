package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	srv_name := "_tcp.test.ovp.bskyb.com"
	infobx, err := do_auth()
	if err != nil {
		color.Red(fmt.Sprintf("%s", err))
		os.Exit(1)
	}

	color.Blue("\tRECORD SRV")
	fmt.Println("Create_record_srv")
	resp, err := infobx.Create_srv_record(srv_name, "host1.ovp.bskyb.com", 10, 25, 8080, 660)
	if err != nil {
		color.Red("Failed to create srv record:\n%s", err)
	} else {
		color.Green("srv record created:\n%s", resp)
	}
	resp, err = infobx.Create_srv_record(srv_name, "host2.ovp.bskyb.com", 20, 25, 8080, 660)
	if err != nil {
		color.Red("Failed to create srv record:\n%s", err)
	} else {
		color.Green("srv record created:\n%s", resp)
	}
	resp, err = infobx.Create_srv_record(srv_name, "host2.ovp.bskyb.com", 27, 25, 8081, 665)
	if err != nil {
		color.Red("Failed to create srv record:\n%s", err)
	} else {
		color.Green("srv record created:\n%s", resp)
	}
	resp, err = infobx.Create_srv_record(srv_name, "host2.ovp.bskyb.com", 27, 21, 8081, 661)
	if err != nil {
		color.Red("Failed to create srv record:\n%s", err)
	} else {
		color.Green("srv record created:\n%s", resp)
	}

	fmt.Println("Get_srv_record")
	fmt.Println("Get spcific target")
	resp1, err := infobx.Get_srv_record(srv_name, "host2.ovp.bskyb.com", 27, 21, 8081)
	if err != nil {
		color.Red("Failed to get srv record:\n%s", err)
	} else {
		color.Green("Get srv record:\n%s", resp1)
	}
	d := resp1[0].(map[string]interface{})
	fmt.Println(d["port"])
	fmt.Println("get any target")
	resp1, err = infobx.Get_srv_record(srv_name, "", -1, -1, -1)
	if err != nil {
		color.Red("Failed to get srv record:\n%s", err)
	} else {
		color.Green("Get srv record:\n%s", resp1)
	}

	fmt.Println("Delete_srv_record")
	err = infobx.Delete_srv_record(srv_name, "", -1, -1, -1)
	if err != nil {
		color.Red("Failed to delete:\n%s", err)
	} else {
		color.Green("Delete OK")
	}

	fmt.Println("\tFAILURES")

	fmt.Println("Delete_srv_record")
	err = infobx.Delete_zone_auth("noop.vp.ovp.bskyb.com")
	if err != nil {
		color.Green("Failed to delete:\n%s", err)
	} else {
		color.Red("Delete OK")
	}

	fmt.Println("Get_srv_record")
	resp1, err = infobx.Get_zone_auth("noop.vp.ovp.bskyb.com")
	if err != nil {
		color.Red("Failed to get srv record:\n%s", err)
	} else {
		color.Green("Get srv record:\n%s", resp1)
	}

	fmt.Println("DONE\n\n")

}
