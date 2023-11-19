package controllers

import (
	"net/http"

	services "example.com/goproject9/services/customer"
	"gorm.io/gorm"
)

type NewCreateSingleCustomerControllerModel struct {
	appName   string
	dbContext *gorm.DB
}

type INewCreateSingleCustomerController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewCreateSingleCustomerController(appName string, dbContext *gorm.DB) INewCreateSingleCustomerController {
	return &NewCreateSingleCustomerControllerModel{
		appName:   appName,
		dbContext: dbContext,
	}
}

func (model NewCreateSingleCustomerControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		services.CreateSingleCustomer(rw, rq, model.dbContext)
	}
}
