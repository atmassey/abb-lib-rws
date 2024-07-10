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

### Create a backup on the Robot Controller

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

### List actions that can be performed on the controller

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




