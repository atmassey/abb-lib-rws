package abb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// GetMechUnits returns a list of all the mechunits on the robot controller
func (c *Client) GetMechUnits() (*MechUnits, error) {
	mechUnits := MechUnitsJson{}
	mechUnitsDecoded := MechUnits{}
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/motionsystem/mechunits", nil)
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
		return nil, fmt.Errorf("HTTP Status: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&mechUnits)
	if err != nil {
		return nil, err
	}
	for _, mechUnit := range mechUnits.Embedded.State {
		mechUnitsDecoded.ActivationAllowed = append(mechUnitsDecoded.ActivationAllowed, mechUnit.ActivationAllowed)
		mechUnitsDecoded.DriveModule = append(mechUnitsDecoded.DriveModule, mechUnit.DriveModule)
		mechUnitsDecoded.Mode = append(mechUnitsDecoded.Mode, mechUnit.Mode)
		mechUnitsDecoded.Title = append(mechUnitsDecoded.Title, mechUnit.Title)
	}
	defer closeErrorCheck(resp.Body)
	return &mechUnitsDecoded, nil
}

// ClearSMBData clears the SMB data for a specific mechunit
// type_ can be either "robot" or "controller"
func (c *Client) ClearSMBData(MechUnit string, type_ string) error {
	body := url.Values{}
	body.Add("type", type_)
	if type_ != "robot" && type_ != "controller" {
		return fmt.Errorf("invalid type: %s", type_)
	}
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+MechUnit+"/smbdata", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "clear")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %s", resp.Status)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

func (c *Client) GetErrorState() (*MotionErrorState, error) {
	var motionErrorState MotionErrorStateJson
	var motionErrorStateDecoded MotionErrorState
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/motionsystem/errorstate", nil)
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
		return nil, fmt.Errorf("HTTP Status: %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&motionErrorState)
	if err != nil {
		return nil, err
	}
	for _, motionError := range motionErrorState.Embedded.State {
		motionErrorStateDecoded.State = append(motionErrorStateDecoded.State, motionError.State)
		motionErrorStateDecoded.Count = append(motionErrorStateDecoded.Count, motionError.Count)
	}
	defer closeErrorCheck(resp.Body)
	return &motionErrorStateDecoded, nil
}
