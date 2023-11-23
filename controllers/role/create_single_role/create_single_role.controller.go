package controllers

import (
	"net/http"

	role_services "example.com/goproject9/services/role"
	"gorm.io/gorm"
)

type NewCreateSingleRoleControllerModel struct {
	dbContext *gorm.DB
}

type INewCreateSingleRoleController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewCreateSingleRoleController(dbContext *gorm.DB) INewCreateSingleRoleController {
	return &NewCreateSingleRoleControllerModel{
		dbContext: dbContext,
	}
}

func (model NewCreateSingleRoleControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		role_services.CreateSingleRole(rw, rq, model.dbContext)
	}
}
