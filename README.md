# ABB Robot Web Service API Wrapper

Go package for [ABB Robot Web Service API](https://developercenter.robotstudio.com/api/rwsApi/index.html)

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

import "github.com/atmassey/abb-lib-rws"

func main() {
	//create a new client
	client := abb.NewClient("localhost", "Default User", "robotics")
	actions, err := client.GetControllerActions()
	if err != nil {
		panic(err)
	}
    	//list all actions that can be performed on the controller
	for _, option := range actions.Body.Div.Select.Options {
		fmt.Println("Option value:", option.Value)
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
	fmt.Printf("Signals: %d\n", len(signals.Body.Div.UL.LIs))

	for _, signal := range signals.Body.Div.UL.LIs {
		name, sigType, lvalue := "", "", ""
		for _, span := range signal.Spans {
			switch span.Class {
			case "name":
				name = span.Content
			case "type":
				sigType = span.Content
			case "lvalue":
				lvalue = span.Content
			}
		}
		fmt.Printf("Name: %s, Type: %s, Value: %s\n", name, sigType, lvalue)
	}
}

```



