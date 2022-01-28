package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (h *Handler) register(app *fiber.App) {
	app.Get("/dapr/subscribe", h.subscribe)
	app.Post("/pubsub", h.pubsub)
}

func (h *Handler) subscribe(ctx *fiber.Ctx) error {
	return ctx.JSON([]fiber.Map{
		{
			"pubsubname": "pubsub",
			"topic":      "pubsub",
			"route":      "pubsub",
		},
	})
}

func (h *Handler) pubsub(ctx *fiber.Ctx) error {

	fmt.Printf("%+v\n", string(ctx.Body()))

	return ctx.JSON("OK")
}
