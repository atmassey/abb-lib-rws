package structures

import "encoding/xml"

type CameraStatusRaw struct {
	XMLName xml.Name         `xml:"html"`
	Head    CameraStatusHead `xml:"head"`
	Body    CameraStatusBody `xml:"body"`
}

type CameraStatusHead struct {
	Title string `xml:"title"`
	Base  string `xml:"base,attr"`
}

type CameraStatusBody struct {
	Div CameraStatusDiv `xml:"div"`
}

type CameraStatusDiv struct {
	Class string             `xml:"class,attr"`
	Link  CameraStatusLink   `xml:"a"`
	List  []CameraStatusList `xml:"ul>li"`
}

type CameraStatusLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type CameraStatusList struct {
	Class string           `xml:"class,attr"`
	Title string           `xml:"title,attr"`
	Span  CameraStatusSpan `xml:"span"`
}

type CameraStatusSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}
