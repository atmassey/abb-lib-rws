package abb

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {

	abb := NewClient("localhost", "Default User", "robotics")
	actions, err := abb.GetControllerActions()
	if err != nil {
		t.Error(err)
	}
	for _, option := range actions.Body.Div.Select.Options {
		fmt.Println("Option value:", option.Value)
	}
}
