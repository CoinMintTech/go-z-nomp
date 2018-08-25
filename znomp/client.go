package znomp

import (
	"github.com/steenzout/go-cfg"
)

// Client z-nomp API client.
type Client struct {
	authenticate AuthFunc
	endpoint     string
}

// NewClient returns a new z-nomp API client.
func NewClient(c cfg.Config) *Client {
	return &Client{
		authenticate: AuthenticationFunc(c),
		endpoint:     Endpoint(c),
	}
}
