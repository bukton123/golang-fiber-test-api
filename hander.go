package main

import (
	"api/pkg/logging"
	"api/spec"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	userService spec.UserService
}

func (h *Handler) register(app *fiber.App) {
	app.Get("/users", h.main)
	app.Get("/users/err", h.err)
	app.Post("/users", h.create)
}

func (h *Handler) main(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"data": h.userService.Find()})
}

func (h *Handler) err(ctx *fiber.Ctx) error {
	return errors.New("fail connection")
}

func (h *Handler) create(ctx *fiber.Ctx) error {
	body := new(spec.UserBody)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(logging.HandlerErrorHttp(ServiceName, 1, "validate"))
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"created": true, "data": body})
}
