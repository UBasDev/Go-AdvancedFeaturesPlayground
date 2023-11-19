package main

import (
	"net/http"
	"net/http/pprof"

	customer_controllers "example.com/goproject9/controllers/customer"
	product_controllers "example.com/goproject9/controllers/product"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	if app.debug { //PPROF PROFILERS
		router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	}

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
