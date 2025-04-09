package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func Init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Environment file loaded")

	// Get the MongoDB connection URI from environment variables
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL environment variable is not set")
	}

	// Create a new MongoDB client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Define a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB cluster
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	MongoClient = client

}
