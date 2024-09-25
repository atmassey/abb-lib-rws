package abb

import (
	"fmt"
	"net/http"
)

// ClearProfinetAlarms clears the alarms for a specific profinet device
func (c *Client) ClearProfinetAlarms(Device string, Network string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"rw/iosystem/devices/"+Network+"/"+Device+"/alarms/clear", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Status: %s", resp.Status)
	}
	return nil
}
