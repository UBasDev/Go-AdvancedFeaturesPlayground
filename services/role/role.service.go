package services

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/goproject9/entities"
	requests "example.com/goproject9/requests/role"
	"gorm.io/gorm"
)

func CreateSingleRole(rw http.ResponseWriter, rq *http.Request, dbContext *gorm.DB) {
	var request_body requests.CreateSingleRoleRequest
	if err := json.NewDecoder(rq.Body).Decode(&request_body); err != nil {
		log.Printf("Check your request body, unable to decode: %s", err)
		http.Error(rw, "Check your request body, unable to decode", http.StatusBadRequest)
		return
	}
	roleToCreate, _ := entities.NewRoleEntity(request_body.Key, request_body.Code, request_body.Value, request_body.Description)
	results := dbContext.Model(&entities.RoleEntity{}).Create(&roleToCreate)
	if results.Error != nil {
		log.Printf("Unable to write this role object to the database: %s", results.Error)
		http.Error(rw, "Unable to write this role object to the database", http.StatusBadRequest)
		return
	}
	log.Printf("%d rows have been created in database", results.RowsAffected)
}
