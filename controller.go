package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	defer closeErrorCheck(resp.Body)
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
	defer closeErrorCheck(resp.Body)
	for _, option := range actions.Body.Div.Select.Options {
		actionsStruct.Actions = append(actionsStruct.Actions, option.Value)
	}
	return &actionsStruct, nil
}

// SetControllerLanguage sets the language of the controller
// language can be either "en", "zh", etc. refer to RFC 3066
func (c *Client) SetControllerLanguage(language string) error {
	body := url.Values{}
	body.Add("lang", language)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-lang")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// CompressionResource will compress or decompress a file a give path
// comp must be either "comp" for compression or "dcomp" for decompression
func (c *Client) CompressionResource(srcpath string, dstpath string, comp string) error {
	body := url.Values{}
	body.Add("srcpath", srcpath)
	body.Add("dstpath", dstpath)
	if comp != "comp" && comp != "dcomp" {
		return fmt.Errorf("invalid compression type: %s", comp)
	}
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/compress", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", comp)
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
