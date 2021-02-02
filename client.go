package thousandeyes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

const (
	apiEndpoint = "https://api.thousandeyes.com/v6"
)

// APILinks - List of APILink
type APILinks []APILink

// APILink - an api link
type APILink struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
}

type errorObject struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Limiter - Rate limiter interface
type Limiter interface {
	Wait()
}

//HTTPClient - an http client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientOptions - Thousandeyes client options for accountID, AuthToken & rate limiter
type ClientOptions struct {
	Limiter   Limiter
	AccountID string
	AuthToken string
}

// Client wraps http client
type Client struct {
	AuthToken      string
	AccountGroupID string
	APIEndpoint    string
	HTTPClient     http.Client
	Limiter        Limiter
}

// DefaultLimiter -  thousandeyes rate limit is 240 per minute
type DefaultLimiter struct{}

// Wait - Satisfying the Limiter interface and wait on 300ms to avoid TE 240 per minute default
func (l DefaultLimiter) Wait() {
	time.Sleep(time.Millisecond * 300)
}

// NewClient creates an API client
func NewClient(opts *ClientOptions) *Client {
	return &Client{
		AuthToken:      opts.AuthToken,
		AccountGroupID: opts.AuthToken,
		APIEndpoint:    apiEndpoint,
		HTTPClient: http.Client{
			Timeout: time.Second * 10,
		},
		Limiter: opts.Limiter,
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
	if !reflect.ValueOf(c.Limiter).IsNil() {
		c.Limiter.Wait()
	}
	endpoint := c.APIEndpoint + path + ".json"
	req, _ := http.NewRequest(method, endpoint, body)
	if c.AccountGroupID != "" {
		q := req.URL.Query()
		q.Add("aid", c.AccountGroupID)
		req.URL.RawQuery = q.Encode()
	}
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
