package abb

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// SetCameraState sets the state of the camera to either run or standby.
// Set state to true for run and false for standby.
func (c *Client) SetCameraState(Name string, State bool) error {
	body := url.Values{}
	if State {
		body.Add("state", "run")
	} else {
		body.Add("state", "standby")
	}
	body.Add("name", Name)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/camera", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-state")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
