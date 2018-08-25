package znomp

import (
	"fmt"
	"log"
	"strings"

	"github.com/steenzout/go-cfg"
	"github.com/steenzout/go-env"
)

const (
	envAuth             = "ZNOMP_AUTH"
	envAuthHTTPPassword = "ZNOMP_AUTH_HTTP_PASSWORD"
	envAuthHTTPUser     = "ZNOMP_AUTH_HTTP_USER"
	envEndpoint         = "ZNOMP_ENDPOINT"
)

// Config returns package configuration.
func Config(c cfg.Config) cfg.Config {
	endpoint := strings.ToLower(strings.TrimSpace(env.GetString(envEndpoint)))
	c = WithEndpoint(c, endpoint)

	auth := strings.ToLower(strings.TrimSpace(env.GetOptionalString(envAuth, "http")))

	if auth == "http" {
		username := env.GetString(envAuthHTTPUser)
		password := env.GetString(envAuthHTTPPassword)

		return WithAuthenticationFunc(c, httpAuth(username, password))
	}

	errm := fmt.Sprintf("invalid authentication method: %s", auth)
	log.Println(errm)
	panic(errm)
}
