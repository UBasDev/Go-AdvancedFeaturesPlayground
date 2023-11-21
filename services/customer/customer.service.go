package services

import (
	"encoding/json"
	"log"
	"net/http"

	customer_entity "example.com/goproject9/entities/customer"
	requests "example.com/goproject9/requests/customer"
	"gorm.io/gorm"
)

func CreateSingleCustomer(rw http.ResponseWriter, rq *http.Request, dbContext *gorm.DB) {
	// buffer := make([]byte, 1024)
	// byteCountRead, err := rq.Body.Read(buffer)
	// if err != nil && err != io.EOF {
	// 	log.Fatalf("Check your request body, unable to read: %s", err)
	// 	http.Error(rw, "Check your request body, unable to read", http.StatusBadRequest)
	// }
	// fmt.Printf("Okunan byte sayısı: %d\n", byteCountRead)
	// fmt.Printf("okunan request body datası: %s\n", string(buffer))
	// var customerToCreate = &requests.CreateSingleCustomerRequest{}
	// if err := json.Unmarshal(buffer, customerToCreate); err != nil {
	// 	log.Fatalf("Check your request body, unable to unmarshal: %s", err)
	// 	http.Error(rw, "Check your request body, unable to unmarshal", http.StatusBadRequest)
	// }
	var request_body requests.CreateSingleCustomerRequest
	if err := json.NewDecoder(rq.Body).Decode(&request_body); err != nil {
		log.Printf("Check your request body, unable to decode: %s", err.Error())
		http.Error(rw, "Check your request body, unable to decode", http.StatusBadRequest)
		return
	}
	customerToCreate, err := customer_entity.NewCustomerEntity(request_body.Firstname, request_body.Lastname, request_body.Email, request_body.Age, request_body.Gender)
	if err != nil {
		log.Printf("An error occured while creating a new customer instance: %s", err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	results := dbContext.Create(&customerToCreate)
	if results.Error != nil {
		log.Printf("Unable to write this customer object to the database: %s", results.Error.Error())
		http.Error(rw, "Unable to write this customer object to the database", http.StatusBadRequest)
		return
	}
	log.Printf("%d rows have been created in database", results.RowsAffected)
}
