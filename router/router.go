package router

import (
	"glutara/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", middleware.MainHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/about", middleware.AboutHandler).Methods("GET", "OPTIONS")

	return router
}