package models

import (
	"time"
)

type Reminder struct {
	UserID			int64		`json:"UserID"`
	ReminderID		int64		`json:"ReminderID"`
	Name			string		`json:"name"`
	Description		string		`json:"description"`
	Time			time.Time	`json:"time"`
}
