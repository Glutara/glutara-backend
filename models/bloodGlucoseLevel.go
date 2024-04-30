package models

import (
	"time"
)

type BloodGlucoseLevel struct {
	UserID					int64		`json:"UserID"`
	BloodGlucoseLevelID		int64		`json:"BloodGlucoseLevelID"`
	Input					float32		`json:"Input"`
	Prediction				float32		`json:"Prediction"`
	Time					time.Time	`json:"Time"`
}

type InferenceReqToken struct {
	Instances 	[]float32 `json:"instances"`
}

type PredictionResToken struct {
	Predictions []float32 `json:"predictions"`
}

type GraphicToken struct {
	Date 	time.Time	`json:"Date"`
}

type GraphicData struct {
	Prediction				float32		`json:"Prediction"`
	Time					time.Time	`json:"Time"`
}