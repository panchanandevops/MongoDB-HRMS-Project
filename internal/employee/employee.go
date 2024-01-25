// internal/employee/employee.go
package employee

import (
	"log"

	"github.com/panchanandevops/MongoDB-HRMS-Project/internal/mongodb"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Employee represents the structure of an employee
type Employee struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Salary float64            `json:"salary"`
	Age    float64            `json:"age"`
}

// GetAllEmployees retrieves all employees from the MongoDB database
func GetAllEmployees(c *fiber.Ctx) error {
	query := bson.D{}
	cursor, err := mongodb.Mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	var employees []Employee
	if err := cursor.All(c.Context(), &employees); err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

// CreateEmployee creates a new employee record in the MongoDB database
func CreateEmployee(c *fiber.Ctx) error {
	collection := mongodb.Mg.Db.Collection("employees")

	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		log.Println(err)
		return c.Status(400).SendString(err.Error())
	}

	employee.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(employee)
}

// UpdateEmployee updates an existing employee record in the MongoDB database
func UpdateEmployee(c *fiber.Ctx) error {
	idParam := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	employee := new(Employee)
	if err := c.BodyParser(employee); err != nil {
		log.Println(err)
		return c.Status(400).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	_, err = mongodb.Mg.Db.Collection("employees").UpdateOne(c.Context(), filter, update)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(employee)
}

// DeleteEmployee deletes an employee record from the MongoDB database
func DeleteEmployee(c *fiber.Ctx) error {
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	filter := bson.D{{Key: "_id", Value: employeeID}}
	result, err := mongodb.Mg.Db.Collection("employees").DeleteOne(c.Context(), filter)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		log.Println("No record deleted")
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("Record deleted")
}
