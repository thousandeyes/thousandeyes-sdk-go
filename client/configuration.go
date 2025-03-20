package client

import (
	"context"
	"fmt"
	"net/http"
)

// This sdkVersion will be updated by release.yaml
// Do not modify this line manually
var sdkVersion = "v3"

func SDKVersion() string {
	return sdkVersion
}

// Configuration stores the configuration of the API client
type Configuration struct {
	UserAgent  string `json:"userAgent,omitempty"`
	Debug      bool   `json:"debug,omitempty"`
	ServerURL  string
	AuthToken  string
	HTTPClient *http.Client
	Context    context.Context
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		Debug:     false,
		ServerURL: "https://api.thousandeyes.com/v7",
	}
	return cfg
}

func (c *Configuration) BuildUserAgent() string {
	var sdkUserAgent = fmt.Sprintf("ThousandEyesSDK-Go/%s", SDKVersion())
	if c.UserAgent == "" {
		return sdkUserAgent
	}
	return sdkUserAgent + " " + c.UserAgent
}

func (c *Configuration) WithAuthToken(authToken string) *Configuration {
	c.AuthToken = authToken
	return c
}

func (c *Configuration) WithServerUrl(serverUrl string) *Configuration {
	c.ServerURL = serverUrl
	return c
}
