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

// SubscribeToIOSignal is used to subscribe to an IO signal and returns a channel with the signal value and simulation state.
// Example signal: LOCAL/PANEL/MAN1 for manual mode
func (c *Client) SubscribeToIOSignal(Signal string) (chan map[string]string, error) {
	returnChannel := make(chan map[string]string)
	body := url.Values{}
	body.Add("resources", "1")
	body.Add("1", "/rw/iosystem/signals/"+Signal+";state")
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
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			MessageXML := structures.IOSignalXML{}
			err = xml.Unmarshal(message, &MessageXML)
			if err != nil {
				return
			}
			mapString := make(map[string]string)
			mapString["value"] = MessageXML.Body.Div.List.Span[0].Text
			mapString["state"] = MessageXML.Body.Div.List.Span[1].Text
			returnChannel <- mapString
		}
	}()
	return returnChannel, nil
}

// UnblockSignals will remove simulation for all simulated logical I/O signals.
func (c *Client) UnblockSignals() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("DELETE", "http://"+c.Host+"/rw/iosystem/signals", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("action", "unblock-signal")
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
