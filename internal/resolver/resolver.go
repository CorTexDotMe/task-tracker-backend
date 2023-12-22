package resolver

import "task-tracker-backend/internal/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepository repository.UserRepository
	taskRepository repository.TaskRepository
}
