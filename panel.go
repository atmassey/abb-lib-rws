package abb

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/atmassey/abb-lib-rws/structures"
	"github.com/gorilla/websocket"
)

type RestartController interface {
	Warmstart() error
	IStart() error
	PStart() error
	BStart() error
}

// CAUTION: A warmstart will restart the controller and all running programs will be stopped. (Warmstart)
func (c *Client) Warmstart() error {
	return c.RestartController("restart")
}

// CAUTION: A "istart" will restart the controller and factory reset the controller.
func (c *Client) IStart() error {
	return c.RestartController("istart")
}

// CAUTION: A "pstart" will restart the controller and delete all rapid programs but keep all configuration data.
func (c *Client) PStart() error {
	return c.RestartController("pstart")
}

// CAUTION: A "bstart" will restart the controller and revert it to its last auto-saved state.
func (c *Client) BStart() error {
	return c.RestartController("bstart")
}

// RestartController is used to restart the controller with the specified action.
// Possible values: {restart | istart | pstart | bstart}
func (c *Client) RestartController(Action string) error {
	body := url.Values{}
	body.Add("restart-mode", Action)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel", bytes.NewBufferString(body.Encode()))
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

// GetOperationMode returns the current operation mode of the controller.
// Possible values: (INIT | AUTO_CH | MANF_CH | MANR | MANF | AUTO | UNDEF)
func (c *Client) GetOperationMode() (string, error) {
	var opmode structures.OperationMode
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/panel/opmode", nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("json", "1")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&opmode)
	if err != nil {
		return "", err
	}
	mode := opmode.Embedded.State[0].Opmode
	if mode == "" {
		return "", fmt.Errorf("OP Mode Not Found: %v", opmode)
	}
	return mode, nil
}

// SubscribeToControllerState is subscribed to the controller state websocket
// that will send an update anytime the controller state changes.
// The map key returned is mapString["state"].
// Possible states are {init | motoron | motoroff | guardstop | emergencystop | emergencystopreset | sysfail}
func (c *Client) SubscribeToControllerState() (chan map[string]string, error) {
	returnChannel := make(chan map[string]string)
	body := url.Values{}
	body.Add("resources", "1")
	body.Add("1", "/rw/panel/ctrlstate")
	body.Add("1-p", "1")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/subscription", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer closeErrorCheck(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	ws_url := resp.Header.Get("Location")
	header := resp.Cookies()
	var session, session_ab string
	for _, c := range header {
		if c.Name == "-http-session-" {
			session = c.Value
		} else if c.Name == "ABBCX" {
			session_ab = c.Value
		}
	}
	requestHeader := http.Header{}
	cookie1 := &http.Cookie{Name: "-http-session-", Value: session}
	cookie2 := &http.Cookie{Name: "ABBCX", Value: session_ab}
	requestHeader.Add("Cookie", cookie1.String()+"; "+cookie2.String())
	requestHeader.Add("Origin", strings.Split(ws_url, "/poll")[0])
	requestHeader.Add("Sec-WebSocket-Protocol", "robapi2_subscription")
	conn, _, err := websocket.DefaultDialer.Dial(ws_url, requestHeader)
	if err != nil {
		return nil, err
	}
	go func() {
		defer func() {
			conn.Close()
			close(returnChannel)
		}()
		err = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		if err != nil {
			return
		}
		conn.SetPongHandler(func(string) error {
			if err := conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
				return err
			}
			return nil
		})
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			MessageXML := structures.PanelXML{}
			err = xml.Unmarshal(message, &MessageXML)
			if err != nil {
				return
			}
			mapString := make(map[string]string)
			mapString["state"] = MessageXML.Body.Div.List.Span.Text
			returnChannel <- mapString
		}
	}()
	return returnChannel, nil
}

// SubscribeToOperationMode is subscribed to the operation mode websocket
// that will send an update anytime the operation mode changes.
// The map key returned is mapString["mode"].
// Possible states are {INIT | AUTO_CH | MANF_CH | MANR | MANF | AUTO | UNDEF}
func (c *Client) SubscribeToOperationMode() (chan map[string]string, error) {
	returnChannel := make(chan map[string]string)
	body := url.Values{}
	body.Add("resources", "1")
	body.Add("1", "/rw/panel/opmode")
	body.Add("1-p", "1")
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/subscription", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer closeErrorCheck(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	ws_url := resp.Header.Get("Location")
	header := resp.Cookies()
	var session, session_ab string
	for _, c := range header {
		if c.Name == "-http-session-" {
			session = c.Value
		} else if c.Name == "ABBCX" {
			session_ab = c.Value
		}
	}
	requestHeader := http.Header{}
	cookie1 := &http.Cookie{Name: "-http-session-", Value: session}
	cookie2 := &http.Cookie{Name: "ABBCX", Value: session_ab}
	requestHeader.Add("Cookie", cookie1.String()+"; "+cookie2.String())
	requestHeader.Add("Origin", strings.Split(ws_url, "/poll")[0])
	requestHeader.Add("Sec-WebSocket-Protocol", "robapi2_subscription")
	conn, _, err := websocket.DefaultDialer.Dial(ws_url, requestHeader)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	go func() {
		defer func() {
			close(returnChannel)
		}()
		err = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		if err != nil {
			return
		}
		conn.SetPongHandler(func(string) error {
			if err := conn.SetReadDeadline(time.Now().Add(60 * time.Second)); err != nil {
				return err
			}
			return nil
		})
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			MessageXML := structures.PanelXML{}
			err = xml.Unmarshal(message, &MessageXML)
			if err != nil {
				return
			}
			mapString := make(map[string]string)
			mapString["mode"] = MessageXML.Body.Div.List.Span.Text
			returnChannel <- mapString
		}
	}()
	return returnChannel, nil
}

// AcknowledgeOpMode is used to acknowledge the operation mode change.
func (c *Client) AcknowledgeOpMode(Mode string) error {
	if Mode != "auto" && Mode != "manf" && Mode != "coldet" {
		return fmt.Errorf("invalid mode %s", Mode)
	}
	body := url.Values{}
	body.Add("opmode", Mode)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel/opmode", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "acknowledge")
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

// LockOpMode is used to lock the operation mode selection.
func (c *Client) LockOpMode(Pin int16, Permanent bool) error {
	if Pin < 0 || Pin > 9999 {
		return fmt.Errorf("invalid pin %d", Pin)
	}
	body := url.Values{}
	body.Add("pin", fmt.Sprintf("%04d", Pin))
	if Permanent {
		body.Add("permanent", "1")
	} else {
		body.Add("permanent", "0")
	}
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel/opmode", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "lock")
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

// UnlockOpMode is used to unlock the operation mode selection.
func (c *Client) UnlockOpMode(Pin int16) error {
	if Pin < 0 || Pin > 9999 {
		return fmt.Errorf("invalid pin %d", Pin)
	}
	body := url.Values{}
	body.Add("pin", fmt.Sprintf("%v", Pin))
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel/opmode", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "unlock")
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

// SetSpeedRatio is used to set the speed ratio of the controller. The value should be between 0 and 100.
func (c *Client) SetSpeedRatio(SpeedRatio int8) error {
	if SpeedRatio < 0 || SpeedRatio > 100 {
		return fmt.Errorf("invalid speed ratio %d", SpeedRatio)
	}
	body := url.Values{}
	body.Add("speed-ratio", fmt.Sprintf("%d", SpeedRatio))
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/panel/speedratio", bytes.NewBufferString(body.Encode()))
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
