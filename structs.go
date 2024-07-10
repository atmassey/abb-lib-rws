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

type IOSignalsHTML struct {
	XMLName xml.Name      `xml:"html"`
	Head    IOSignalsHead `xml:"head"`
	Body    IOSignalsBody `xml:"body"`
}

type IOSignalsHead struct {
	XMLName xml.Name      `xml:"head"`
	Title   string        `xml:"title"`
	Base    IOSignalsBase `xml:"base"`
}

type IOSignalsBase struct {
	XMLName xml.Name `xml:"base"`
	Href    string   `xml:"href,attr"`
}

type IOSignalsBody struct {
	XMLName xml.Name     `xml:"body"`
	Div     IOSignalsDiv `xml:"div"`
}

type IOSignalsDiv struct {
	XMLName xml.Name        `xml:"div"`
	Class   string          `xml:"class,attr"`
	Links   []IOSignalsLink `xml:"a"`
	UL      IOSignalsUL     `xml:"ul"`
}

type IOSignalsLink struct {
	XMLName xml.Name `xml:"a"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
}

type IOSignalsUL struct {
	XMLName xml.Name      `xml:"ul"`
	LIs     []IOSignalsLI `xml:"li"`
}

type IOSignalsLI struct {
	XMLName  xml.Name      `xml:"li"`
	Class    string        `xml:"class,attr"`
	Title    string        `xml:"title,attr"`
	Link     IOSignalsLink `xml:"a"`
	Name     string        `xml:"span.name"`
	Type     string        `xml:"span.type"`
	Category string        `xml:"span.category"`
	LValue   string        `xml:"span.lvalue"`
	LState   string        `xml:"span.lstate"`
}
