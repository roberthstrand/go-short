package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	h "github.com/roberthstrand/go-short/internal/handlers"
)

var (
	rootUrl = os.Getenv("ROOT_URL")
)

func routes(app *fiber.App) {
	// If root of of the app is requested,
	// redirect to my website
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(rootUrl, fiber.StatusPermanentRedirect)
	})

	// Health check, used when the app is
	// deployed to Kubernetes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Add a new short URL
	app.Post("/", h.AddUrl)

	// If a short URL is requested,
	// redirect to the full URL
	app.Get("/:url", h.GetUrl)
}
