package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GetUsers gets a list of users from the controller
func (c *Client) GetUsers() (*UserResources, error) {
	var users UserResources
	c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/users", nil)
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
	users_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(users_raw, &users)
	if err != nil {
		return nil, err
	}
	defer closeErrorCheck(resp.Body)
	return &users, nil
}

// LoginAsLocalUser A client is normally logged in as a remote client. To login as local client it needs access to an enabling device.
// To successfully login as local user, the client should make the request and within 5 seconds press and release the enabling button.
// Accepted types are local or remote.
func (c *Client) LoginAsLocalUser(Type_ string) error {
	if Type_ != "local" && Type_ != "remote" {
		return fmt.Errorf("invalid type %s", Type_)
	}
	body := url.Values{}
	body.Add("type", Type_)
	c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/users", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("action", "set-locale")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}
