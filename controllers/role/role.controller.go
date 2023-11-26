package controllers

import (
	"net/http"

	create_single_role "example.com/goproject9/controllers/role/create_single_role"
	delete_single_role "example.com/goproject9/controllers/role/delete_single_role"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddRoleControllers(router *http.ServeMux, dbContext *gorm.DB) {
	create_single_role_handler := create_single_role.NewCreateSingleRoleController(dbContext)
	soft_delete_single_role_handler := delete_single_role.NewSoftDeleteSingleRoleController(dbContext)
	hard_delete_single_role_handler := delete_single_role.NewHardDeleteSingleRoleController(dbContext)
	router.Handle("/api/role/createSingleRole", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_single_role_handler)))
	router.Handle("/api/role/softDeleteSingleRole", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(soft_delete_single_role_handler)))
	router.Handle("/api/role/hardDeleteSingleRole", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(hard_delete_single_role_handler)))
}
