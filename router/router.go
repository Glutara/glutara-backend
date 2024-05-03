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

	// Protected Endpoints
	protected := router.Group("")
	protected.Use(middlewares.AuthMiddleware())

	// Reminder Handler
	protected.GET("/api/:UserID/reminders", handlers.GetReminders)
	protected.POST("/api/:UserID/reminders", handlers.CreateReminder)
	protected.DELETE("/api/:UserID/reminders/:ReminderID", handlers.DeleteReminder)
	protected.PUT("/api/:UserID/reminders/:ReminderID", handlers.UpdateReminder)

	// Sleep Log Handler
	protected.GET("/api/:UserID/sleeps", handlers.GetSleeps)
	protected.POST("/api/:UserID/sleeps", handlers.CreateSleep)
	protected.DELETE("/api/:UserID/sleeps/:SleepID", handlers.DeleteSleep)
	protected.PUT("/api/:UserID/sleeps/:SleepID", handlers.UpdateSleep)

	// Exercise Log Handler
	protected.GET("/api/:UserID/exercises", handlers.GetExercises)
	protected.POST("/api/:UserID/exercises", handlers.CreateExercise)
	protected.DELETE("/api/:UserID/exercises/:ExerciseID", handlers.DeleteExercise)
	protected.PUT("/api/:UserID/exercises/:ExerciseID", handlers.UpdateExercise)

	// Meal Log Handler
	protected.GET("/api/:UserID/meals", handlers.GetMeals)
	protected.POST("/api/:UserID/meals", handlers.CreateMeal)
	protected.DELETE("/api/:UserID/meals/:MealID", handlers.DeleteMeal)
	protected.PUT("/api/:UserID/meals/:MealID", handlers.UpdateMeal)

	// Medication Log Handler
	protected.GET("/api/:UserID/medications", handlers.GetMedications)
	protected.POST("/api/:UserID/medications", handlers.CreateMedication)
	protected.DELETE("/api/:UserID/medications/:MedicationID", handlers.DeleteMedication)
	protected.PUT("/api/:UserID/medications/:MedicationID", handlers.UpdateMedication)

	// Blood Glucose Level Handler
	protected.GET("/api/:UserID/glucoses/info/graphic", handlers.GetBloodGlucoseGraphic)
	protected.GET("/api/:UserID/glucoses/info/average", handlers.GetBloodGlucoseAverage)
	protected.POST("/api/:UserID/glucoses", handlers.CreateBloodGlucoseLevel)

	// User Logs Handler
	protected.GET("/api/:UserID/logs", handlers.GetLogs)

	// Relation Handdler
	protected.GET("/api/:UserID/relations", handlers.GetRelations)
	protected.POST("/api/:UserID/relations", handlers.CreateRelation)
	protected.GET("/api/:UserID/relations/related", handlers.GetRelatedsInfo)

	// Scan Handler
	protected.POST("/api/:UserID/scan", handlers.ScanFood)
}

func applyCorsMiddleware(router *gin.Engine) {
	router.Use(middlewares.CorsMiddleware())
}
