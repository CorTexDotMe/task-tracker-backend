package model

import "gorm.io/gorm"

// Default User entity
type User struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex" json:"name"`
	Password string `json:"password"`
}

// DTO object with data to create new User
type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
