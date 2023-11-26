package controllers

import (
	"net/http"

	role_services "example.com/goproject9/services/role"
	"gorm.io/gorm"
)

type NewHardDeleteSingleRoleControllerModel struct {
	dbContext *gorm.DB
}

type INewHardDeleteSingleRoleController interface {
	ServeHTTP(rw http.ResponseWriter, rq *http.Request)
}

func NewHardDeleteSingleRoleController(dbContext *gorm.DB) INewHardDeleteSingleRoleController {
	return &NewHardDeleteSingleRoleControllerModel{
		dbContext: dbContext,
	}
}

func (model NewHardDeleteSingleRoleControllerModel) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		role_services.HardDeleteSingleRole(rw, rq, model.dbContext)
	}
}
