package client

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

var (
	rateLimitResetHeaders = []string{
		"x-organization-rate-limit-reset",
		"x-instant-test-rate-limit-reset",
	}
	resetHeaderPattern = regexp.MustCompile(`^\s*[0-9]+\s*$`)
)

// Custom backoff function
func customBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}

	for _, header := range rateLimitResetHeaders {
		if resetTimeStr := resp.Header.Get(header); resetTimeStr != "" {
			if waitDuration := parseResetHeader(&resetTimeStr); waitDuration != nil {
				return *waitDuration
			}

			// Default backoff if no valid reset time is found
			return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
		}
	}

	// Default case for no rateLimitResetHeaders
	return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
}

// Function to parse the reset header
func parseResetHeader(value *string) *time.Duration {
	if value == nil || !resetHeaderPattern.MatchString(*value) {
		return nil
	}

	// Parse the header value to a Unix timestamp
	resetTimeUnix, err := strconv.ParseInt(*value, 10, 64)
	if err != nil {
		return nil
	}

	// Calculate the duration until the reset time
	resetTime := time.Unix(resetTimeUnix, 0)
	waitDuration := time.Until(resetTime)

	// Return nil if the reset time has already passed
	if waitDuration < 0 {
		zeroDuration := time.Duration(0)
		return &zeroDuration
	}
	return &waitDuration
}
