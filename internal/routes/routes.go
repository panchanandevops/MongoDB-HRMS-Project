package routes

import (
	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/employee"
	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/mongodb"

	"github.com/gofiber/fiber/v2"
)

func Connect() error {
	return mongodb.Connect()
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/employee", employee.GetAllEmployees)
	app.Post("/employee", employee.CreateEmployee)
	app.Put("/employee/:id", employee.UpdateEmployee)
	app.Delete("/employee/:id", employee.DeleteEmployee)
}
