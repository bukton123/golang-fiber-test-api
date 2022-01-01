package api

import (
	"api/pkg/logging"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	HandlerError struct {
		Error handlerErrorMessage `json:"error"`
	}

	handlerErrorMessage struct {
		Type    string `json:"type"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func HandlerErrorHttp(service string, code int, message string) HandlerError {
	return HandlerError{
		Error: handlerErrorMessage{
			Type:    fmt.Sprintf("%sException", service),
			Code:    code,
			Message: message,
		},
	}
}

func ErrorHandlerFiber(service string, version string, dev bool) fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		text := http.StatusText(http.StatusInternalServerError)
		if dev {
			text = err.Error()
		}

		logging.Error(service, version, err.Error())
		return ctx.Status(code).JSON(HandlerErrorHttp(service, 0, text))
	}
}
