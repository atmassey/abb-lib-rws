package main

import (
	"testing"
)

func TestBackup(t *testing.T) {
	//create a new client
	err := GetDirectoryTree("/hd0a/BACKUP/test", "./test")
	if err != nil {
		t.Error(err)
	}
}
