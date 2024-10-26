package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/atmassey/abb-lib-rws/structures"
	"github.com/gorilla/websocket"
)

// SaveElogSystemDump dumps log file to the specified path on the controller.
// Example path: $HOME/my_dump_file.txt
func (c *Client) SaveElogSystemDump(Path string) error {
	body := url.Values{}
	body.Add("path", "/fileservice/"+Path)
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/elog", bytes.NewBufferString(body.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("action", "saveraw")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return nil
}

// getElogMessages is a helper function that returns the messages from the Elog system based on the endpoint.
// This function is used in conjunction with SubscribeToElog for the Elog websocket.
func (c *Client) getElogMessages(Endpoint string) (*structures.ElogMessagesXML, error) {
	var messages structures.ElogMessagesXML
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+Endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("lang", "en")
	req.URL.RawQuery = q.Encode()
	resp, err := c.Client.Do(req)
	defer closeErrorCheck(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %w", err)
	}
	messages_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(messages_raw, &messages)
	if err != nil {
		return nil, err
	}
	return &messages, nil
}

// SubscribeToElog subscribes to the Elog websocket for all events happening at the robot.
// This function returns a map of strings. The keys within the map are as follows
// "msgtype", "code", "tstamp", "title", "desc", "conseqs", "causes", "actions", "argc",
// "arg1", and "arg2".
func (c *Client) SubscribeToElog() (chan map[string]string, error) {
	returnChannel := make(chan map[string]string)
	body := url.Values{}
	body.Add("resources", "1")
	body.Add("1", "/rw/elog/1")
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
		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			MessageXML := structures.ElogXML{}
			err = xml.Unmarshal(message, &MessageXML)
			if err != nil {
				return
			}
			endpoint := MessageXML.Body.Div.List.Endpoint.Href
			msg, err := c.getElogMessages(endpoint)
			if err != nil {
				continue
			}
			mapString := make(map[string]string)
			for _, m := range msg.Body.Div.List.Span {
				mapString[m.Class] = m.Text
			}
			returnChannel <- mapString
		}
	}()
	return returnChannel, nil
}

// ClearElogMessages clears all messages from the Elog system on domain 0.
func (c *Client) ClearElogMessages() error {
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("POST", "http://"+c.Host+"/rw/elog/0", nil)
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
		return fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	defer closeErrorCheck(resp.Body)
	return nil
}
