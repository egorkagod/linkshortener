package main

import (
	"net/http"
)

const SERVER_URL = "localhost:8080"

func main() {
	InitMux()
	http.ListenAndServe(SERVER_URL, Mux)
}