package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	db "github.com/roberthstrand/go-short/internal/backend/mongodb"
)

func GetUrl(c *fiber.Ctx) error {
	url := c.Params("url")

	if os.Getenv("MONGODB_URI") != "" {
		client, ctx, err := db.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		collection := client.Database("go-short").Collection("urls").FindOne(ctx, map[string]string{"short": url})
		if collection.Err() != nil {
			return c.Status(fiber.StatusNotFound).SendString("URL not found")
		}

		var result map[string]string
		collection.Decode(&result)
		return c.Redirect(result["fullUrl"], fiber.StatusPermanentRedirect)
	}

	return c.Status(fiber.ErrBadRequest.Code).SendString("No backend set")
}
