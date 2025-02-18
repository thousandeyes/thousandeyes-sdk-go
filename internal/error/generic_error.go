package error

import (
	"fmt"
	"reflect"
	"strings"
)

// GenericAPIError Provides access to the body, error and model on returned errors.
type GenericAPIError struct {
	Body         []byte
	ErrorMessage string
	Model        interface{}
}

func (g GenericAPIError) Error() string {
	return g.ErrorMessage
}

// FormatErrorMessage format error message using title and detail when model implements rfc7807
func FormatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}

// ReportError Prevent trying to import "fmt"
func ReportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
