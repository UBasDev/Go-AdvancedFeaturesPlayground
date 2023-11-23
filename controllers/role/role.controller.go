package controllers

import (
	"net/http"

	create_single_role "example.com/goproject9/controllers/role/create_single_role"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddRoleControllers(router *http.ServeMux, dbContext *gorm.DB) {
	create_single_role_handler := create_single_role.NewCreateSingleRoleController(dbContext)
	router.Handle("/api/role/createSingleRole", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_single_role_handler)))
}
