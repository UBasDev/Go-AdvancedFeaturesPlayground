package controllers

import (
	"net/http"

	services "example.com/goproject9/services/customer"
	"gorm.io/gorm"
)

type NewUpdateSingleCustomerByRoleControllerModel struct {
	dbContext *gorm.DB
}

type INewUpdateSingleCustomerByRoleController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewUpdateSingleCustomerByRoleController(dbContext *gorm.DB) INewUpdateSingleCustomerByRoleController {
	return &NewUpdateSingleCustomerByRoleControllerModel{
		dbContext: dbContext,
	}
}

func (model NewUpdateSingleCustomerByRoleControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPut {
		services.UpdateSingleCustomerByRole(rw, rq, model.dbContext)
		return
	}
}
