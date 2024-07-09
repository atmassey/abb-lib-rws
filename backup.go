package abb

import "net/http"

func (c *Client) GetBackupResources() (*http.Response, error) {
	req, err := http.NewRequest("GET", "http://"+c.IP+"/backup/resources?json=1", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}
