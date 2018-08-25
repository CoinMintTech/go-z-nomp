//
// Copyright 2018 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package znomp

import (
	"github.com/steenzout/go-cfg"
)

type configKey string

const (
	keyAuth             = configKey(iota)
	keyAuthHTTPPassword = configKey(iota)
	keyAuthHTTPUser     = configKey(iota)
)

// Authentication returns the authentication method.
func Authentication(c cfg.Config) Auth {
	return c.Value(keyAuth).(Auth)
}

// AuthHTTPassword returns the password for HTTP Basic Authentication.
func AuthHTTPassword(c cfg.Config) string {
	return c.Value(keyAuth).(string)
}

// AuthHTTPUser returns the user for HTTP Basic Authentication.
func AuthHTTPUser(c cfg.Config) string {
	return c.Value(keyAuth).(string)
}

// WithAuthentication sets the authentication method.
func WithAuthentication(c cfg.Config, a Auth) cfg.Config {
	return cfg.WithValue(c, keyAuth, a)
}

// WithAuthHTTPassword set password for HTTP Basic Authentication.
func WithAuthHTTPassword(c cfg.Config, v string) cfg.Config {
	return cfg.WithValue(c, keyAuthHTTPPassword, v)
}

// WithAuthHTTPUser set username for HTTP Basic Authentication.
func WithAuthHTTPUser(c cfg.Config, v string) cfg.Config {
	return cfg.WithValue(c, keyAuthHTTPUser, v)
}
