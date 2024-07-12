package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetRobotType() (*RobotType, error) {
	var robotType RobotType
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/system/robottype", nil)
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
	robotTypeRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(robotTypeRaw, &robotType)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return &robotType, nil
}
