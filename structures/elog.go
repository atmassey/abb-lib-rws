package structures

import "encoding/xml"

type ElogXML struct {
	XMLName xml.Name `xml:"html"`
	Head    ElogHead `xml:"head"`
	Body    ElogBody `xml:"body"`
}

type ElogHead struct {
	Title string `xml:"title"`
}

type ElogBody struct {
	Div ElogDiv `xml:"div"`
}

type ElogDiv struct {
	Class string   `xml:"class,attr"`
	Link  ElogLink `xml:"a"`
	List  ElogList `xml:"ul>li"`
}

type ElogLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type ElogList struct {
	Class    string       `xml:"class,attr"`
	Endpoint ElogEndpoint `xml:"a"`
	Span     ElogSpan     `xml:"span"`
}

type ElogSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}

type ElogEndpoint struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:",chardata"`
}

type ElogMessagesXML struct {
	XMLName xml.Name        `xml:"html"`
	Head    ElogMessageHead `xml:"head"`
	Body    ElogMessageBody `xml:"body"`
}

type ElogMessageHead struct {
	Title string `xml:"title"`
	Base  string `xml:"base,attr"`
}

type ElogMessageBody struct {
	Div ElogMessageDiv `xml:"div"`
}

type ElogMessageDiv struct {
	Class string          `xml:"class,attr"`
	Link  ElogMessageLink `xml:"a"`
	List  ElogMessageList `xml:"ul>li"`
}

type ElogMessageLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type ElogMessageList struct {
	Class string            `xml:"class,attr"`
	Title string            `xml:"title"`
	Span  []ElogMessageSpan `xml:"span"`
}

type ElogMessageSpan struct {
	Class string `xml:"class,attr"`
	Text  string `xml:",chardata"`
}
