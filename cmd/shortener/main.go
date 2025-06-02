package main

import (
	"net/http"
)

const ServerURL = "localhost:8080"

func main() {
	InitMux()
	http.ListenAndServe(ServerURL, Mux)
}