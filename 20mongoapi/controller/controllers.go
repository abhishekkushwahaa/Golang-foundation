package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/abhishekkushwahaa/mongoapi/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Retrieve MongoDB connection string from environment variable
	connectionString := os.Getenv("MONGO_URL")
	if connectionString == "" {
		log.Fatal("MONGO_URL environment variable not set")
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())
	fmt.Println("Connected to MongoDB!")

	// Initialize collection handle
	collection = client.Database(dbName).Collection(colName)

}

// MongoDB Helper Functions

// InsertOne helper function to insert one document into MongoDB

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal("Error inserting movie:", err)
	}
	fmt.Println("Inserted movie with ID:", inserted.InsertedID)
}