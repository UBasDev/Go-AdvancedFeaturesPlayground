package controllers

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type NewCreateSingleProductControllerModel struct {
	appName   string
	dbContext *gorm.DB
}

type INewCreateSingleProductController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewCreateSingleProductController(appName string, dbContext *gorm.DB) INewCreateSingleProductController {
	return &NewCreateSingleProductControllerModel{
		appName:   appName,
		dbContext: dbContext,
	}
}

func (model NewCreateSingleProductControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Product1 response has been sent")
	string1 := fmt.Sprintf("Product controller works! Application name: %s", model.appName)
	rw.Write([]byte(string1))
}
