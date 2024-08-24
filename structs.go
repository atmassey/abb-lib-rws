package abb

import (
	"encoding/xml"
	"net/http"
)

type Client struct {
	Host     string
	Username string
	Password string
	Client   *http.Client
}

type RobotType struct {
	XMLName xml.Name      `xml:"html"`
	Head    RobotTypeHead `xml:"head"`
	Body    RobotTypeBody `xml:"body"`
}

type RobotTypeHead struct {
	XMLName xml.Name `xml:"head"`
	Title   string   `xml:"title"`
	Base    string   `xml:"base,attr"`
}

type RobotTypeBody struct {
	XMLName xml.Name       `xml:"body"`
	State   RobotTypeState `xml:"div"`
}

type RobotTypeState struct {
	XMLName xml.Name         `xml:"div"`
	Self    RobotTypeSelf    `xml:"a"`
	Robots  []RobotTypeRobot `xml:"ul>li"`
}

type RobotTypeSelf struct {
	XMLName xml.Name `xml:"a"`
	Rel     string   `xml:"rel,attr"`
}

type RobotTypeRobot struct {
	XMLName   xml.Name `xml:"li"`
	Title     string   `xml:"title,attr"`
	RobotType string   `xml:"span"`
}

type IOSignalsJson struct {
	Links    IOSignalsJsonLinks `json:"_links"`
	Embedded IOSignalsJsonState `json:"_embedded"`
}

type IOSignalsJsonLinks struct {
	Base IOSignalsJsonBase `json:"base"`
}

type IOSignalsJsonBase struct {
	Href string `json:"href"`
}

type IOSignalsJsonState struct {
	State []IOSignalsJsonMeta `json:"_state"`
}

type IOSignalsJsonMeta struct {
	Links           IOSignalsJsonMetaLinks `json:"_links"`
	TypeT           string                 `json:"_type"`
	Name            string                 `json:"name"`
	Type            string                 `json:"type"`
	Category        string                 `json:"category"`
	Value           int                    `json:"lvalue"`
	SimulationState string                 `json:"lstate"`
}

type IOSignalsJsonMetaLinks struct {
	Self IOSignalsJsonMetaLinksSelf `json:"self"`
}

type IOSignalsJsonMetaLinksSelf struct {
	Href string `json:"href"`
}

type IOSignals struct {
	SignalName  []string
	SignalType  []string
	SignalValue []int
}
