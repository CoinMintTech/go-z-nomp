package znomp

import (
	"fmt"
	"log"
	"strings"

	"github.com/steenzout/go-cfg"
	"github.com/steenzout/go-env"
)

const (
	envAuth             = "AUTH"
	envAuthHTTPPassword = "AUTH_HTTP_PASSWORD"
	envAuthHTTPUser     = "AUTH_HTTP_USER"
)

// Config returns package configuration.
func Config(c cfg.Config) cfg.Config {
	v := strings.ToLower(strings.TrimSpace(env.GetOptionalString(envAuth, "http")))

	if v == "http" {
		username := env.GetString(envAuthHTTPUser)
		password := env.GetString(envAuthHTTPPassword)

		c = WithAuthentication(c, HTTPBasicAuth)
		c = WithAuthHTTPassword(c, password)
		c = WithAuthHTTPUser(c, username)

		return c
	}

	errm := fmt.Sprintf("invalid authentication method: %s", v)
	log.Println(errm)
	panic(errm)
}
