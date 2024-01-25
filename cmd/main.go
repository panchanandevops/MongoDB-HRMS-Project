package main

import (
	"log"

	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := routes.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	routes.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
