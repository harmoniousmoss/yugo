package main

import (
	"fmt"
	"log"

	"gorps/handlers"
	"gorps/libs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load domain from .env file
	domain := libs.LoadEnv("DOMAIN")

	// Initialize Fiber app
	app := fiber.New()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Testing RPS tool is running!")
	})

	// Run the server
	go func() {
		fmt.Println("Server is running on port 8080")
		if err := app.Listen(":8080"); err != nil {
			log.Fatalf("Failed to start server: %s", err)
		}
	}()

	// Perform RPS test
	handlers.TestRPS(domain, "urls.txt")
}
