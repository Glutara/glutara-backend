package models

type Relation struct {
	UserID		int64	`json:"UserID"`
	RelationID	int64	`json:"RelationID"`
	Name		string	`json:"Name"`
	Phone		string	`json:"Phone"`
	Longitude	float64	`json:"Longitude"`
	Latitude	float64	`json:"Latitude"`
}