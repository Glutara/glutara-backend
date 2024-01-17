package models

import (
	"time"
)

type Sleep struct {
	UserID			int64		`json:"UserID"`
	SleepID			int64		`json:"SleepID"`
	StartTime		time.Time	`json:"StartTime"`
	EndTime			time.Time	`json:"EndTime"`
}