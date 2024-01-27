package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/routes"
)

func main() {

	// Connect to MongoDB database
	if err := routes.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber app instance
	app := fiber.New()

	// Register routes with the Fiber app
	routes.RegisterRoutes(app)

	// Start the Fiber app on port 3000
	log.Fatal(app.Listen(":3000"))
}
