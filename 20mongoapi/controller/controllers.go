package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/abhishekkushwahaa/mongoapi/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// UpdateOne helper function to update one document in MongoDB

func updateOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"title": "The Matrix Reloaded"}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal("Error updating movie:", err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

// DeleteOne helper function to delete one document from MongoDB

func deleteOneMovie(movieId string){
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal("Error deleting movie:", err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)
}

// DeleteAll helper function to delete multiple documents from MongoDB

func deleteAllMovies() int64 {
	filter := bson.M{}
	result, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal("Error deleting movies:", err)
	}

	fmt.Println("Number of movie deleted:", result.DeletedCount)
	return result.DeletedCount
}

// Get All Movies helper function to retrieve all documents from MongoDB

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal("Error finding movies:", err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal("Error decoding movie:", err)
		}

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}