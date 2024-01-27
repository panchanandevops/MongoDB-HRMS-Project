// internal/mongodb/mongo.go
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance represents the MongoDB client and database
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var Mg MongoInstance // Global instance of MongoInstance (Change 'mg' to 'Mg')

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

// Connect establishes a connection to the MongoDB database
func Connect() error {

	// Create a new MongoDB client with connection options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	// Set the global MongoInstance with the connected client and database (Change 'mg' to 'Mg')
	Mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}
