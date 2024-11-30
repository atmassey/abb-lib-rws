package abb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/atmassey/abb-lib-rws/structures"
)

// GetMechUnits returns a list of all the mechunits on the robot controller
func (c *Client) GetMechUnits() (*structures.MechUnits, error) {
	mechUnits := structures.MechUnitsJson{}
	mechUnitsDecoded := structures.MechUnits{}
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

// Returns the error state of the motion system
func (c *Client) GetErrorState() (*structures.MotionErrorState, error) {
	var motionErrorState structures.MotionErrorStateJson
	var motionErrorStateDecoded structures.MotionErrorState
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

// SetMotionSupervisionMode sets the motion supervision mode for a specific mechanical unit
func (c *Client) SetMotionSupervisionMode(MechanicalUnit string, Mode bool) error {
	body := url.Values{}
	body.Add("mode", fmt.Sprintf("%t", Mode))
	body.Add("mechunit-name", MechanicalUnit)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/motionsupervision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-mode")
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

// SetMotionSupervisionSensitivity sets the motion supervision sensitivity for a specific mechanical unit
func (c *Client) SetMotionSupervisionSensitivity(MechanicalUnit string, Sensitivity string) error {
	body := url.Values{}
	body.Add("sensitivity", Sensitivity)
	body.Add("mechunit-name", MechanicalUnit)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/motionsupervision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-level")
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

// SetPathSupervisionMode sets the path supervision mode for a specific mechanical unit
func (c *Client) SetPathSupervisionMode(Mode bool, MechUnit string) error {
	var ActualMode string
	if Mode {
		ActualMode = "ON"
	} else {
		ActualMode = "OFF"
	}
	body := url.Values{}
	body.Add("mode", ActualMode)
	body.Add("mechunit", MechUnit)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/pathsupervision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-mode")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetPathSupervisionLevel sets the path supervision level for a specific mechanical unit
func (c *Client) SetPathSupervisionLevel(Level string, MechUnit string) error {
	body := url.Values{}
	body.Add("level", Level)
	body.Add("mechunit", MechUnit)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/pathsupervision", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-level")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetNonMotionExecutionMode sets the non-motion execution mode
func (c *Client) SetNonMotionExecutionMode(Mode string) error {
	mode_buffered := strings.ToUpper(Mode)
	if mode_buffered != "ON" && mode_buffered != "OFF" {
		return fmt.Errorf("invalid mode: %s", Mode)
	}
	body := url.Values{}
	body.Add("mode", Mode)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/nonmotionexecution", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-mode")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetComplianceLeadThrough sets the compliance lead through for a specific mechanical unit
// Set Status to true to enable compliance lead through, false to disable
func (c *Client) SetComplianceLeadThrough(Mechunit string, Status bool) error {
	var status string
	if Status {
		status = "active"
	} else {
		status = "inactive"
	}
	body := url.Values{}
	body.Add("status", status)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+Mechunit, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-lead-through")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetFineCalibration sets the fine calibration for a specific mechanical unit
func (c *Client) SetFineCalibration(Mechunit string, AxisValue int) error {
	body := url.Values{}
	body.Add("axis", fmt.Sprintf("%d", AxisValue))
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+Mechunit, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "fine-calibrate")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetAxisPose sets the axis pose for a specific mechanical unit
func (c *Client) SetAxisPose(Mechunit string, Axisnum int, Positions structures.AxisPositon) error {
	axis_string := fmt.Sprintf("%d", Axisnum)
	body := url.Values{}
	body.Add("axis", fmt.Sprintf("%d", Axisnum))
	body.Add("x", Positions.X)
	body.Add("y", Positions.Y)
	body.Add("z", Positions.Z)
	body.Add("q1", Positions.Q1)
	body.Add("q2", Positions.Q2)
	body.Add("q3", Positions.Q3)
	body.Add("q4", Positions.Q4)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+Mechunit+"/axes/"+axis_string, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "set-axispose")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// SetSyncRevCounter will update the sync rev counter for a specific mechanical unit and axis
func (c *Client) UpdateSyncRevCounter(Mechunit string, Axis string) error {
	body := url.Values{}
	body.Add("syncType", "1")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+Mechunit+"/axes/"+Axis, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "update-syncrevcounter")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}

// UpdateCommutate will update the commutate for a specific mechanical unit and axis
func (c *Client) UpdateCommutate(Mechunit string, Axis string) error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/motionsystem/mechunits/"+Mechunit+"/axes/"+Axis, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "update-commutate")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("HTTP Status: %v", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}
