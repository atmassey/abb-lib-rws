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

func (c *Client) LoginAsLocalUser(Type_ string) error {
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
