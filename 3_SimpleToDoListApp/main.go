package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

}

func startServer() {
	err := http.ListenAndServe(os.Getenv("APP_PORT"), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
