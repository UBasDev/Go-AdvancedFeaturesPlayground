package controllers

import (
	"net/http"

	create_multiple_customers_with_transaction_controller "example.com/goproject9/controllers/customer/create_multiple_customers_with_transaction"
	create_single_customer_controller "example.com/goproject9/controllers/customer/create_single_customer"
	get_all_customers_controller "example.com/goproject9/controllers/customer/get_all_customers"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

func AddCustomerControllers(router *http.ServeMux, appName string, dbContext *gorm.DB) {
	create_single_customer_handler := create_single_customer_controller.NewCreateSingleCustomerController(appName, dbContext)
	create_multiple_customers_with_transactions_handler := create_multiple_customers_with_transaction_controller.NewCreateMultipleCustomersWithTransactionController(appName, dbContext)
	get_all_customers_handler := get_all_customers_controller.NewGetAllCustomersController(dbContext)

	router.Handle("/api/customer/createSingleCustomer", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_single_customer_handler)))
	router.Handle("/api/customer/createMultipleCustomersWithTransaction", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_multiple_customers_with_transactions_handler)))
	router.Handle("/api/customer/getAllCustomers", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(get_all_customers_handler)))
}
