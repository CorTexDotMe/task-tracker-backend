package model

import (
	"time"

	"gorm.io/gorm"
)

// Default Task entity
type Task struct {
	gorm.Model
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status"`
	Done        bool       `json:"done"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
	UserId      uint       `json:"userId"`
	User        *User      `gorm:"reference:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

// DTO object with data to create new Task
type NewTask struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"status,omitempty"`
	DueDate     *string `json:"dueDate,omitempty"`
}
