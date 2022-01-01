package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func (h *Handler) register(app *fiber.App) {
	app.Get("/users", h.main)
	app.Get("/users/err", h.err)
}

func (h *Handler) main(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"test": 123})
}

func (h *Handler) err(ctx *fiber.Ctx) error {
	return errors.New("fail connection")
}
