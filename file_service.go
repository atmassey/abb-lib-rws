package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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

// Delete a directory resource from the controller.
// The resource should be the name of the enviorment variable plus the directory.
// Example: $TEMP/my_test_directory
func (c *Client) DeleteDirectoryResource(Resource string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("DELETE", "http://"+c.IP+"/fileservice/"+Resource, nil)
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
	return nil
}
