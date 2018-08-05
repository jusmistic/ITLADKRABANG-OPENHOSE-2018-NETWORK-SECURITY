package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jusmistic/programming/prog-201-buy-me-more/src/database"
)

// App struct
type App struct {
	Router *mux.Router
}

// Initialize you app
func (a *App) Initialize() {
	// Defind db uri
	// Todo replace to use env or yaml ?
	dbURI := ""

	// Connect to a database
	log.Println("Connecting to database...")
	database.Connect(dbURI)
}

// Run you app
func (a *App) Run(serv string) {
	log.Println("Register a route")
	r := mux.NewRouter()
	RegisterRoute(r)

	// Close database
	defer database.Close()
	log.Println("Run server")
	log.Fatal(http.ListenAndServe(serv, r))
}
