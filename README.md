# ABB Robot Web Service API Wrapper

[![Build and Test](https://github.com/atmassey/abb-lib-rws/actions/workflows/go.yml/badge.svg)](https://github.com/atmassey/abb-lib-rws/actions/workflows/go.yml)
[![golangci-lint](https://github.com/atmassey/abb-lib-rws/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/atmassey/abb-lib-rws/actions/workflows/golangci-lint.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go package for [ABB Robot Web Service API](https://developercenter.robotstudio.com/api/RWS)

![ABB Robot Web Service API](https://github.com/atmassey/abb-lib-rws/blob/main/docs/rws.png?raw=true)

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

#### Get all actions that can be performed on the controller 

```Go
package main

import (
	"fmt"

	"github.com/atmassey/abb-lib-rws"
	)

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	actions, err := client.GetControllerActions()
	if err != nil {
		panic(err)
	}
    	//list all actions that can be performed on the controller
	for _, action := range actions.Actions {
		fmt.Printf("Action: %s\n", action)
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
		fmt.Printf("Name: %s, Type: %s, Value: %s\n", name,
			 signals_struct.SignalType[i], signals_struct.SignalValue[i])
	}
}

```



