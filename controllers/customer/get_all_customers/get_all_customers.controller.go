package controllers

import (
	"net/http"

	services "example.com/goproject9/services/customer"
	"gorm.io/gorm"
)

type NewGetAllCustomersControllerModel struct {
	dbContext *gorm.DB
}

type INewGetAllCustomersController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewGetAllCustomersController(dbContext *gorm.DB) INewGetAllCustomersController {
	return &NewGetAllCustomersControllerModel{
		dbContext: dbContext,
	}
}

func (model NewGetAllCustomersControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodGet {
		services.GetAllCustomers(rw, rq, model.dbContext)
		return
	}
}
