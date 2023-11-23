package controllers

import (
	"net/http"

	create_single_screen_controller "example.com/goproject9/controllers/screen/create_single_screen"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddScreenControllers(router *http.ServeMux, dbContext *gorm.DB) {
	create_single_screen_handler := create_single_screen_controller.NewCreateSingleScreenController(dbContext)
	router.Handle("/api/screen/createSingleScreen", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_single_screen_handler)))
}
