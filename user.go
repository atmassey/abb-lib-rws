package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetUsers() (*UserResources, error) {
	var users UserResources
	c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.IP+"/users", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
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
	return &users, nil
}
