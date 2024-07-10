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

func TestBackup(t *testing.T) {

	abb := NewClient("localhost", "Default User", "robotics")
	err := abb.CreateBackup("$TEMP/backup_test")
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDirectory(t *testing.T) {

	abb := NewClient("localhost", "Default User", "robotics")
	err := abb.CreateDirectory("$TEMP", "my_test_directory")
	if err != nil {
		t.Error(err)
	}
}

func TestRestartController(t *testing.T) {

	abb := NewClient("localhost", "Default User", "robotics")
	err := abb.RestartController("restart")
	if err != nil {
		t.Error(err)
	}
}
