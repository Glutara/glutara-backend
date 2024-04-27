package router

import (
	"github.com/gin-gonic/gin"

	"glutara/handlers"
	"glutara/middlewares"
)

// Router is exported and used in main.go
func ConfigureRouter(router *gin.Engine) {
	applyCorsMiddleware(router)
	router.GET("/api", handlers.MainHandler)

	// Auth Handler
	router.POST("/api/auth/register", handlers.Register)
	router.POST("/api/auth/login", handlers.Login)

	// Reminder Handler
	router.GET("/api/:UserID/reminders", handlers.GetReminders)
	router.POST("/api/:UserID/reminders", handlers.CreateReminder)
	router.DELETE("/api/:UserID/reminders/:ReminderID", handlers.DeleteReminder)
	router.PUT("/api/:UserID/reminders/:ReminderID", handlers.UpdateReminder)

	// Sleep Log Handler
	router.GET("/api/:UserID/sleeps", handlers.GetSleeps)
	router.POST("/api/:UserID/sleeps", handlers.CreateSleep)
	router.DELETE("/api/:UserID/sleeps/:SleepID", handlers.DeleteSleep)
	router.PUT("/api/:UserID/sleeps/:SleepID", handlers.UpdateSleep)

	// Exercise Log Handler
	router.GET("/api/:UserID/exercises", handlers.GetExercises)
	router.POST("/api/:UserID/exercises", handlers.CreateExercise)
	router.DELETE("/api/:UserID/exercises/:ExerciseID", handlers.DeleteExercise)
	router.PUT("/api/:UserID/exercises/:ExerciseID", handlers.UpdateExercise)

	// Meal Log Handler
	router.GET("/api/:UserID/meals", handlers.GetMeals)
	router.POST("/api/:UserID/meals", handlers.CreateMeal)
	router.DELETE("/api/:UserID/meals/:MealID", handlers.DeleteMeal)
	router.PUT("/api/:UserID/meals/:MealID", handlers.UpdateMeal)

	// Medication Log Handler
	router.GET("/api/:UserID/medications", handlers.GetMedications)
	router.POST("/api/:UserID/medications", handlers.CreateMedication)
	router.DELETE("/api/:UserID/medications/:MedicationID", handlers.DeleteMedication)
	router.PUT("/api/:UserID/medications/:MedicationID", handlers.UpdateMedication)

	// Blood Glucose Level Handler
	router.GET("/api/:UserID/glucoses", handlers.GetBloodGlucoseLevels)
	router.POST("/api/:UserID/glucoses", handlers.CreateBloodGlucoseLevel)

	// User Logs Handler
	router.GET("/api/:UserID/logs", handlers.GetLogs)
	
	// Scan Handler
	router.POST("/api/:UserID/scan", handlers.ScanFood)
}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
