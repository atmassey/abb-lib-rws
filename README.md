# ABB Robot Web Service API Wrapper

[![Build and Test](https://github.com/atmassey/abb-lib-rws/actions/workflows/go.yml/badge.svg)](https://github.com/atmassey/abb-lib-rws/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/atmassey/abb-lib-rws.svg)](https://pkg.go.dev/github.com/atmassey/abb-lib-rws)
[![Go Report Card](https://goreportcard.com/badge/github.com/atmassey/abb-lib-rws)](https://goreportcard.com/report/github.com/atmassey/abb-lib-rws)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go package for [ABB Robot Web Service API](https://developercenter.robotstudio.com/api/RWS)

<img src="https://github.com/atmassey/abb-lib-rws/blob/main/docs/rws.png?raw=true" alt="Robot Web Service API Wrapper" width="350" height="350">

## ABB Robot Option Requirement

- 616-1 PC Interface

## Installation

```bash
go get github.com/atmassey/abb-lib-rws
```

## Examples
There are a few full examples in the examples directory that can be referenced.

#### Create a backup on the controller 

```Go
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
```

#### Subscribe on the Controller State Websocket

```Go
package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
	)

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	// subscribe to controller state websocket
	msg, err := client.SubscribeToControllerState()
	if err != nil {
		panic(err)
	}
	for range msg {
		message, ok := <-msg
		if !ok {
			fmt.Print("Channel closed")
			break
		}
		fmt.Printf("Controller State: %v", message["state"])
	}
}

```

#### List all the IO signals and their values on the controller

```Go
package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
)

func main() {
	client := abb.NewClient("localhost", "Default User", "robotics")
	signals, err := client.GetIOSignals()
	if err != nil {
		panic(err)
	}
	for i, name := range signals.SignalName {
		fmt.Printf("Name: %s, Type: %s, Value: %v\n", name,
			signals.SignalType[i], signals.SignalValue[i])
	}
}

```



