package main

import (
	"log"
	"net/http"

	"github.com/Christoffer-lindow/portfolio/internal/config"
	"github.com/Christoffer-lindow/portfolio/internal/driver"
	"github.com/Christoffer-lindow/portfolio/internal/handlers"
)

var app config.AppConfig
var portNumber = ":8080"
var constants *Constants

func initHandlers(db *driver.DB) {
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
}

func connectDb() (*driver.DB, error) {
	constants = NewConstants()
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(constants.DatabaseUrl)
	if err != nil {
		return db, err
	}
	log.Println("Connected to database")
	return db, nil
}

func main() {
	app.InProduction = true
	db, err := connectDb()
	if err != nil {
		log.Fatal("Could not connect to DB")
	}
	defer db.SQL.Close()
	initHandlers(db)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("Starting server on port" + portNumber)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Could not start server")
	}

}
