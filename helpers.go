package abb

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.ToLower(b) == a {
			return true
		}
	}
	return false
}

func closeErrorCheck(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Println(err)
	}
}

func closeFileCheck(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
