package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
}

type Tierlist struct {
	gorm.Model
	UserID uint
	User User `json:"creator"`
}