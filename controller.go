package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/atmassey/abb-lib-rws/structures"
)

// GetControllerResources returns a struct of the XML response for capturing the controllers
// resources.
func (c *Client) GetControllerResources() (*structures.ControllerResources, error) {
	var ControllerResources structures.ControllerResources
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
func (c *Client) GetControllerActions() (*structures.ControllerActions, error) {
	var actions structures.ControllerActionsHTML
	var actionsStruct structures.ControllerActions
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

// RestoreSafetyController will reset the safety controller
// Be careful with this function as it will reset the safety controller to its factory default state
func (c *Client) FactoryDefaultSafetyController() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/safety", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "reset")
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

func (c *Client) SetClock(Time structures.Clock) error {
	if _, err := strconv.Atoi(Time.Year); err != nil {
		return fmt.Errorf("invalid year format")
	}
	if _, err := strconv.Atoi(Time.Month); err != nil {
		return fmt.Errorf("invalid month format")
	}
	if _, err := strconv.Atoi(Time.Day); err != nil {
		return fmt.Errorf("invalid day format")
	}
	if _, err := strconv.Atoi(Time.Hour); err != nil {
		return fmt.Errorf("invalid hour format")
	}
	if _, err := strconv.Atoi(Time.Minute); err != nil {
		return fmt.Errorf("invalid minute format")
	}
	if _, err := strconv.Atoi(Time.Second); err != nil {
		return fmt.Errorf("invalid second format")
	}
	body := url.Values{}
	body.Add("sys-clock-year", Time.Year)
	body.Add("sys-clock-month", Time.Month)
	body.Add("sys-clock-day", Time.Day)
	body.Add("sys-clock-hour", Time.Hour)
	body.Add("sys-clock-minute", Time.Minute)
	body.Add("sys-clock-second", Time.Second)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("PUT", "http://"+c.Host+"/ctrl/clock", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// SetIdentity sets the controller name and id
func (c *Client) SetIdentity(ControllerName string, ControllerId string) error {
	body := url.Values{}
	body.Add("ctrl-name", ControllerName)
	body.Add("ctrl-id", ControllerId)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("PUT", "http://"+c.Host+"/ctrl/identity", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// SetControllerNetworkConfiguration sets the network configuration of the controller
// Method can be either "fixip", "dhcp", or "noip"
func (c *Client) SetControllerNetworkConfiguration(Method string, Address string, Mask string, Gateway string) error {
	if Method != "fixip" && Method != "dhcp" && Method != "noip" {
		return fmt.Errorf("invalid method %s", Method)
	}
	body := url.Values{}
	body.Add("method", Method)
	body.Add("address", Address)
	body.Add("mask", Mask)
	body.Add("gateway", Gateway)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/network", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("action", "set")
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

// UnlockSafetyController will unlock the safety controller if the proper user is logged in and authenticated.
func (c *Client) UnlockSafetyController() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/safety", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "unlock")
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

// Add Rotue Table Entry will add a route table entry to the controller network stack
func (c *Client) AddRouteTableEntry(destination string, gateway string) error {
	body := url.Values{}
	body.Add("destination", destination)
	body.Add("gateway", gateway)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/network/route/add", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// RemoveRoutTableEntry will remove a route table entry from the controller network stack
func (c *Client) RemoveRouteTableEntry(destination string) error {
	body := url.Values{}
	body.Add("destination", destination)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/network/route/remove", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

func (c *Client) SetBootDevice(path string) error {
	body := url.Values{}
	body.Add("path", path)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/system", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("action", "set-bootdevice")
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

func (c *Client) RemoveValidationInfo() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/safety", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "invalidate-cfg")
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

// AddValidationInfo will add validation information to the controller with user who validated the configuration
func (c *Client) AddValidationInfo(validated_by string) error {
	body := url.Values{}
	body.Add("validated-by", validated_by)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/safety", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("action", "validate-cfg")
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

// SoftwareSyncAcknowledgement will acknowledge a software safety sync request
func (c *Client) SoftwareSyncAcknowledgement(index int) error {
	index_string := strconv.Itoa(index)
	if index_string == "" || (index != 1 && index != 2) {
		return fmt.Errorf("invalid index: %d must be 1 or 2", index)
	}
	body := url.Values{}
	body.Add("index", index_string)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/safety", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "syncack")
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
