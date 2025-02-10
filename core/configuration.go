package core

import (
	"context"
	"net/http"
)

// Configuration stores the configuration of the API client
type Configuration struct {
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	Debug         bool              `json:"debug,omitempty"`
	ServerURL     string
	AuthToken     string
	HTTPClient    *http.Client
	Context       context.Context
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "ThousandEyes Go SDK v3",
		Debug:         false,
		ServerURL:     "https://api.thousandeyes.com/v7",
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) WithAuthToken(authToken string) *Configuration {
	c.AuthToken = authToken
	return c
}

func (c *Configuration) WithServerUrl(serverUrl string) *Configuration {
	c.ServerURL = serverUrl
	return c
}
