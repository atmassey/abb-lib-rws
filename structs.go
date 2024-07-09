package abb

import "net/http"

type Client struct {
	IP       string
	Username string
	Password string
	Client   *http.Client
}
