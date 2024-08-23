package main

import (
	"time"

	"github.com/atmassey/abb-lib-rws"
)

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	//create a new backup on the robot controller
	err := client.CreateBackup("$TEMP/my_test_directory")
	if err != nil {
		panic(err)
	}
	//wait for the backup to be created on the robot controller
	time.Sleep(20 * time.Second)
	//ftp download the backup to the local machine
	err = GetDirectoryTree("hd0a/TEMP/my_test_directory", "./test")
	if err != nil {
		panic(err)
	}
	//delete the backup on the robot controller
	err = client.DeleteDirectory("$TEMP/my_test_directory")
	if err != nil {
		panic(err)
	}
}
