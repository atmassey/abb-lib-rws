package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
)

func main() {
	client := abb.NewClient("localhost", "Default User", "robotics")
	msg, err := client.SubscribeToOperationMode()
	if err != nil {
		panic(err)
	}
	for range msg {
		message, ok := <-msg
		if !ok {
			fmt.Print("Channel closed")
			break
		}
		fmt.Printf("Operation Mode: %v \n", message["mode"])
	}
}
