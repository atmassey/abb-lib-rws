package structures

import "encoding/xml"

type PanelXML struct {
	XMLName xml.Name  `xml:"html"`
	Head    PanelHead `xml:"head"`
	Body    PanelBody `xml:"body"`
}

type PanelHead struct {
	Title string `xml:"title"`
}

type PanelBody struct {
	Div   PanelDiv `xml:"div"`
	Class string   `xml:"class,attr"`
}

type PanelDiv struct {
	Poll []PanelLink `xml:"a"`
	List PanelList   `xml:"ul>li"`
}

type PanelLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:",chardata"`
}

type PanelLinkName struct {
	Href string `xml:"href,attr"`
	Name string `xml:"name,attr"`
}

type PanelList struct {
	Href PanelMetaLink `xml:"a"`
	Span PanelSpan     `xml:"span"`
}

type PanelMetaLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type PanelSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
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
