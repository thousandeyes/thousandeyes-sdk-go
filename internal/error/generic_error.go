package error

import (
	"fmt"
	"net/http"
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

// DecodeError decodes a modeled API error and preserves the response details.
func DecodeError(
	decode func(interface{}, []byte, string) error,
	model interface{},
	body []byte,
	response *http.Response,
) *GenericAPIError {
	apiErr := &GenericAPIError{
		Body:         body,
		ErrorMessage: response.Status,
	}
	if model == nil {
		return apiErr
	}

	if err := decode(model, body, response.Header.Get("Content-Type")); err != nil {
		apiErr.ErrorMessage = err.Error()
		return apiErr
	}

	apiErr.ErrorMessage = FormatErrorMessage(response.Status, string(body), model)
	apiErr.Model = model
	return apiErr
}

type titleProvider interface {
	GetTitleOk() (*string, bool)
}

type detailProvider interface {
	GetDetailOk() (*string, bool)
}

// FormatErrorMessage format error message using title and detail when model implements rfc7807
func FormatErrorMessage(status string, responseBody string, v interface{}) string {
	message := strings.TrimSpace(responseBody)
	title, hasTitle := problemTitle(v)
	detail, hasDetail := problemDetail(v)

	switch {
	case hasTitle && hasDetail:
		message = fmt.Sprintf("%s (%s)", title, detail)
	case hasTitle:
		message = title
	case hasDetail:
		message = detail
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, message))
}

func problemTitle(v interface{}) (string, bool) {
	provider, ok := v.(titleProvider)
	if !ok {
		return "", false
	}

	title, ok := provider.GetTitleOk()
	if !ok || title == nil || *title == "" {
		return "", false
	}
	return *title, true
}

func problemDetail(v interface{}) (string, bool) {
	provider, ok := v.(detailProvider)
	if !ok {
		return "", false
	}

	detail, ok := provider.GetDetailOk()
	if !ok || detail == nil || *detail == "" {
		return "", false
	}
	return *detail, true
}

// ReportError Prevent trying to import "fmt"
func ReportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
