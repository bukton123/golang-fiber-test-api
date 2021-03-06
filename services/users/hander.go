package main

import (
	"api/pkg/api"
	"api/pkg/validator"
	"api/services/users/spec"
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Handler struct {
	validate    *validator.ValidateSchema
	userService spec.UserService
}

func (h *Handler) register(app *fiber.App) {
	app.Get("/users", h.main)
	app.Get("/users/err", h.err)
	app.Post("/users", h.create)
}

func (h *Handler) main(ctx *fiber.Ctx) error {
	result, err := h.userService.Find()
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"data": result})
}

func (h *Handler) err(ctx *fiber.Ctx) error {
	return errors.New("fail connection")
}

func (h *Handler) create(ctx *fiber.Ctx) error {
	body := new(spec.UserBody)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(api.HandlerErrorHttp(ServiceName, 1, "validate"))
	}

	if validate := h.validate.ValidateHandleError(body); validate != nil {
		return ctx.JSON(validate)
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"created": true, "data": body})
}
