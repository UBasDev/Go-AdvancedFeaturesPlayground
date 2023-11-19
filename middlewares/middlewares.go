package middlewares

import (
	"log"
	"net/http"
)

func CustomMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

func CustomMiddleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareTwo")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}
