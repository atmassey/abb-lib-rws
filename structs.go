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

type FileResources struct {
	XMLName xml.Name        `xml:"html"`
	Body    FileServiceBody `xml:"body"`
}

type FileServiceBody struct {
	Div FileServiceDiv `xml:"div"`
}

type FileServiceDiv struct {
	Class string          `xml:"class,attr"`
	Lists []FileServiceLi `xml:"ul>li"`
}

type FileServiceLi struct {
	Class      string           `xml:"class,attr"`
	Title      string           `xml:"title,attr"`
	ParentLink *FileServiceLink `xml:"a"`
	SelfLink   *FileServiceLink `xml:"a"`
	DeviceType FileServiceSpan  `xml:"span[class='fs-device-type']"`
	FreeSpace  FileServiceSpan  `xml:"span[class='fs-free-space']"`
	TotalSpace FileServiceSpan  `xml:"span[class='fs-total-space']"`
	Enabled    FileServiceSpan  `xml:"span[class='fs-enabled']"`
	ReadOnly   FileServiceSpan  `xml:"span[class='fs-readonly']"`
}

type FileServiceLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
}

type FileServiceSpan struct {
	Text string `xml:",chardata"`
}
