package controllers

import (
	"net/http"

	services "example.com/goproject9/services/customer"
	"gorm.io/gorm"
)

type NewCreateMultipleCustomersWithTransactionControllerModel struct {
	appName   string
	dbContext *gorm.DB
}

type INewCreateMultipleCustomersWithTransactionController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewCreateMultipleCustomersWithTransactionController(appName string, dbContext *gorm.DB) INewCreateMultipleCustomersWithTransactionController {
	return &NewCreateMultipleCustomersWithTransactionControllerModel{
		appName:   appName,
		dbContext: dbContext,
	}
}

func (model NewCreateMultipleCustomersWithTransactionControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		services.CreateMultipleCustomersWithTransactionManagement(rw, rq, model.dbContext)
		return
	}
}
