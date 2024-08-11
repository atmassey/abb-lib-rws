package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
)

func closeRespWithErrorCheck(r *ftp.Response) {
	if err := r.Close(); err != nil {
		fmt.Println("Error closing connection:", err)
	}
}

func closeFileWithErrorCheck(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("Error closing file:", err)
	}
}

func closeConnWithErrorCheck(c *ftp.ServerConn) {
	if err := c.Quit(); err != nil {
		fmt.Println("Error closing connection:", err)
	}
}
