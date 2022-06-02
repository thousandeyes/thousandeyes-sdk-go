package thousandeyes

import (
	"encoding/json"
	"reflect"
	"strings"
)

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

// booleanFieldsFromStruct receives a struct pointer and returns the
// JSON keys for fields tagged with "te:int-bool".
func booleanFieldsFromStruct(structPtr interface{}) map[string]bool {
	booleanFields := map[string]bool{}

	v := reflect.ValueOf(structPtr).Elem()
	t := reflect.TypeOf(v.Interface())

	for i := 0; i < t.NumField(); i++ {
		if tag, ok := t.Field(i).Tag.Lookup("te"); ok && tag == "int-bool" {
			jsonKey := strings.Split(t.Field(i).Tag.Get("json"), ",")[0]
			booleanFields[jsonKey] = true
		}
	}

	return booleanFields
}

// jsonBoolToInt is a helper routine that transforms ThousandEyes
// boolean fields to an int value. It is used by JSON marshalers.
func jsonBoolToInt(structPtr interface{}, data []byte) ([]byte, error) {
	booleanFields := booleanFieldsFromStruct(structPtr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, err
	}

	for jsonKey, jsonValue := range jsonMap {
		if booleanFields[jsonKey] {
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
func jsonIntToBool(structPtr interface{}, data []byte) ([]byte, error) {
	booleanFields := booleanFieldsFromStruct(structPtr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(data, &jsonMap); err != nil {
		return nil, err
	}

	for jsonKey, jsonValue := range jsonMap {
		if booleanFields[jsonKey] {
			jsonMap[jsonKey] = jsonValue.(float64) == 1
		}
	}

	return json.Marshal(jsonMap)
}
