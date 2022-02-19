package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		log.Fatal(err.Error())
	}
}
