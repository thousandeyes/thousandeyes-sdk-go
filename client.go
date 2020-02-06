package thousandeyes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiEndpoint = "https://api.thousandeyes.com/v6"
)

type ApiLink struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
}

type errorObject struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client wraps http client
type Client struct {
	AuthToken   string
	ApiEndpoint string
	HTTPClient  http.Client
}

// NewClient creates an API client
func NewClient(authToken string) *Client {
	return &Client{
		AuthToken:   authToken,
		ApiEndpoint: apiEndpoint,
		HTTPClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) put(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return c.do("PUT", path, bytes.NewBuffer(data), headers)
	}
	return c.do("PUT", path, nil, headers)
}

func (c *Client) post(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), headers)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	endpoint := c.ApiEndpoint + path + ".json"
	req, _ := http.NewRequest(method, endpoint, body)
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Set("content-type", "application/json")
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	resp, err := c.HTTPClient.Do(req)
	return c.checkResponse(resp, err)
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("Error calling the API endpoint: %v", err)
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		var eo *errorObject
		var getErr error
		if eo, getErr = c.getErrorFromResponse(resp); getErr != nil {
			return resp, fmt.Errorf("Response did not contain formatted error: %s. HTTP response code: %v. Raw response: %+v", getErr, resp.StatusCode, resp)
		}
		return resp, fmt.Errorf("Failed call API endpoint. HTTP response code: %v. Error: %v", resp.StatusCode, eo)
	}
	return resp, nil
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*errorObject, error) {
	var result errorObject
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", err)
	}
	return &result, nil
}
