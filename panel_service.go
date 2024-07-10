package abb

import (
	"fmt"
	"net/http"
)

// Restart the controller with any of the following actions: restart | istart | pstart | bstart
func (c *Client) RestartController(Action string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("restart-mode", Action)
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	return nil
}
