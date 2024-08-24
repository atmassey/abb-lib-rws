package abb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

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
