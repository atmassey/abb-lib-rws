package abb

import (
	"fmt"
	"io"
	"os"
)

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
