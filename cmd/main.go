package main

import (
	"fmt"
	"log"
	"os"
	"time"

	initializers "example.com/goproject9/initialize/database"
	"example.com/goproject9/migrations"

	"gorm.io/gorm"
)

type application struct {
	appName           string
	server            server
	isDebugActive     bool
	isProfilersActive bool
	errLog            *log.Logger
	infoLog           *log.Logger
	serverReadTimeOut time.Duration
	dbContext         *gorm.DB
}

type server struct {
	host string
	port string
	url  string
}

var dBContext *gorm.DB

func init() {
	//dsn := `host=localhost user=postgres password=admin dbname=go1 port=5432 sslmode=disable`
	dbUrl := "postgres://postgres:admin@localhost:5432/go1?sslmode=disable"
	dBContext = initializers.InitializeDbConnection(dbUrl)
	migrations := migrations.NewMigrationsPackage()
	go migrations.MigrateAllEntities(dBContext)
}

func main() {
	server := server{
		host: "127.0.0.1",
		port: "3001",
	}
	server.url = fmt.Sprintf("%s:%s", server.host, server.port)
	app := &application{
		appName:           "goproject9",
		server:            server,
		isDebugActive:     true,
		isProfilersActive: true,
		infoLog:           log.New(os.Stdout, "INFO\t", log.LstdFlags),
		errLog:            log.New(os.Stdout, "ERROR\t", log.LstdFlags),
		serverReadTimeOut: time.Duration(300) * time.Second,
		dbContext:         dBContext,
	}
	if err := app.listenAndServer(); err != nil {
		panic(err)
	}
	fmt.Println("test1")
}
