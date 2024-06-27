package errors

import (
	"fmt"
	"strings"
)

type SDKError struct {
	Code    int
	Message string
}

func (err *SDKError) Error() string {
	return fmt.Sprintf("[SDKError] code=%d, message=%s", err.Code, err.Message)
}

func NewSDKError(code int, errors []string) error {
	return &SDKError{
		Code:    code,
		Message: strings.Join(errors, ","),
	}
}

func (err *SDKError) GetCode() int {
	return err.Code
}

func (err *SDKError) GetMessage() string {
	return err.Message
}
