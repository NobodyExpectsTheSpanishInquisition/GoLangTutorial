package main

import (
	"GoLangTutorial/3_SimpleToDoListApp/home"
	"github.com/gorilla/mux"
	"net/http"
)

func registerRoutes() {
	r := mux.Router{}

	homepage(&r)

	http.Handle("/", &r)
}

func homepage(router *mux.Router) {
	router.HandleFunc("/", home.Home).Methods("GET")
}
