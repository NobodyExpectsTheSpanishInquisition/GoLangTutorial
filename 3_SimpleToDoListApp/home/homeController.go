package home

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome in TODO LIST APP")
	if err != nil {
		log.Fatal(err.Error())
	}
}
