package main

import (
	"log"
	"net/http"
)

func startServer() {
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
