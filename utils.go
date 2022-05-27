package thousandeyes

import (
	"encoding/json"
)

var booleanFields = []string{
	"alertsEnabled",
	"bandwidthMeasurements",
	"bgpMeasurements",
	"enabled",
	"followRedirects",
	"includeCoveredPrefixes",
	"includeHeaders",
	"liveShare",
	"mtuMeasurements",
	"networkMeasurements",
	"recursiveQueries",
	"registerEnabled",
	"savedEvent",
	"throughputMeasurements",
	"useActiveFtp",
	"useExplicitFtps",
	"useNTLM",
	"usePublicBGP",
	"verifyCertificate",
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// elem checks that the first parameter is an element of the received
// slice.
func elem(str string, elems []string) bool {
	for _, e := range elems {
		if e == str {
			return true
		}
	}
	return false
}

// jsonBoolToInt is a helper routine that transforms ThousandEyes
// boolean fields to an int value. It is used by JSON marshalers.
func jsonBoolToInt(data []byte) ([]byte, error) {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, err
	}

	for jsonKey, jsonValue := range jsonMap {
		if elem(jsonKey, booleanFields) {
			if jsonValue.(bool) {
				jsonMap[jsonKey] = 1
			} else {
				jsonMap[jsonKey] = 0
			}
		}
	}

	return json.Marshal(jsonMap)
}

// jsonIntToBool is a helper routine that transforms ThousandEyes int
// fields to a boolean value. It is used by JSON unmarshalers.
func jsonIntToBool(data []byte) ([]byte, error) {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, err
	}

	for jsonKey, jsonValue := range jsonMap {
		if elem(jsonKey, booleanFields) {
			jsonMap[jsonKey] = jsonValue.(float64) == 1
		}
	}

	return json.Marshal(jsonMap)
}
