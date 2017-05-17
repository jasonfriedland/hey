// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package basic provides HTTP Basic auth functionality.
package basic

import "strings"

// Config holds the auth.Provider config.
type Config struct {
	Username string
	Password string
}

// Name is the unique name of this provider
func (config Config) Name() string {
	return "basic"
}

// GetProvider reads a config file and unmarshalls its contents.
func GetProvider(authConfig *string) Config {
	t := strings.Split(*authConfig, ":")
	return Config{Username: t[0], Password: t[1]}
}

// GenerateAuthToken generates a unique Bearer token per request.
func (config Config) GenerateAuthToken() (string, error) {
	return "", nil
}
