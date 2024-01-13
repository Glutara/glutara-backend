package models

import (
	"time"
)

type Reminder struct {
	UserID			int64		`json:"UserID"`
	ReminderID		int64		`json:"ReminderID"`
	Name			string		`json:"Name"`
	Description		string		`json:"Description"`
	Time			time.Time	`json:"Time"`
}
