package main

import "github.com/atmassey/abb-lib-rws"

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	//restart the controller | similar to a warm start
	err := client.RestartController("restart")
	if err != nil {
		panic(err)
	}
}
