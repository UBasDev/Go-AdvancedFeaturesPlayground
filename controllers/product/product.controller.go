package controllers

import (
	"net/http"

	create_single_product_controller "example.com/goproject9/controllers/product/create_single_product"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddProductControllers(router *http.ServeMux, appName string, dbContext *gorm.DB) {
	handler1 := create_single_product_controller.NewCreateSingleProductController(appName, dbContext)
	router.Handle("/api/product/createSingleProduct", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(handler1)))
}
