package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
)

func main() {
	client := abb.NewClient("localhost", "Default User", "robotics")
	signal_path := "LOCAL/PANEL/MAN1"
	signal, err := client.SubscribeToIOSignal(signal_path)
	if err != nil {
		panic(err)
	}
	for range signal {
		message, ok := <-signal
		if !ok {
			break
		}
		// Print the signal value
		fmt.Printf("Signal Value: %v \n", message["value"])
		fmt.Printf("Signal Simulation State: %v \n", message["state"])
	}

}
