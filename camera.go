package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/atmassey/abb-lib-rws/structures"
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

// RestartCamera restarts the camera with the given name.
func (c *Client) RestartCamera(Name string) error {
	body := url.Values{}
	body.Add("name", Name)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/camera", bytes.NewBufferString(body.Encode()))
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
	defer closeErrorCheck(resp.Body)
	return nil
}

// FlashCameraLEDs flashes the LEDs on the camera with the given name.
func (c *Client) FlashCameraLEDs(Name string) error {
	body := url.Values{}
	body.Add("name", Name)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/ctrl/camera", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "flash-led")
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

// SetCameraName sets the name of the camera with the given index.
func (c *Client) SetCameraName(Index int, NewName string) error {
	Index_String := fmt.Sprintf("%d", Index)
	body := url.Values{}
	body.Add("index", Index_String)
	body.Add("name", NewName)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/vision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-cameraname")
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

// RefreshCameras refreshes the cameras on the robot.
func (c *Client) RefreshCameras() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/vision", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "refresh")
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

// SetCameraDCHP sets camera to dhcp mode
func (c *Client) SetCameraDHCP(CameraName string) error {
	body := url.Values{}
	body.Add("name", CameraName)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/ethernet", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-dhcp")
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

// SetCameraDNS sets the cameras DNS settings
func (c *Client) SetCameraDNS(CameraName string, Suffix string, Server string) error {
	body := url.Values{}
	body.Add("name", CameraName)
	body.Add("dns-suffix", Suffix)
	body.Add("dns-server", Server)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/ethernet", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-dns-settings")
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

// SetUserCredentials sets the user credentials for the camera
// Restart the controller after setting the credentials
func (c *Client) SetCameraUserCredentials(CameraName string, Username string, Password string) error {
	body := url.Values{}
	body.Add("index", CameraName)
	body.Add("user", Username)
	body.Add("password", Password)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/vision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-user-credential")
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

// SetCameraIP sets the IP settings for the camera
// Restart the controller after setting the IP settings
func (c *Client) SetCameraIP(CameraName string, IP string, Subnet string, Gateway string) error {
	body := url.Values{}
	body.Add("name", CameraName)
	body.Add("address", IP)
	body.Add("netmask", Subnet)
	body.Add("gateway", Gateway)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/vision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-ip-settings")
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

// GetCameraStatus gets the status of the camera with the given name.
func (c *Client) GetCameraStatus(Cameraname string) (status map[string]string, err error) {
	var status_struct structures.CameraStatusRaw
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/vision", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("name", Cameraname)
	q.Add("resource", "camera-status")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	status_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(status_raw, &status_struct)
	if err != nil {
		return nil, err
	}
	status = make(map[string]string)
	status["camera-status"] = status_struct.Body.Div.List[0].Span.Text
	return status, nil
}
