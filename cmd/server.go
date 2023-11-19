package main

import (
	"fmt"
	"net/http"
)

type IListenAndServer interface {
	listenAndServer() error
}

func (app *application) listenAndServer() error {
	host := fmt.Sprintf("%s:%s", app.server.host, app.server.port)
	server := http.Server{
		Handler:     app.routes(),
		Addr:        host,
		ReadTimeout: app.serverReadTimeOut,
	}
	app.infoLog.Printf("Server is listening on: %s\n", host)
	return server.ListenAndServe()
}
