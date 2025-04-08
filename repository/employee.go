package repository

import (
	"context"
	"fmt"

	"github.com/heshanu/gorest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) CreateEmployee(emp *models.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (r *EmployeeRepo) FindEmployeeById(empID string) (*models.Employee, error) {
	var emp models.Employee

	err := r.MongoCollection.FindOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}}).Decode(&emp)

	if err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]models.Employee, error) {
	result, err := r.MongoCollection.Find(context.Background(),
		bson.D{})

	if err != nil {
		return nil, err
	}

	var emps []models.Employee
	err = result.All(context.Background(), &emps)
	if err != nil {
		return nil, fmt.Errorf("results decode error")
	}
	return emps, nil
}

func (r *EmployeeRepo) UpdateEmployeeById(empID string, updateEmp *models.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
		bson.D{{Key: "$set", Value: updateEmp}},
	)

	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeById(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(
		context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
	)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}
