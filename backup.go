package abb

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// CreateBackup creates a backup of the controller to a specified directory.
// The backup path show include the environment variable along with the directory.
// Example: /$TEMP/my_backup_directory
func (c *Client) CreateBackup(Dir string) error {
	body := url.Values{}
	body.Add("backup", "/fileservice/"+Dir+"/")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/backup", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "backup")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// RestoreBackup restores a backup from a specified directory.
// The backup path show include the environment variable along with the directory.
// This does require a UAS grant.
// Example: /$SYSPAR/my_backup_directory
func (c *Client) RestoreBackup(Dir string) error {
	body := url.Values{}
	body.Add("backup", "/fileservice/"+Dir)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/backup", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "restore")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}
