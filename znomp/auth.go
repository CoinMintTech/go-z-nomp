package znomp

import (
	"net/http"
)

// Auth type of API authentication to use.
type Auth uint8

const (
	// HTTPBasicAuth constants for HTTP Basic Authentication.
	HTTPBasicAuth = Auth(iota)
)

// AuthFunc HTTP client authentication function.
type AuthFunc func(req *http.Request) error

func httpAuth(u, p string) AuthFunc {
	return func(req *http.Request) error {
		req.SetBasicAuth(u, p)

		return nil
	}
}
