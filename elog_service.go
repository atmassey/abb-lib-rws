package abb

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

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

func (c *Client) SubscribeToElog(ResourceId int, Priority int) (chan string, error) {
	returnChannel := make(chan string)
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
	requestHeader.Add("Cookie", cookie1.String())
	requestHeader.Add("Cookie", cookie2.String())

	conn, _, err := websocket.DefaultDialer.Dial(ws_url, requestHeader)
	if err != nil {
		return nil, err
	}
	defer closeWSCheck(conn)
	go func() {
		defer func() {
			closeWSCheck(conn)
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			returnChannel <- string(message)
		}
	}()
	return returnChannel, nil
}
