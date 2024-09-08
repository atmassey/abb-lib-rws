package structures

import "encoding/xml"

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

type IOSignalXML struct {
	XMLName xml.Name     `xml:"html"`
	Head    IOSignalHead `xml:"head"`
	Body    IOSignalBody `xml:"body"`
}

type IOSignalHead struct {
	Title string `xml:"title"`
}

type IOSignalBody struct {
	Div   IOSignalDiv `xml:"div"`
	Class string      `xml:"class,attr"`
}

type IOSignalDiv struct {
	Poll []IOSignalLink `xml:"a"`
	List IOSignalList   `xml:"ul>li"`
}

type IOSignalLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:",chardata"`
}

type IOSignalList struct {
	Href IOSignalMetaLink `xml:"a"`
	Span []IOSignalSpan   `xml:"span"`
}

type IOSignalMetaLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type IOSignalSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}
