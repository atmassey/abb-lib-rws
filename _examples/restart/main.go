package main

import "github.com/atmassey/abb-lib-rws"

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	//warmtstart the robot controller
	err := client.Warmstart()
	if err != nil {
		panic(err)
	}
}
