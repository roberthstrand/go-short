package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	db "github.com/roberthstrand/go-short/internal/backend/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUrl(c *fiber.Ctx) error {
	doc := bson.M{
		"fullUrl": c.FormValue("fullUrl"),
		"short":   c.FormValue("short"),
	}

	client, ctx, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("go-short").Collection("urls")
	//todo: check if short url already exists
	_, err = collection.InsertOne(ctx, doc)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString("Error inserting URL: " + err.Error())
	}
	return c.SendString("URL inserted")
}
