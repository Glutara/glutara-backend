package router

import (
	"glutara/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	// Auth Handler
	router.HandleFunc("/api/auth/register", middleware.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", middleware.Login).Methods("POST")
	
	// Reminder Handler
	router.HandleFunc("/api/{UserID}/reminders", middleware.GetReminders).Methods("GET")
	router.HandleFunc("/api/{UserID}/reminders", middleware.CreateReminder).Methods("POST")
	router.HandleFunc("/api/{UserID}/reminders/{ReminderID}", middleware.DeleteReminder).Methods("DELETE")
	router.HandleFunc("/api/{UserID}/reminders/{ReminderID}", middleware.UpdateReminder).Methods("PUT")

	// Sleep Log Handler
	router.HandleFunc("/api/{UserID}/sleeps", middleware.GetSleeps).Methods("GET")
	router.HandleFunc("/api/{UserID}/sleeps", middleware.CreateSleep).Methods("POST")
	router.HandleFunc("/api/{UserID}/sleeps/{SleepID}", middleware.DeleteSleep).Methods("DELETE")
	router.HandleFunc("/api/{UserID}/sleeps/{SleepID}", middleware.UpdateSleep).Methods("PUT")
	
	return router
}