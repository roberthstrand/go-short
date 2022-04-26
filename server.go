package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	srvPort  = os.Getenv("SERVER_PORT")
	certFile = os.Getenv("CERT_FILE")
	certKey  = os.Getenv("CERT_KEY")
)

func main() {
	// Start a new instance of fiber,
	// with the logger middleware
	app := fiber.New()
	app.Use(logger.New())

	// Routes are set up in routes.go
	routes(app)

	// Start the server
	// Listen on port set in SERVER_PORT env var
	if srvPort != "" {
		log.Fatal(app.Listen(":" + srvPort))
	}
	// With TLS
	if certFile != "" && certKey != "" {
		log.Fatal(app.ListenTLS(":"+srvPort, "cert.pem", "key.pem"))
	}
	// If no port is set
	// listen on port 80
	if srvPort == "" {
		log.Fatal(app.Listen(":80"))
	}
}
