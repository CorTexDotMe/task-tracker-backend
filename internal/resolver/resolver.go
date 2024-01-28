package resolver

import (
	"task-tracker-backend/internal/repository"
	"task-tracker-backend/internal/repository/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepository repository.UserRepository
	taskRepository repository.TaskRepository
}

func NewResolver() *Resolver {
	return &Resolver{userRepository: gorm.GetUserRepository(), taskRepository: gorm.GetTaskRepository()}
}
