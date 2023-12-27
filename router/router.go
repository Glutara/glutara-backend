package router

import (
	"glutara/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", middleware.MainHandler).Methods("GET")
	router.HandleFunc("/api/posts-get", middleware.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts-post", middleware.AddPost).Methods("POST")

	return router
}