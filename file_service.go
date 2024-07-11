package abb

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Delete a directory on the controller.
// The directory should be the name of the enviorment variable plus the directory.
// Example: $TEMP/my_test_directory
func (c *Client) DeleteDirectory(Path string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("DELETE", "http://"+c.Host+"/fileservice/"+Path, nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

// Create a directory on the controller.
// The resource should be the name of the enviorment variable plus the directory.
// Example: Env = $TEMP, Dir = my_test_directory
func (c *Client) CreateDirectory(Env string, Dir string) error {
	body := url.Values{}
	body.Add("fs-newname", Dir)
	body.Add("fs-action", "create")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/fileservice/"+Env, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

// Get a file from the controller and save it with the specified filename.
// Example: Source = $TEMP/my_test_file.txt, Filename = my_test_file.txt
func (c *Client) GetFile(Source string, Filename string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/fileservice/"+Source, nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	file, err := os.Create(Filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Delete a file on the controller.
// The file should be the name of the enviorment variable plus the file.
// Example: $TEMP/my_test_file.txt
func (c *Client) DeleteFile(Path string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("DELETE", "http://"+c.Host+"/fileservice/"+Path, nil)
	if err != nil {
		return err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

// Upload a file to the controller.
// The source path should be the path to the file on the local machine.
// The destination path should be the name of the enviorment variable plus the file.
// Example: Source = /home/user/my_test_file.txt, Dest = $TEMP
func (c *Client) UploadFile(SourcePath string, DestPath string) error {
	file, err := os.Open(SourcePath)
	if err != nil {
		return err
	}
	defer file.Close()
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("PUT", "http://"+c.Host+"/fileservice/"+DestPath+"/"+file.Name(), file)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}
