package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutesAndMiddleware(R *mux.Router) {
	// Middleware
	R.Use(HandlePreflightMiddleware)
	R.Use(LoggingMiddleware)

	// Routes
	R.HandleFunc("/start", Start).Methods(http.MethodOptions, http.MethodPost)
	R.HandleFunc("/stop", Stop).Methods(http.MethodOptions, http.MethodGet)
	R.HandleFunc("/metrics", Metrics).Methods(http.MethodOptions, http.MethodGet)
	R.HandleFunc("/status", Status).Methods(http.MethodOptions, http.MethodGet)
	R.HandleFunc("/graph", Graph).Methods(http.MethodOptions, http.MethodGet)
}
