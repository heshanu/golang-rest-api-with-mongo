package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/heshanu/gorest/models"
	"github.com/heshanu/gorest/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")

	res := &Response{}

	var emp models.Employee

	defer json.NewEncoder(w).Encode(res)

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid body", err)
		res.Error = "Invalid body"
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	insertedID, err := repo.CreateEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Invalid body", err)
		res.Error = "Invalid body"
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusAccepted)

	log.Println("employee successfully created", insertedID, emp)

}

func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	result, err := repo.FindAllEmployee()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusAccepted)

	log.Println("employee successfully fetched")

}

func (svc *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get emp with related ID
	empID := mux.Vars(r)["id"]
	log.Printf("employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	result, err := repo.FindEmployeeById(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusAccepted)

	log.Println("employee successfully fetched")

}

func (svc *EmployeeService) UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get emp with related ID
	empID := mux.Vars(r)["id"]
	log.Printf("employee id", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("empty employee Id")
		res.Error = "empty employee Id"
		return
	}

	var emp models.Employee

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.UpdateEmployeeById(empID, &emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusAccepted)

	log.Println("employee successfully fetched")

}

func (svc *EmployeeService) DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//get emp with related ID
	empID := mux.Vars(r)["id"]
	log.Printf("employee id", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	result, err := repo.DeleteEmployeeById(empID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusAccepted)

	log.Println("employee successfully deletes")

}
