package main

import "github.com/gofiber/fiber/v2"

type Handler struct {}

func (h *Handler) register (app *fiber.App) {
	app.Get("/", h.main)
}

func (h *Handler) main (ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{ "test": 123 })
}
