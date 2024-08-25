package abb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/atmassey/abb-lib-rws/structures"
)

// GetIOSignals returns a struct of all IO signals on the robot with their names and values.
func (c *Client) GetIOSignals() (*structures.IOSignals, error) {
	var signals structures.IOSignals
	var signalsRaw structures.IOSignalsJson
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/iosystem/signals", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("json", "1")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&signalsRaw)
	if err != nil {
		return nil, err
	}
	defer closeErrorCheck(resp.Body)
	for _, signal := range signalsRaw.Embedded.State {
		signals.SignalName = append(signals.SignalName, signal.Name)
		signals.SignalType = append(signals.SignalType, signal.Type)
		signals.SignalValue = append(signals.SignalValue, signal.Value)
	}
	return &signals, nil
}

// UpdateIODevice is used to enable or disable an IO device.
// Possible values: {enable | disable}
// Possible Device path example: Local/DRV_1
func (c *Client) UpdateIODevice(State string, DevicePath string) error {
	if State != "enable" && State != "disable" {
		return fmt.Errorf("invalid state: %s", State)
	}
	body := url.Values{}
	body.Add("lstate", State)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/iosystem/devices/"+DevicePath, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set")
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
