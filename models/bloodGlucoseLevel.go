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
	Predictions [][]float32 `json:"predictions"`
}

type GraphicData struct {
	Prediction				float32		`json:"Prediction"`
	Time					time.Time	`json:"Time"`
}

type AverageGlucoseLevel struct {
	Today	float32		`json:"Today"`
	Week	float32		`json:"Week"`
	Month	float32		`json:"Month"`
}