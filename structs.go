package abb

import (
	"net/http"
)

type Client struct {
	Host     string
	Username string
	Password string
	Client   *http.Client
}
