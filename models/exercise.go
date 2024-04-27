package models

import (
	"time"
)

type ExerciseLog struct {
	Type		string		`json:"Type"`
	Data 		Exercise 	`json:"Data"`
}

type Exercise struct {
	UserID			int64		`json:"UserID"`
	ExerciseID		int64		`json:"ExerciseID"`
	Name			string		`json:"Name"`
	Intensity		int64		`json:"Intensity"`
	Date			time.Time	`json:"Date"`
	StartTime		time.Time	`json:"StartTime"`
	EndTime			time.Time	`json:"EndTime"`
}