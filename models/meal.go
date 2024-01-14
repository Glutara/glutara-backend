package models

import (
	"time"
)

type Meal struct {
	UserID			int64		`json:"UserID"`
	MealID			int64		`json:"MealID"`
	Name			string		`json:"Name"`
	Calories		int64		`json:"Calories"`
	Type			int64		`json:"Type"`
	Date			time.Time	`json:"Date"`
	StartTime		time.Time	`json:"StartTime"`
	EndTime			time.Time	`json:"EndTime"`
}