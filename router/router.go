package router

import (
	"glutara/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api", middleware.MainHandler).Methods("GET", "OPTIONS")

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

	// Exercise Log Handler
	router.HandleFunc("/api/{UserID}/exercises", middleware.GetExercises).Methods("GET")
	router.HandleFunc("/api/{UserID}/exercises", middleware.CreateExercise).Methods("POST")
	router.HandleFunc("/api/{UserID}/exercises/{ExerciseID}", middleware.DeleteExercise).Methods("DELETE")
	router.HandleFunc("/api/{UserID}/exercises/{ExerciseID}", middleware.UpdateExercise).Methods("PUT")

	// Meal Log Handler
	router.HandleFunc("/api/{UserID}/meals", middleware.GetMeals).Methods("GET")
	router.HandleFunc("/api/{UserID}/meals", middleware.CreateMeal).Methods("POST")
	router.HandleFunc("/api/{UserID}/meals/{MealID}", middleware.DeleteMeal).Methods("DELETE")
	router.HandleFunc("/api/{UserID}/meals/{MealID}", middleware.UpdateMeal).Methods("PUT")

	// Medication Log Handler
	router.HandleFunc("/api/{UserID}/medications", middleware.GetMedications).Methods("GET")
	router.HandleFunc("/api/{UserID}/medications", middleware.CreateMedication).Methods("POST")
	router.HandleFunc("/api/{UserID}/medications/{MedicationID}", middleware.DeleteMedication).Methods("DELETE")
	router.HandleFunc("/api/{UserID}/medications/{MedicationID}", middleware.UpdateMedication).Methods("PUT")

	// Blood Glucose Level Handler
	router.HandleFunc("/api/{UserID}/glucoses", middleware.GetBloodGlucoseLevels).Methods("GET")
	router.HandleFunc("/api/{UserID}/glucoses", middleware.CreateBloodGlucoseLevel).Methods("POST")

	// Scan Handler
	router.HandleFunc("/api/{UserID}/scan", middleware.ScanFood).Methods("POST")

	return router
}