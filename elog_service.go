package abb

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// Dumps elog file to the specified path on the controller.
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
