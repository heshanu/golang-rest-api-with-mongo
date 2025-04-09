package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/heshanu/gorest/config"
	"github.com/heshanu/gorest/usecase"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables first, before any other initialization
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize config (which presumably uses the env vars)
	config.Init()
	defer config.MongoClient.Disconnect(context.Background())

	collection := config.MongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// Create employee service using dependency injection
	empService := usecase.EmployeeService{MongoCollection: collection}

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/employee", empService.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", empService.GetAllEmployee).Methods("GET")
	r.HandleFunc("/employee/{id}", empService.GetEmployeeById).Methods("GET")
	r.HandleFunc("/employee/{id}", empService.UpdateEmployeeById).Methods("PUT")
	r.HandleFunc("/employee/{id}", empService.DeleteEmployeeById).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Println("PORT not set, defaulting to 8081")
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("running..."))
	log.Println("server is running...")
}
