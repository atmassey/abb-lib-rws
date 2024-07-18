package abb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *Client) GetIOSignals() (*IOSignals, error) {
	var signals IOSignals
	var signals_raw IOSignalsJson
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
	err = json.NewDecoder(resp.Body).Decode(&signals_raw)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("Syntax error at byte offset %d", e.Offset)
		}
		log.Printf("Raw JSON: %v", signals_raw)
		return nil, err
	}
	defer resp.Body.Close()
	for _, signal := range signals_raw.Embedded.State {
		signals.SignalName = append(signals.SignalName, signal.Name)
		signals.SignalType = append(signals.SignalType, signal.Type)
		signals.SignalValue = append(signals.SignalValue, signal.Value)
	}
	return &signals, nil
}
