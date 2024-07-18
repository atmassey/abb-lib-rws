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

type ControllerActions struct {
	Actions []string
}

type ControllerActionsHTML struct {
	XMLName xml.Name              `xml:"html"`
	Body    ControllerActionsBody `xml:"body"`
}

type ControllerActionsBody struct {
	Div ControllerActionsDiv `xml:"div"`
}

type ControllerActionsDiv struct {
	Select ControllerActionsSelect `xml:"form>select"`
}

type ControllerActionsSelect struct {
	Options []ControllerActionsOption `xml:"option"`
}

type ControllerActionsOption struct {
	Value string `xml:"value,attr"`
}

type ControllerResources struct {
	XMLName xml.Name                `xml:"html"`
	Head    ControllerResourcesHead `xml:"head"`
	Body    ControllerResourcesBody `xml:"body"`
}

type ControllerResourcesHead struct {
	Title string                  `xml:"title"`
	Base  ControllerResourcesBase `xml:"base"`
}

type ControllerResourcesBase struct {
	Href string `xml:"href,attr"`
}

type ControllerResourcesBody struct {
	Div ControllerResourcesDiv `xml:"div"`
}

type ControllerResourcesDiv struct {
	Class string                    `xml:"class,attr"`
	Links []ControllerResourcesLink `xml:"a"`
	Lists []ControllerResourcesLi   `xml:"ul>li"`
}

type ControllerResourcesLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type ControllerResourcesLi struct {
	Class string                    `xml:"class,attr"`
	Title string                    `xml:"title,attr"`
	Link  ControllerResourcesLink   `xml:"a"`
	Spans []ControllerResourcesSpan `xml:"span"`
}

type ControllerResourcesSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}

type UserResources struct {
	XMLName xml.Name          `xml:"html"`
	Head    UserResourcesHead `xml:"head"`
	Body    UserResourcesBody `xml:"body"`
}

type UserResourcesHead struct {
	Title string            `xml:"title"`
	Base  UserResourcesBase `xml:"base"`
}

type UserResourcesBase struct {
	Href string `xml:"href,attr"`
}

type UserResourcesBody struct {
	Div UserResourcesDiv `xml:"div"`
}

type UserResourcesDiv struct {
	Class string              `xml:"class,attr"`
	Links []UserResourcesLink `xml:"a"`
	Lists []UserLi            `xml:"ul>li"`
}

type UserResourcesLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type UserLi struct {
	Class string            `xml:"class,attr"`
	Title string            `xml:"title,attr"`
	Link  UserResourcesLink `xml:"a"`
	Span  UserResourcesSpan `xml:"span"`
}

type UserResourcesSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
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

type OperationMode struct {
	Links    OperationModeBase  `json:"_links"`
	Embedded OperationModeState `json:"_embedded"`
}
type OperationModeLinks struct {
	Base OperationModeBase `json:"base"`
}

type OperationModeBase struct {
	Href string `json:"href"`
}
type OperationModeState struct {
	State []OperationModeMeta `json:"_state"`
}

type OperationModeMeta struct {
	Type   string `json:"_type"`
	Title  string `json:"_title"`
	Opmode string `json:"opmode"`
}

type SystemEnergy struct {
	Links    SystemEneryLinks     `json:"_links"`
	Embedded SystemEnergyEmbedded `json:"_embedded"`
}

type SystemEneryLinks struct {
	Base SystemEneryBase `json:"base"`
}

type SystemEneryBase struct {
	Href string `json:"href"`
}

type SystemEnergyEmbedded struct {
	State []StateListLinks `json:"_state"`
}

type StateListLinks struct {
	Links             StateListLinksSelf `json:"_links"`
	Type              string             `json:"_type"`
	Title             string             `json:"_title"`
	State             string             `json:"state"`
	EnergyState       string             `json:"energy-state"`
	ChangeCount       string             `json:"change-count"`
	TimeStamp         string             `json:"time-stamp"`
	ResetTime         string             `json:"reset-time"`
	IntervalLength    string             `json:"interval-length"`
	IntervalEnergy    string             `json:"interval-energy"`
	AccumulatedEnergy string             `json:"accumulated-energy"`
	MechUnits         []SystemEnergyAxes `json:"mechunits"`
}

type StateListLinksSelf struct {
	Self StateListLinksSelfHref `json:"self"`
}

type StateListLinksSelfHref struct {
	Href string `json:"href"`
}

type SystemEnergyAxes struct {
	Type  string             `json:"_type"`
	Title string             `json:"_title"`
	Axis  []SystemEnergyAxis `json:"axes"`
}

type SystemEnergyAxis struct {
	Type           string `json:"_type"`
	Title          string `json:"_title"`
	IntervalEnergy string `json:"interval-energy"`
}

type SystemEnergyMetrics struct {
	AccumulatedEnergy string
	AxisEnergy        []SystemAxisEnergy
}

type SystemAxisEnergy struct {
	Axis   string
	Energy string
}

type InstalledSystemProducts struct {
	Title   []string
	Version []string
}

type InstalledProducts struct {
	Links InstalledProductsLinks   `json:"_links"`
	State []InstalledProductsState `json:"state"`
}

type InstalledProductsLinks struct {
	Base InstalledProductsBase `json:"base"`
	Self InstalledProductsSelf `json:"self"`
}

type InstalledProductsBase struct {
	Href string `json:"href"`
}

type InstalledProductsSelf struct {
	Href string `json:"href"`
}

type InstalledProductsState struct {
	Type        string `json:"_type"`
	Title       string `json:"_title"`
	VersionName string `json:"version-name"`
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
