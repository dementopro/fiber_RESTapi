package mongo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/toc/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB represents the MongoDB utility.
type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	logger     *logrus.Logger
}

// NewMongoDB creates a new MongoDB utility instance.
func NewMongoDB(cfg config.MongoClient, logger *logrus.Logger) (*MongoDB, error) {
	// Get the MongoDB connection details from environment variables
	// Create a MongoDB connection string
	// connectionString := fmt.Sprintf("mongodb://%s:%s@%s", username, password, url)
	// connectionString := os.Getenv("MONGOURI")

	// Set up options for the MongoDB client
	clientOptions := options.Client().ApplyURI(EnvMongoURI())

	// TODO: Load additional options from an external file (e.g., config.yaml)

	// Create a new MongoDB client
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logger.WithError(err).Error("Failed to connect to MongoDB")
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.WithError(err).Error("Failed to ping MongoDB")
		return nil, err
	}

	// Set the database and collection
	database := client.Database(cfg.Database)
	collection := database.Collection(cfg.Collection)

	// logrus.Println("Connected to MongoDB!")
	logger.Info("Connected to MongoDB!")

	return &MongoDB{
		client:     client,
		database:   database,
		collection: collection,
		logger:     logger,
	}, nil
}

// Disconnect disconnects from MongoDB.
func (db *MongoDB) Disconnect() error {
	err := db.client.Disconnect(context.Background())
	if err != nil {
		db.logger.WithError(err).Error("Failed to disconnect from MongoDB")
		return err
	}
	db.logger.Info("Disconnected from MongoDB")
	return nil
}

// Create inserts a document into the collection.
func (db *MongoDB) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	result, err := db.collection.InsertOne(ctx, document)
	if err != nil {
		db.logger.WithError(err).Error("Failed to create document in MongoDB")
		return nil, err
	}
	return result, nil
}

// Read retrieves documents from the collection based on the provided filter.
func (db *MongoDB) Read(ctx context.Context, filter interface{}, result interface{}) ([]interface{}, error) {
	cur, err := db.collection.Find(ctx, filter)
	if err != nil {
		db.logger.WithError(err).Error("Failed to read documents from MongoDB")
		return nil, err
	}
	defer cur.Close(ctx)

	var results []interface{}

	for cur.Next(ctx) {
		err := cur.Decode(result)
		if err != nil {
			db.logger.WithError(err).Error("Failed to decode document from MongoDB")
			return nil, err
		}
		results = append(results, result)

	}
	return results, nil
}

// Read retrieves documents from the collection based on the provided filter.
func (db *MongoDB) ReadOne(ctx context.Context, id string, result interface{}) (interface{}, error) {
	err := db.collection.FindOne(ctx, bson.M{"tenant_id": id}).Decode(result)
	if err != nil {
		db.logger.WithError(err).Error("Failed to read document from MongoDB")
		return nil, err
	}

	return result, nil
}

// Update updates a document in the collection based on the provided filter.
func (db *MongoDB) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	result, err := db.collection.UpdateMany(ctx, filter, update)
	return result, err
}

// Update updates a document in the collection based on the provided filter.
func (db *MongoDB) UpdateOne(ctx context.Context, id string, update interface{}) (*mongo.UpdateResult, error) {
	result, err := db.collection.UpdateOne(ctx, bson.M{"tenant_id": id}, bson.M{"$set": update})
	if err != nil {
		db.logger.WithError(err).Error("Failed to update document in MongoDB")
		return nil, err
	}
	return result, nil
}

// Delete documents from the collection based on the provided filter.
func (db *MongoDB) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	result, err := db.collection.DeleteMany(ctx, filter)
	if err != nil {
		db.logger.WithError(err).Error("Failed to delete documents from MongoDB")
		return nil, err
	}
	return result, nil
}

// Delete deletes documents from the collection based on the provided filter.
func (db *MongoDB) DeleteOne(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	result, err := db.collection.DeleteOne(ctx, bson.M{"tenant_id": id})
	if err != nil {
		db.logger.WithError(err).Error("Failed to delete document from MongoDB")
		return nil, err
	}
	return result, nil
}
