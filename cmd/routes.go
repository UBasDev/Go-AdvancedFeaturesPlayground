package main

import (
	"net/http"

	customer_controllers "example.com/goproject9/controllers/customer"
	product_controllers "example.com/goproject9/controllers/product"
)

func (app *application) routes() http.Handler {

	router := http.NewServeMux()
	product_controllers.AddProductControllers(router, app.appName, app.dbContext)
	customer_controllers.AddCustomerControllers(router, app.appName, app.dbContext)
	router.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) { //Diğer routelarla eşleşmeyen bütün requestleri kabul eder. Genellikle 404 responseları için kullanılır.
		if rq.URL.Path != "/" {
			http.NotFound(rw, rq)
			return
		}
	})
	return router
}
