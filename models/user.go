package models

type User struct {
	ID       			int64  	`json:"ID"`
	Email    			string  `json:"Email"`
	Password 			string  `json:"Password"`
	Name     			string  `json:"Name"`
	Role     			int64  	`json:"Role"`
	Phone    			string 	`json:"Phone"`
	LatestBloodGlucose 	float32	`json:"LatestBloodGlucose"`
}

type LatestBloodGlucoseInfo struct {
	LatestBloodGlucose 	float32	`json:"LatestBloodGlucose"`	
}

type LoginRequest struct {	
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
