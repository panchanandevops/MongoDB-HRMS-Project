// routes/routes.go
package routes

import (
	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/employee"
	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/mongodb"

	"github.com/gofiber/fiber/v2"
)

// Connect initializes the MongoDB connection
func Connect() error {
	return mongodb.Connect()
}

// RegisterRoutes defines and registers all routes for the application
func RegisterRoutes(app *fiber.App) {

	// Define routes for employee management
	app.Get("/employee", employee.GetAllEmployees)
	app.Post("/employee", employee.CreateEmployee)
	app.Put("/employee/:id", employee.UpdateEmployee)
	app.Delete("/employee/:id", employee.DeleteEmployee)
}
