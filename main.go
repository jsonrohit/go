// main.go
package main

import (
	"fiber/src/config"
	"fiber/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Create Fiber app with configuration
	app := fiber.New(fiber.Config{
		AppName:      "Fiber CRUD API",
		ErrorHandler: config.ErrorHandler,
	})

	// Setup routes
	routes.SetupRoutes(app)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	// Start server
	log.Printf("🚀 Server starting on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
