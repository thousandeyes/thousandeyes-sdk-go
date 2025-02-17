package core

import (
	"context"
	"net/http"
)

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
		UserAgent: "ThousandEyes Go SDK v3",
		Debug:     false,
		ServerURL: "https://api.thousandeyes.com/v7",
	}
	return cfg
}

func (c *Configuration) WithAuthToken(authToken string) *Configuration {
	c.AuthToken = authToken
	return c
}

func (c *Configuration) WithServerUrl(serverUrl string) *Configuration {
	c.ServerURL = serverUrl
	return c
}
