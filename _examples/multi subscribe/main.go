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
	state, err := client.SubscribeToControllerState()
	if err != nil {
		panic(err)
	}
	mode, err := client.SubscribeToOperationMode()
	if err != nil {
		panic(err)
	}
	for {
		select {
		case l, ok := <-log:
			if !ok {
				fmt.Println("log channel closed")
				return
			}
			fmt.Printf("Title: %s Description: %s\n", l["title"], l["desc"])
		case s, ok := <-state:
			if !ok {
				fmt.Println("state channel closed")
				return
			}
			fmt.Printf("Controller State: %s\n", s["state"])
		case m, ok := <-mode:
			if !ok {
				fmt.Println("mode channel closed")
				return
			}
			fmt.Printf("Operation Mode: %s\n", m["mode"])
		}
	}
}
