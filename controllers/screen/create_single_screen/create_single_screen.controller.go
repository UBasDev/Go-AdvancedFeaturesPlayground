package controllers

import (
	"net/http"

	services "example.com/goproject9/services/screen"
	"gorm.io/gorm"
)

type NewCreateSingleScreenControllerModel struct {
	dbContext *gorm.DB
}

type INewCreateSingleScreenController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewCreateSingleScreenController(dbContext *gorm.DB) INewCreateSingleScreenController {
	return &NewCreateSingleScreenControllerModel{
		dbContext: dbContext,
	}
}

func (model NewCreateSingleScreenControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		services.CreateSingleScreen(rw, rq, model.dbContext)
	}
}
