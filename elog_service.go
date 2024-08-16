package abb

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

func (c *Client) SubscribeToElog(ResourceId int, Priority int) (chan ElogXML, error) {
	returnChannel := make(chan ElogXML)
	string_id := strconv.Itoa(ResourceId)
	string_priority := strconv.Itoa(Priority)
	body := url.Values{}
	body.Add("resources", string_id)
	body.Add(string_id, "/rw/elog/"+string_id)
	body.Add(string_id+"-p", string_priority)
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
			fmt.Println("Session: ", session)
		} else if c.Name == "ABBCX" {
			session_ab = c.Value
			fmt.Println("Session AB: ", session_ab)
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
	fmt.Println("Connected to websocket")
	go func() {
		defer func() {
			fmt.Printf("Closing connection\n")
			conn.Close()
			close(returnChannel)
		}()
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil })
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				return
			}
			MessageXML := ElogXML{}
			fmt.Printf("Received message: %s\n", message)
			err = xml.Unmarshal(message, &MessageXML)
			if err != nil {
				fmt.Printf("Error unmarshalling message: %v\n", err)
				return
			}
			seqnum := MessageXML.Body.Div.List.Span.Text
			endpoint := MessageXML.Body.Div.List.Endpoint.Href
			fmt.Printf("Seqnum: %s Endpoint: %s", seqnum, endpoint)
			returnChannel <- MessageXML
		}
	}()
	return returnChannel, nil
}
