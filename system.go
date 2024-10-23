package abb

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/atmassey/abb-lib-rws/structures"
)

// GetRobotType returns a struct of the robot type.
func (c *Client) GetRobotType() (*structures.RobotType, error) {
	var robotType structures.RobotType
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
	defer closeErrorCheck(resp.Body)
	return &robotType, nil
}

// GetSystemEnergyMetrics returns a struct of energy metrics for each axis on the controller.
// The struct includes the axis title and the energy consumption for the axis.
// The struct also includes the total accumulated energy consumption.
func (c *Client) GetSystemEnergyMetrics() (*structures.SystemEnergyMetrics, error) {
	var EnergyMetricsDecoded structures.SystemEnergyMetrics
	var EnergyMetricsRaw structures.SystemEnergy
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/system/energy", nil)
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
	defer closeErrorCheck(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&EnergyMetricsRaw)
	if err != nil {
		return nil, err
	}
	EnergyMetricsDecoded.AccumulatedEnergy = EnergyMetricsRaw.Embedded.State[0].AccumulatedEnergy
	for _, state := range EnergyMetricsRaw.Embedded.State {
		for _, MechUnits := range state.MechUnits {
			for _, axis := range MechUnits.Axis {
				axisEnergy := structures.SystemAxisEnergy{Axis: axis.Title, Energy: axis.IntervalEnergy}
				EnergyMetricsDecoded.AxisEnergy = append(EnergyMetricsDecoded.AxisEnergy, axisEnergy)
			}
		}
	}
	return &EnergyMetricsDecoded, nil
}

// GetInstalledProducts returns a struct of installed products on the controller.
// The struct includes the product title and version.
func (c *Client) GetInstalledProducts() (*structures.InstalledSystemProducts, error) {
	var InstalledProducts structures.InstalledProducts
	var InstalledProductsDecoded structures.InstalledSystemProducts
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/system/products", nil)
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
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&InstalledProducts)
	if err != nil {
		return nil, err
	}
	installedProducts := structures.InstalledSystemProducts{}
	for _, product := range InstalledProducts.State {
		title := product.Title
		version := product.VersionName
		InstalledProductsDecoded.Title = append(installedProducts.Title, title)
		InstalledProductsDecoded.Version = append(installedProducts.Version, version)
	}
	return &InstalledProductsDecoded, nil
}

// KeylessMotorOn is used to turn on the motors without utilizing the key.
func (c *Client) KeylessMotorOn() error {
	body := url.Values{}
	body.Add("state", "run")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/cfg", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "keyless")
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

func (c *Client) RenameSystem(old_name string, new_name string) error {
	body := url.Values{}
	body.Add("newname", new_name)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/system/"+old_name, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "rename")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// ResetAccumulatedEnergy resets the accumulated energy consumption on the controller.
func (c *Client) ResetAccumulatedEnergy() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/system/energy", nil)
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

// RequestMastership requests mastership of all domains on the controller.
// ie. CFG, Motion, RAPID, etc.
func (c *Client) RequestMastershipAll() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/mastership", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "request")
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

// ReleaseMastership releases mastership of all domains on the controller.
// ie. CFG, Motion, RAPID, etc.
func (C *Client) ReleaseMastershipAll() error {
	C.Client = C.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+C.Host+"/rw/mastership", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "release")
	req.URL.RawQuery = q.Encode()
	resp, err := C.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// RequestMastershipIndividual requests mastership of a specific domain on the controller.
// ie. CFG, Motion, RAPID, etc.
func (c *Client) RequestMastershipIndividual(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}
	if domain != "cfg" && domain != "motion" && domain != "rapid" {
		return fmt.Errorf("invalid domain")
	}
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/mastership/"+domain, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "request")
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

// ReleaseMastershipIndividual releases mastership of a specific domain on the controller.
// ie. CFG, Motion, RAPID, etc.
func (C *Client) ReleaseMastershipIndividual(domain string) error {
	if domain == "" {
		return fmt.Errorf("domain cannot be empty")
	}
	if domain != "cfg" && domain != "motion" && domain != "rapid" {
		return fmt.Errorf("invalid domain")
	}
	C.Client = C.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+C.Host+"/rw/mastership/"+domain, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "release")
	req.URL.RawQuery = q.Encode()
	resp, err := C.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// CreateDIPCQueue creates a DIPC queue on the controller with the specified name, size, and max message size.
func (c *Client) CreateDIPCQueue(name string, size uint16, max_msg_size uint16) error {
	if max_msg_size < 1 || max_msg_size > 444 {
		return fmt.Errorf("max_msg_size must be between 1 and 444")
	}
	if size < 1 {
		return fmt.Errorf("size must be greater than 0")
	}
	size_str := fmt.Sprintf("%d", size)
	max_msg_size_str := fmt.Sprintf("%d", max_msg_size)
	body := url.Values{}
	body.Add("dipc-queue-name", name)
	body.Add("dipc-queue-size", size_str)
	body.Add("dipc-max-msg-size", max_msg_size_str)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/dipc", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "dipc-create")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}
