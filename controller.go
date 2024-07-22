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
	req, err := http.NewRequest("GET", "http://"+c.Host+"/ctrl", nil)
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
	resourcesRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(resourcesRaw, &ControllerResources)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &ControllerResources, nil
}

// GetControllerActions returns the actions that can be performed on the controller
func (c *Client) GetControllerActions() (*ControllerActions, error) {
	var actions ControllerActionsHTML
	var actionsStruct ControllerActions
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/ctrl", nil)
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
	actionsRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(actionsRaw, &actions)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	for _, option := range actions.Body.Div.Select.Options {
		actionsStruct.Actions = append(actionsStruct.Actions, option.Value)
	}
	return &actionsStruct, nil
}
