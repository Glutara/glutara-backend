package models

type Relation struct {
	UserID				int64	`json:"UserID"`
	Name				string	`json:"Name"`
	Phone 				string	`json:"Phone"`
	RelationID			int64	`json:"RelationID"`
	RelationName		string	`json:"RelationName"`
	RelationPhone		string	`json:"RelationPhone"`
	Longitude			float64	`json:"Longitude"`
	Latitude			float64	`json:"Latitude"`
}

type Related struct {
	UserID				int64	`json:"UserID"`
	Name				string	`json:"Name"`
	Phone 				string	`json:"Phone"`
}

type RelatedInfo struct {
	UserID				int64	`json:"UserID"`
	Name				string	`json:"Name"`
	Phone 				string	`json:"Phone"`
	LatestBloodGlucose 	float32	`json:"LatestBloodGlucose"`
}