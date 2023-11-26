package controllers

import (
	"net/http"

	services "example.com/goproject9/services/customer"
	"gorm.io/gorm"
)

type NewUpdateSingleCustomerByScreenControllerModel struct {
	dbContext *gorm.DB
}

type INewUpdateSingleCustomerByScreenController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewUpdateSingleCustomerByScreenController(dbContext *gorm.DB) INewUpdateSingleCustomerByScreenController {
	return &NewUpdateSingleCustomerByScreenControllerModel{
		dbContext: dbContext,
	}
}

func (model NewUpdateSingleCustomerByScreenControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPut {
		services.UpdateSingleCustomerByScreen(rw, rq, model.dbContext)
		return
	}
}
