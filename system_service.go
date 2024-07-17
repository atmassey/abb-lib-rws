package abb

import (
	"encoding/json"
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

func (c *Client) GetSystemEnergyMetrics() (*SystemEnergyMetrics, error) {
	var EnergyMetricsDecoded SystemEnergyMetrics
	var EnergyMetricsRaw SystemEnergy
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
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&EnergyMetricsRaw)
	if err != nil {
		return nil, err
	}
	EnergyMetricsDecoded.AccumulatedEnergy = EnergyMetricsRaw.Embedded.State[0].AccumulatedEnergy
	for _, state := range EnergyMetricsRaw.Embedded.State {
		for _, MechUnits := range state.MechUnits {
			for _, axis := range MechUnits.Axis {
				axisEnergy := SystemAxisEnergy{Axis: axis.Title, Energy: axis.IntervalEnergy}
				EnergyMetricsDecoded.AxisEnergy = append(EnergyMetricsDecoded.AxisEnergy, axisEnergy)
			}
		}
	}
	return &EnergyMetricsDecoded, nil
}

func (c *Client) GetInstalledProducts() (*InstalledSystemProducts, error) {
	var InstalledProducts InstalledProducts
	var InstalledProductsDecoded InstalledSystemProducts
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
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&InstalledProducts)
	if err != nil {
		return nil, err
	}
	installedProducts := InstalledSystemProducts{}
	for _, product := range InstalledProducts.State {
		title := product.Title
		version := product.VersionName
		InstalledProductsDecoded.Title = append(installedProducts.Title, title)
		InstalledProductsDecoded.Version = append(installedProducts.Version, version)
	}
	return &InstalledProductsDecoded, nil
}
