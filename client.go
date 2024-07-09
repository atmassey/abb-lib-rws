package abb

import (
	"net/http"

	"github.com/icholy/digest"
)

func NewClient(IP string, Username string, Password string) *Client {
	abb := new(Client)
	abb.IP = IP
	abb.Username = Username
	abb.Password = Password
	abb.Client = &http.Client{}
	return abb
}

func (c *Client) GetIP() string {
	return c.IP
}

func (c *Client) GetUsername() string {
	return c.Username
}

func (c *Client) GetPassword() string {
	return c.Password
}

func (c *Client) SetIP(IP string) {
	c.IP = IP
}

func (c *Client) SetUsername(Username string) {
	c.Username = Username
}

func (c *Client) SetPassword(Password string) {
	c.Password = Password
}

func (c *Client) DigestAuthenticate() *http.Client {
	client := &http.Client{Transport: &digest.Transport{Username: c.Username, Password: c.Password}}
	return client
}
