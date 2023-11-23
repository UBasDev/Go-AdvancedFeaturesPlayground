package services

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/goproject9/entities"
	requests "example.com/goproject9/requests/customer"
	responses "example.com/goproject9/responses/customer"
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
		log.Printf("Check your request body, unable to decode: %s", err)
		http.Error(rw, "Check your request body, unable to decode", http.StatusBadRequest)
		return
	}
	roleToAssign := entities.RoleEntity{}
	resultsFromRole := dbContext.Model(&entities.RoleEntity{}).First(&roleToAssign, "id = ?", request_body.RoleId)
	if resultsFromRole.Error != nil {
		log.Printf("Unable to find this role with given id: %s", resultsFromRole.Error)
		http.Error(rw, "Unable to find this role with given id", http.StatusBadRequest)
		return
	}
	screensToAssign := []entities.ScreenEntity{}
	resultsFromScreens := dbContext.Model(&entities.ScreenEntity{}).Find(&screensToAssign, request_body.ScreenIds)
	if resultsFromScreens.Error != nil {
		log.Printf("Unable to find this screen with given id: %s", resultsFromScreens.Error)
		http.Error(rw, "Unable to find this screen with given id", http.StatusBadRequest)
		return
	}
	customerToCreate, err := entities.NewCustomerEntity(request_body.Firstname, request_body.Lastname, request_body.Email, request_body.Age, request_body.Gender)
	if err != nil {
		log.Printf("An error occured while creating a new customer instance: %s", err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	customerToCreate.RoleId1 = roleToAssign.Id
	customerToCreate.Screens1 = make([]*entities.ScreenEntity, len(screensToAssign))
	for i := range screensToAssign {
		customerToCreate.Screens1[i] = &screensToAssign[i]
	}
	resultsFromCustomerCreate := dbContext.Model(&entities.CustomerEntity{}).Create(&customerToCreate)
	if resultsFromCustomerCreate.Error != nil {
		log.Printf("Unable to write this customer object to the database: %s", resultsFromCustomerCreate.Error)
		http.Error(rw, "Unable to write this customer object to the database", http.StatusBadRequest)
		return
	}
	//log.Printf("%d rows have been created in database", resultsFromCustomerCreate.RowsAffected)
	resultsFromProfileCreate := dbContext.Model(&entities.ProfileEntity{}).Create(&entities.ProfileEntity{
		TokenCount:         request_body.TokenCount,
		BalanceIntegerPart: request_body.BalanceIntegerPart,
		BalanceDecimalPart: request_body.BalanceDecimalPart,
		CustomerId1:        customerToCreate.Id,
	})
	if resultsFromProfileCreate.Error != nil {
		log.Printf("Unable to write this profile object to the database: %s", resultsFromProfileCreate.Error)
		http.Error(rw, "Unable to write this profile object to the database", http.StatusBadRequest)
		return
	}
}

func CreateMultipleCustomersWithTransactionManagement(rw http.ResponseWriter, rq *http.Request, dbContext *gorm.DB) {
	var request_body requests.CreateSingleCustomerRequest
	if err := json.NewDecoder(rq.Body).Decode(&request_body); err != nil {
		log.Printf("Check your request body, unable to decode: %s", err)
		http.Error(rw, "Check your request body, unable to decode", http.StatusBadRequest)
		return
	}
	customerToCreate, err := entities.NewCustomerEntity(request_body.Firstname, request_body.Lastname, request_body.Email, request_body.Age, request_body.Gender)
	if err != nil {
		log.Printf("An error occured while creating a new customer instance: %s", err.Error())
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	transaction1 := dbContext.Begin()
	results := transaction1.Model(&entities.CustomerEntity{}).Create(&customerToCreate)
	if results.Error != nil {
		log.Printf("Unable to write this customer object to the database: %s", results.Error)
		http.Error(rw, "Unable to write this customer object to the database", http.StatusBadRequest)
		transaction1.Rollback()
		return
	}
	transaction1.SavePoint("savePoint1")
	results2 := transaction1.Model(&entities.CustomerEntity{}).Create(&customerToCreate)
	if results2.Error != nil {
		log.Printf("Unable to write this customer object to the database: %s", results.Error)
		transaction1.RollbackTo("savePoint1")
	}
	results3 := transaction1.Commit()
	if results3.Error != nil {
		log.Printf("Unable to write this customer object to the database: %s", results.Error)
		http.Error(rw, "Unable to write this customer object to the database", http.StatusBadRequest)
		transaction1.Rollback()
		return
	}
}

func GetAllCustomers(rw http.ResponseWriter, rq *http.Request, dbContext *gorm.DB) {
	response := []responses.GetAllCustomersResponse{}
	all_customers := []entities.CustomerEntity{}
	results := dbContext.Model(&entities.CustomerEntity{}).Find(&all_customers)
	if results.Error != nil {
		log.Printf("Unable to retrieve all customers from database: %s", results.Error)
		http.Error(rw, "Unable to retrieve all customers from database", http.StatusBadRequest)
		return
	}
	for _, current_customer := range all_customers {
		mappedCustomerToAddToResponse := responses.GetAllCustomersResponse{
			Id:        current_customer.Id,
			Firstname: current_customer.Firstname,
			Lastname:  current_customer.Lastname,
			Email:     current_customer.Email,
			Age:       current_customer.Age,
			Gender:    current_customer.Gender,
			CreatedAt: current_customer.CreatedAt,
			UpdatedAt: current_customer.UpdatedAt,
		}
		if current_customer.DeletedAt.Valid {
			mappedCustomerToAddToResponse.DeletedAt = current_customer.DeletedAt.Time
		}
		response = append(response, mappedCustomerToAddToResponse)
	}
	serializedCustomer, err := json.Marshal(response)
	if err != nil {
		log.Printf("Unable to serialize all customers from database: %s", results.Error)
		http.Error(rw, "Unable to serialize all customers from database", http.StatusBadRequest)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	if _, err = rw.Write(serializedCustomer); err != nil {
		log.Printf("Unable to serialize all customers from database: %s", results.Error)
		http.Error(rw, "Unable to serialize all customers from database", http.StatusBadRequest)
		return
	}
}
