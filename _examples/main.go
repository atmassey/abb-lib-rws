package main

import "github.com/atmassey/abb-lib-rws"

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	//create a new backup on the robot controller
	err := client.CreateBackup("$TEMP/my_test_directory")
	if err != nil {
		panic(err)
	}
}
