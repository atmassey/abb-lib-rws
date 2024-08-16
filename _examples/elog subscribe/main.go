package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
)

func main() {
	client := abb.NewClient("localhost", "Default User", "robotics")
	log, err := client.SubscribeToElog()
	if err != nil {
		panic(err)
	}
	for range log {
		message, ok := <-log
		if !ok {
			fmt.Print("Channel closed")
			break
		}
		fmt.Printf("Title: %v Description: %v", message["title"], message["desc"])
	}
}
