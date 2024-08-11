package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

// recursively download directories and files from the FTP server
func downloadDir(client *ftp.ServerConn, remoteDir, localDir string) error {
	// List files and directories in the remote directory
	entries, err := client.List(remoteDir)
	if err != nil {
		return err
	}

	// Ensure the local directory exists
	err = os.MkdirAll(localDir, 0755)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		remotePath := filepath.Join(remoteDir, entry.Name)
		localPath := filepath.Join(localDir, entry.Name)

		if entry.Type == ftp.EntryTypeFolder {
			// Recursively download the subdirectory
			err := downloadDir(client, remotePath, localPath)
			if err != nil {
				return err
			}
		} else {
			// Download the file
			err := downloadFile(client, remotePath, localPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// download a file from the FTP server
func downloadFile(client *ftp.ServerConn, remoteFile, localFile string) error {
	// Retrieve the remote file
	resp, err := client.Retr(remoteFile)
	if err != nil {
		return err
	}
	defer closeRespWithErrorCheck(resp)

	// Create the local file
	outFile, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer closeFileWithErrorCheck(outFile)

	// Copy the file contents to the local file
	_, err = io.Copy(outFile, resp)
	if err != nil {
		return err
	}

	return nil
}

func GetDirectoryTree(remoteDir string, localDir string) error {
	// FTP server details
	ftpServer := "localhost:21"
	username := "Default User"
	password := "robotics"

	// Connect to the FTP server
	client, err := ftp.Dial(ftpServer, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("Error connecting to FTP server:", err)
		return err
	}
	defer closeConnWithErrorCheck(client)

	// Login to the FTP server
	err = client.Login(username, password)
	if err != nil {
		fmt.Println("Error logging in to FTP server:", err)
		return err
	}

	// Download the directory tree
	err = downloadDir(client, remoteDir, localDir)
	if err != nil {
		fmt.Println("Error downloading directory:", err)
	}

	fmt.Println("Download completed successfully")
	return nil
}
