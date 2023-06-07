package mongo

import (
	"context"
	"fmt"
	"os"

	"github.com/toc/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents the MongoDB utility.
type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// NewMongoDB creates a new MongoDB utility instance.
func NewMongoDB(cfg config.MongoClient) (*MongoDB, error) {
	// Get the MongoDB connection details from environment variables
	// Create a MongoDB connection string
	// connectionString := fmt.Sprintf("mongodb://%s:%s@%s", username, password, url)
	connectionString := os.Getenv("MONGODB_URL")

	// Set up options for the MongoDB client
	clientOptions := options.Client().ApplyURI(connectionString)

	// TODO: Load additional options from an external file (e.g., config.yaml)

	// Create a new MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// Set the database and collection
	database := client.Database(cfg.Database)
	collection := database.Collection(cfg.Collection)

	fmt.Println("Connected to MongoDB!")

	return &MongoDB{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

// Disconnect disconnects from MongoDB.
func (db *MongoDB) Disconnect() error {
	err := db.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("Disconnected from MongoDB!")
	return nil
}

// Create inserts a document into the collection.
func (db *MongoDB) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := db.collection.InsertOne(ctx, document)
	return result, err
}

// Read retrieves documents from the collection based on the provided filter.
func (db *MongoDB) Read(ctx context.Context, filter interface{}, result interface{}) ([]interface{}, error) {
	cur, err := db.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var results []interface{}
	for cur.Next(context.TODO()) {
		err := cur.Decode(result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

// Update updates a document in the collection based on the provided filter.
func (db *MongoDB) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	result, err := db.collection.UpdateMany(ctx, filter, update)
	return result, err
}

// Delete deletes documents from the collection based on the provided filter.
func (db *MongoDB) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := db.collection.DeleteMany(ctx, filter)
	return result, err
}
