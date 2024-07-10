package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetFileResources() (*FileResources, error) {
	var FileResources FileResources
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.IP+"/fileservice", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	resources_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(resources_raw, &FileResources)
	if err != nil {
		return nil, err
	}
	return &FileResources, nil
}

// Delete a directory on the controller.
// The directory should be the name of the enviorment variable plus the directory.
// Example: $TEMP/my_test_directory
func (c *Client) DeleteDirectory(Dir string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("DELETE", "http://"+c.IP+"/fileservice/"+Dir, nil)
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
	req, err := http.NewRequest("POST", "http://"+c.IP+"/fileservice/"+Env, bytes.NewBufferString(body.Encode()))
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
