package znomp

import (
	"net/http"

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

// Blocks returns the global blocks stats.
func (c Client) Blocks() error {
	var (
		err error
		req *http.Request
	)

	req, err = http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return err
	}

	err = c.authenticate(req)
	if err != nil {
		return err
	}

	return nil
}
