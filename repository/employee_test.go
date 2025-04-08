package repository

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/heshanu/gorest/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://<>:test@cluster0.f6bnq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("Error while connecting to mongodb")
	}
	log.Println("mongodb connected successfully")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("ping failed")
	}
	return mongoTestClient

}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := NewMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	//dummy data
	emp1 := uuid.New().String()
	//emp2 := uuid.New().String()

	//connect to collection
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	employeeRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Inserted Employee one", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Heshan",
			Department: "CS",
			EmployeeID: "emp1",
		}

		result, err := employeeRepo.CreateEmployee(&emp)
		if err != nil {
			t.Fatal("Insert one operation failed", err)
		}

		t.Log("Insert 1 successful", result)
	})

	t.Run("Get Employees", func(t *testing.T) {
		result, err := employeeRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("Fetching employee operation failed", err)
		}

		t.Log("Fetched  successful", result)

	})

	t.Run("Update Employee by Id", func(t *testing.T) {
		emp := models.Employee{
			Name:       "Heshan",
			Department: "CS",
			EmployeeID: "emp00",
		}

		result, err := employeeRepo.UpdateEmployeeById(emp1, &emp)
		if err != nil {
			t.Fatal("Updated  operation failed", err)
		}

		t.Log("Updated successful", result)

	})

	t.Run("Get Employee by Id", func(t *testing.T) {
		result, err := employeeRepo.FindEmployeeById(emp1)
		if err != nil {
			t.Fatal("Fetching employee operation failed", err)
		}

		t.Log("Fetched  successful", result.Name)

	})

}
