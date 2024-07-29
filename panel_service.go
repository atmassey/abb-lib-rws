package abb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// RestartController will restart the controller with any of the following actions: restart | istart | pstart | bstart
// CAUTION: A restart will restart the controller and all running programs will be stopped. (Warmstart)
// CAUTION: A "istart" will restart the controller and factory reset the controller.
// CAUTION: A "pstart" will restart the controller and delete all rapid programs but keep all configuration data.
// CAUTION: A "bstart" will restart the controller and revert it to its last auto-saved state.
func (c *Client) RestartController(Action string) error {
	PossibleActions := []string{"restart", "istart", "pstart", "bstart"}
	if !stringInSlice(Action, PossibleActions) {
		return fmt.Errorf("invalid action: %s", Action)
	}
	body := url.Values{}
	body.Add("restart-mode", Action)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "restart")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = q.Encode()
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

// GetOperationMode returns the current operation mode of the controller.
// Possible values: (INIT | AUTO_CH | MANF_CH | MANR | MANF | AUTO | UNDEF)
func (c *Client) GetOperationMode() (string, error) {
	var opmode OperationMode
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/panel/opmode", nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("json", "1")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&opmode)
	if err != nil {
		return "", err
	}
	mode := opmode.Embedded.State[0].Opmode
	if mode == "" {
		return "", fmt.Errorf("OP Mode Not Found: %v", opmode)
	}
	return mode, nil
}
