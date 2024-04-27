package models

import (
	"time"
)

type Log interface {
	GetTime() time.Time
}

type DateLogs struct {
	Date	time.Time	`json:"Date"`
	Logs 	[]Log 		`json:"Logs"`
}

func (sleepLog SleepLog) GetTime() time.Time {
	return sleepLog.Data.StartTime
}

func (exerciseLog ExerciseLog) GetTime() time.Time {
	return exerciseLog.Data.StartTime
}

func (mealLog MealLog) GetTime() time.Time {
	return mealLog.Data.StartTime
}

func (medicationLog MedicationLog) GetTime() time.Time {
	return medicationLog.Data.Time
}