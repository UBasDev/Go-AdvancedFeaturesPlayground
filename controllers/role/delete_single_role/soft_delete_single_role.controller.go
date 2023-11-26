package controllers

import (
	"net/http"

	role_services "example.com/goproject9/services/role"
	"gorm.io/gorm"
)

type NewSoftDeleteSingleRoleControllerModel struct {
	dbContext *gorm.DB
}

type INewSoftDeleteSingleRoleController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewSoftDeleteSingleRoleController(dbContext *gorm.DB) INewSoftDeleteSingleRoleController {
	return &NewSoftDeleteSingleRoleControllerModel{
		dbContext: dbContext,
	}
}

func (model NewSoftDeleteSingleRoleControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		role_services.SoftDeleteSingleRole(rw, rq, model.dbContext)
	}
}
