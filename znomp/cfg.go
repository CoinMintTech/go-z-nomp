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
	keyAuthFunc = configKey(iota)
	keyEndpoint = configKey(iota)
)

// AuthenticationFunc returns the authentication function.
func AuthenticationFunc(c cfg.Config) AuthFunc {
	return c.Value(keyAuthFunc).(AuthFunc)
}

// Endpoint returns the z-nomp endpoint.
func Endpoint(c cfg.Config) string {
	return c.Value(keyEndpoint).(string)
}

// WithAuthenticationFunc sets the authentication function.
func WithAuthenticationFunc(c cfg.Config, f AuthFunc) cfg.Config {
	return cfg.WithValue(c, keyAuthFunc, f)
}

// WithEndpoint sets the z-nomp endpoint.
func WithEndpoint(c cfg.Config, s string) cfg.Config {
	return cfg.WithValue(c, keyEndpoint, s)
}
