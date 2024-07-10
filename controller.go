package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetControllerResources() (*ControllerResources, error) {
	var ControllerResources ControllerResources
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.IP+"/ctrl", nil)
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
	err = xml.Unmarshal(resources_raw, &ControllerResources)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &ControllerResources, nil
}

func (c *Client) GetControllerActions() (*ControllerActions, error) {
	var actions ControllerActions
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.IP+"/ctrl", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("action", "show")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	actions_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(actions_raw, &actions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &actions, nil
}
