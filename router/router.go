package router

import (
	"glutara/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	// Post Handler
	router.HandleFunc("/api", middleware.MainHandler).Methods("GET")
	router.HandleFunc("/api/posts-get", middleware.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts-post", middleware.AddPost).Methods("POST")

	// Auth Handler
	router.HandleFunc("/api/auth/register", middleware.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", middleware.Login).Methods("POST")
	
	// Reminder Handler
	router.HandleFunc("/api/{UserID}/reminders", middleware.GetReminders).Methods("GET")
	router.HandleFunc("/api/{UserID}/reminders", middleware.CreateReminder).Methods("POST")
	router.HandleFunc("/api/{UserID}/reminders/{ReminderID}", middleware.DeleteReminder).Methods("DELETE")
	
	return router
}