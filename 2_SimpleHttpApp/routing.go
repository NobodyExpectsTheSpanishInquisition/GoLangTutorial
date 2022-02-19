package main

import "net/http"

func registerRoutes() {
	RouteHelloWorld()
}

func RouteHelloWorld() {
	http.HandleFunc("/", helloWorld)
}
