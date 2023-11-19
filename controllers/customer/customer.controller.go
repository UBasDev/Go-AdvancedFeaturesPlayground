package controllers

import (
	"net/http"

	create_single_customer_controller "example.com/goproject9/controllers/customer/create_single_customer"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddCustomerControllers(router *http.ServeMux, appName string, dbContext *gorm.DB) {
	handler1 := create_single_customer_controller.NewCreateSingleCustomerController(appName, dbContext)
	router.Handle("/api/customer/createSingleCustomer", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(handler1)))
}
