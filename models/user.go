package models

type User struct {
	ID			int64	`json:"id"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Name		string	`json:"name"`
	Role		int64	`json:"role"`
	Phone		string	`json:"phone"`
}

type LoginToken struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}