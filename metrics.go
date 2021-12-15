package thousandeyes

import (
	"time"
	"net/http"
	"fmt"
	"io"
)

// Metric - the network metrics that a given Agent collected as an outcome of a test 
type Metric struct{
	AvgLatency 	float32 `json:"avgLatency"`
	Loss 		float32 `json:"loss"`
	MaxLatency 	float32 `json:"maxLatency"`
	Jitter 		float32 `json:"jitter"`
	MinLatency 	float32 `json:"minLatency"`
	ServerIP 	string  `json:"serverIp"`
	AgentName   string  `json:"agentName"`
	CountryID   string  `json:"countryId"`
	AgentID     int     `json:"agentId"`
	RoundID     int     `json:"roundId"`
	PermaLink   string  `json:"permalink"`
}

// Net - the outcome of a test 
type Net struct {
    Test GenericTest `json:"test"`
	Metrics []Metric `json:"metrics"`
}

// Metrics - the response for a Network end2end metrics request
type Metrics struct {
	From string `json:"from"`
	To string `json:"to"`
	Net Net `json:"net"`
	Pages Page `json:"pages"`
}

// Page - the number of pages in a Test Data response
type Page struct {
	Current int `json:"current"`
}

func (c *Client) getPath(path string) (*http.Response, error) {
	return c.doPath("GET", path, nil, nil)
}

func (c *Client) doPath(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	if c.Limiter != nil {
		c.Limiter.Wait()
	}
	endpoint := c.APIEndpoint + path
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
	delay := setDelay(req, nil, time.Now())
	time.Sleep(delay)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Store reported rate limit status
	storeLimits(req, resp, time.Now())

	// If request was rate limited, back off and retry.
	// We shouldn't typically need to do this, because the above delays should
	// prevent us from hitting the limit, but there may be other users in an
	// org who might have triggered the limiting.
	if resp.StatusCode == 429 {
		delay := setDelay(req, resp, time.Now())
		time.Sleep(delay)
		resp, err = c.HTTPClient.Do(req)
	}

	return c.checkResponse(resp, err)
}

// GetNetMetrics - return the Network End2End Metrics of a test in a given time window
func (c *Client) GetNetMetrics(testID int, timeWindow string) (*Metrics, error) {
	resp, err := c.getPath(fmt.Sprintf("/net/metrics/%v.json?window=%s", testID, timeWindow))
	if err != nil {
		return &Metrics{}, err
	}
	var target Metrics
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("could not decode JSON response: %v", dErr)
	}
	metrics := target
	return &metrics, nil
}
