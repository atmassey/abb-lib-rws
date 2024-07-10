package abb

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// Restart the controller with any of the following actions: restart | istart | pstart | bstart
func (c *Client) RestartController(Action string) error {
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
