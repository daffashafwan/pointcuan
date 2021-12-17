package model

import "time"

type User struct {
	Id   int64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username" gorm: "unique"` 
	Password string `json:"password"`
	Email string `json:"email"`
	Address string `json:"address"`
	Status string `json:"status"`
	Token string `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}