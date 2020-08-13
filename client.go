package thousandeyes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	apiEndpoint = "https://api.thousandeyes.com/v6"
)

var orgRate RateLimit
var instantTestRate RateLimit

// RateLimit contains data representing rate limit headers returned in
// ThousandEyes API responses.  int64 everywhere for ease of interacting
// with time values.
type RateLimit struct {
	Limit              int64
	Remaining          int64
	Reset              int64
	LastRemaining      int64
	LastTime           time.Time
	ConcurrentMessages int64
}

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

//HTTPClient - an http client
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client wraps http client
type Client struct {
	AuthToken      string
	AccountGroupID string
	APIEndpoint    string
	HTTPClient     http.Client
}

// NewClient creates an API client
func NewClient(authToken string, accountGroupID string) *Client {
	return &Client{
		AuthToken:      authToken,
		AccountGroupID: accountGroupID,
		APIEndpoint:    apiEndpoint,
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

	// Perform any delays required by previously observed rate headers
	delay := setDelay(req, nil)
	time.Sleep(delay)

	resp, err := c.HTTPClient.Do(req)

	// Store reported rate limit status
	storeLimits(req, resp)

	// If request was rate limited, back off and retry.
	// We shouldn't typically need to do this, because the above delays should
	// prevent us from hitting the limit, but there may be other users in an
	// org who might have triggered the limiting.
	if resp.StatusCode == 429 {
		delay := setDelay(req, resp)
		time.Sleep(delay)
		resp, err = c.HTTPClient.Do(req)
	}

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

// setDelay determines the pause time needed to prevent invoking rate limiting
func setDelay(req *http.Request, resp *http.Response) time.Duration {
	// Choose which rate limit applies
	var delay time.Duration
	var rate RateLimit
	now := time.Now()
	instantTest := isInstantTest(req)
	if instantTest {
		rate = instantTestRate
	} else {
		rate = orgRate
	}

	// If the limit is 0, this is either our first request or we are not receiving
	// rate limit data in the headers
	if rate.Limit == 0 {
		return 0
	}

	// If this is the first time we've sent this particular request and we
	// aren't at the end of our remaining requests for the period...
	if resp == nil && rate.Remaining > 1 {
		baseDelay := 1.0 / float64(rate.Limit) * float64(time.Minute.Nanoseconds())
		// Requests may not be sent synchroniously, so we need to calculate the ratio of
		// actual time elapsed to expected delay per message.
		sinceLast := float64(now.Sub(rate.LastTime).Nanoseconds())
		// The rate limit is per minute, so if there was a zero response time
		// then the ideal delay would be the one minute divided by the rate.
		// To account for potential other users, we will multiply by the
		// difference between the remaining count and our last seen remaining
		// count.
		delta := rate.LastRemaining - rate.Remaining
		if delta < 1 {
			delta = 1
		}

		// It's possible that these calls could be made concurrently, in which
		// case the pacing delay would effectively be divided by the batch size.
		// To account for this, we compare the number of messages issued in a
		// window smaller than the base delay and increase accordingly.
		if sinceLast < baseDelay {
			rate.ConcurrentMessages++
			delta += rate.ConcurrentMessages
		} else {
			rate.ConcurrentMessages = 0
		}
		delay = time.Duration(baseDelay * float64(delta))
		log.Printf("[INFO] %v of %v requests / min remain.  Sleeping %v to prevent rate limiting.",
			rate.Remaining, rate.Limit, delay)
	} else {
		// else calculate delay until resume time.
		// Assume our clock is roughly in sync with the clock setting the resume time.
		delay = time.Duration((rate.Reset - now.Unix() + 1) * time.Second.Nanoseconds())
		// ThousandEyes rates reset within one minute (but not guaranteed).
		// If we exceed a minute wait time, something may be wrong.
		if delay > time.Minute {
			delay = time.Minute
		}
		log.Printf("[INFO] Rate Limited: Sleeping %v before resubmitting\n", delay)
	}
	if instantTest {
		instantTestRate.LastTime = now
		instantTestRate.ConcurrentMessages = rate.ConcurrentMessages
	} else {
		orgRate.LastTime = now
		orgRate.ConcurrentMessages = rate.ConcurrentMessages
	}
	return delay
}

// storeLimits assigns the global variables to track current rate limit data
func storeLimits(req *http.Request, resp *http.Response) {
	// We discard errors, because an error or blank result also return 0
	if v := resp.Header.Get("X-Organization-Rate-Limit-Limit"); v != "" {
		orgRate.Limit, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("X-Organization-Rate-Limit-Remaining"); v != "" {
		orgRate.Remaining, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("X-Organization-Rate-Limit-Reset"); v != "" {
		orgRate.Reset, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Limit"); v != "" {
		instantTestRate.Limit, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Remaining"); v != "" {
		instantTestRate.Remaining, _ = strconv.ParseInt(v, 10, 64)
	}
	if v := resp.Header.Get("X-Instant-Test-Rate-Limit-Reset"); v != "" {
		instantTestRate.Reset, _ = strconv.ParseInt(v, 10, 64)
	}

	if isInstantTest(req) {
		instantTestRate.LastTime = time.Now()
	} else {
		orgRate.LastTime = time.Now()
	}
}

func isInstantTest(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/v6/instant") == true || strings.HasPrefix(req.URL.Path, "/v6/endpoint-instant")
}
