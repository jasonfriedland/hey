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

// Package asap provides ASAP auth functionality.
package asap

import (
	"crypto"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"strings"

	"bitbucket.org/atlassian/go-asap"
)

// Config holds the JSON data in the asap config file.
type Config struct {
	Kid        string   `json:"kid"`
	Issuer     string   `json:"issuer"`
	Audience   []string `json:"audience"`
	PrivateKey string   `json:"privateKey"`
	Expiry     int64    `json:"expiry"`
}

// ParseConfig reads a config file and unmarshalls its contents.
func ParseConfig(config []byte) *Config {
	asapConfig := &Config{}
	if err := json.Unmarshal(config, &asapConfig); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return asapConfig
}

// GenerateAuthToken generates a unique Bearer token per request.
func (config *Config) GenerateAuthToken() (string, error) {
	privateKey, err := privateKeyFromBytes([]byte(config.PrivateKey))
	if err != nil {
		return "", err
	}

	asapToken := asap.NewASAP(config.Kid, config.Issuer, nil)
	token, err := asapToken.Sign(strings.Join(config.Audience, ","), privateKey)
	if err != nil {
		return "", err
	}
	authToken := "Bearer " + string(token)
	return authToken, nil
}

// privateKeyFromBytes lovingly borrowed from go-asap/keyprovider/util.go
func privateKeyFromBytes(privateKeyData []byte) (crypto.PrivateKey, error) {
	keyFromDataURL, err := x509.ParsePKCS8PrivateKey(privateKeyData)
	if err == nil {
		return keyFromDataURL, nil
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		return nil, errors.New("No valid PEM data found")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privateKey, err
	}

	return x509.ParseECPrivateKey(block.Bytes)
}
