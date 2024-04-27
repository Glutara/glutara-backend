package models

import (
	"time"
)

type MedicationLog struct {
	Type		string		`json:"Type"`
	Data 		Medication 	`json:"Data"`
}

type Medication struct {
	UserID			int64		`json:"UserID"`
	MedicationID	int64		`json:"MedicationID"`
	Type			int64		`json:"Type"`
	Category		string		`json:"Category"`
	Dose			int64		`json:"Dose"`
	Date			time.Time	`json:"Date"`
	Time			time.Time	`json:"Time"`
}