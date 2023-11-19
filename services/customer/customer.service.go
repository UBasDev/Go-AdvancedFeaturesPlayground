package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	requests "example.com/goproject9/requests/customer"
	"gorm.io/gorm"
)

func CreateSingleCustomer(rw http.ResponseWriter, rq *http.Request, dbContext *gorm.DB) {
	buffer := make([]byte, 1024)
	byteCountRead, err := rq.Body.Read(buffer)
	if err != nil && err != io.EOF {
		log.Fatalf("Check your request body, unable to read: %s", err)
		http.Error(rw, "Check your request body, unable to read", http.StatusBadRequest)
	}
	fmt.Printf("Okunan byte sayısı: %d\n", byteCountRead)
	fmt.Printf("okunan request body datası: %s\n", string(buffer))
	var customerToCreate = &requests.CreateSingleCustomerRequest{}
	err = json.Unmarshal(buffer, customerToCreate)
	if err != nil {
		log.Fatalf("Check your request body, unable to unmarshal: %s", err)
		http.Error(rw, "Check your request body, unable to unmarshal", http.StatusBadRequest)
	}
}
