package controllers

import (
	"net/http"

	create_multiple_customers_with_transaction_controller "example.com/goproject9/controllers/customer/create_multiple_customers_with_transaction"
	create_single_customer_controller "example.com/goproject9/controllers/customer/create_single_customer"
	get_all_customers_controller "example.com/goproject9/controllers/customer/get_all_customers"
	update_single_customer_by_role_controller "example.com/goproject9/controllers/customer/update_single_customer_by_role"
	update_single_customer_by_screen_controller "example.com/goproject9/controllers/customer/update_single_customer_by_screen"
	"example.com/goproject9/middlewares"
	"gorm.io/gorm"
)

const base_route string = "/api/customer"

func AddCustomerControllers(router *http.ServeMux, appName string, dbContext *gorm.DB) {
	create_single_customer_handler := create_single_customer_controller.NewCreateSingleCustomerController(appName, dbContext)
	create_multiple_customers_with_transactions_handler := create_multiple_customers_with_transaction_controller.NewCreateMultipleCustomersWithTransactionController(appName, dbContext)
	get_all_customers_handler := get_all_customers_controller.NewGetAllCustomersController(dbContext)
	update_single_customer_by_screen_handler := update_single_customer_by_screen_controller.NewUpdateSingleCustomerByScreenController(dbContext)
	update_single_customer_by_role_handler := update_single_customer_by_role_controller.NewUpdateSingleCustomerByRoleController(dbContext)
	router.Handle(base_route+"/createSingleCustomer", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_single_customer_handler)))
	router.Handle(base_route+"/createMultipleCustomersWithTransaction", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(create_multiple_customers_with_transactions_handler)))
	router.Handle(base_route+"/getAllCustomers", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(get_all_customers_handler)))
	router.Handle(base_route+"/updateSingleCustomerByScreen", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(update_single_customer_by_screen_handler)))
	router.Handle(base_route+"/updateSingleCustomerByRole", middlewares.CustomMiddleware1(middlewares.CustomMiddleware2(update_single_customer_by_role_handler)))
}
