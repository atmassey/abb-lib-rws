package abb

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetIOSignals() (*IOSignals, error) {
	var signals IOSignalsHTML
	var signals_struct IOSignals
	c.Client = c.DigestAuthenticate()
	req, err := http.NewRequest("GET", "http://"+c.Host+"/rw/iosystem/signals", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Status Code: %d", resp.StatusCode)
	}
	signals_raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(signals_raw, &signals)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	for _, signal := range signals.Body.Div.UL.LIs {
		for _, span := range signal.Spans {
			switch span.Class {
			case "name":
				signals_struct.SignalName = append(signals_struct.SignalName, span.Content)
			case "type":
				signals_struct.SignalType = append(signals_struct.SignalType, span.Content)
			case "lvalue":
				signals_struct.SignalValue = append(signals_struct.SignalValue, span.Content)
			}
		}
	}
	return &signals_struct, nil
}
