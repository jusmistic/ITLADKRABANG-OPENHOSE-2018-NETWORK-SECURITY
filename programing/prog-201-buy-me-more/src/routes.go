package main

import (
	"github.com/gorilla/mux"
	"github.com/jusmistic/programming/prog-201-buy-me-more/src/handler"
)

// RegisterRoute register routes
func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/", handler.Home).Methods("GET")
	r.HandleFunc("/user/login", handler.Login).Methods("GET", "POST")
	r.HandleFunc("/user/register", handler.Register).Methods("GET", "POST")
	r.HandleFunc("/user/logout", handler.Logout).Methods("GET")
	// r.HandleFunc("/user/list", handler.ListUser).Methods("GET")
	r.HandleFunc("/item", handler.ViewItem).Methods("GET")
	r.HandleFunc("/item/", handler.ViewItem).Methods("GET")
	r.HandleFunc("/item/buy", handler.BuyItem).Methods("GET")
}
