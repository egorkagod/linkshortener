package main

import (
	"net/http"
	"github.com/egorkagod/linkshortener/internal/app"
)

var Mux = http.NewServeMux()

func InitMux() {
	Mux.HandleFunc("/", app.LinkHandler)
}