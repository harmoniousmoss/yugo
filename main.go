package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Testing RPS")
	})

	// Start the server (default: port 8080)
	app.Listen(":8080")
}
