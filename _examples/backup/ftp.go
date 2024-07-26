package main

import (
	"os"
	"path/filepath"

	"github.com/secsy/goftp"
)

func downloadFile(client *goftp.Client, remoteFile, localFile string) error {
	// Create the local file
	outFile, err := os.Create(localFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Download the remote file
	err = client.Retrieve(remoteFile, outFile)
	if err != nil {
		return err
	}

	return nil
}

func downloadDir(client *goftp.Client, remoteDir, localDir string) error {
	// List files and directories in the remote directory
	entries, err := client.ReadDir(remoteDir)
	if err != nil {
		return err
	}

	// Ensure the local directory exists
	err = os.MkdirAll(localDir, 0755)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		remotePath := filepath.Join(remoteDir, entry.Name())
		localPath := filepath.Join(localDir, entry.Name())

		if entry.IsDir() {
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

func GetDirectoryTree(remoteDir string, localDir string) error {
	config := goftp.Config{
		User:     "Default User",
		Password: "robotics",
	}
	server := "10.40.36.102"

	client, err := goftp.DialConfig(config, server)
	if err != nil {
		return err
	}

	err = downloadDir(client, remoteDir, localDir)
	if err != nil {
		return err
	}
	return nil
}
