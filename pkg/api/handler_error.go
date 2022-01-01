package api

import (
	"api/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

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
		return ctx.Status(code).JSON(logging.HandlerErrorHttp(service, 0, text))
	}
}
