package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()
	registerRoutes()
	startServer()
}

func startServer() {
	err := http.ListenAndServe(os.Getenv("APP_PORT"), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
}
