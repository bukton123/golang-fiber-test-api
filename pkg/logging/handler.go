package logging

import "fmt"

type (
	handlerError struct {
		Error handlerErrorMessage `json:"error"`
	}

	handlerErrorMessage struct {
		Type    string `json:"type"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func HandlerErrorHttp(service string, code int, message string) handlerError {
	return handlerError{
		Error: handlerErrorMessage{
			Type:    fmt.Sprintf("%sException", service),
			Code:    code,
			Message: message,
		},
	}
}
