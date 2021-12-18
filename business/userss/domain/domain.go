package domain

import "time"

type Domain struct {
	Id   int64 
	Name string 
	Username string 
	Password string 
	Email string 
	Address string 
	Status string 
	Token string 
	CreatedAt time.Time 
	UpdatedAt time.Time 
}