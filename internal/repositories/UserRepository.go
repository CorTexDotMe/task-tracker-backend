package repositories

import (
	"task-tracker-backend/internal/database"
	model "task-tracker-backend/internal/models"
)

type UserRepository struct{}

func (r *UserRepository) SaveFromInput(input model.NewUser) (*model.User, error) {
	user := &model.User{Name: input.Username, Password: input.Password}
	return r.Save(user)
}

// Creates user If no id provided. Updates user if id provided
func (r *UserRepository) Save(user *model.User) (*model.User, error) {
	err := database.DB.Save(user).Error
	return user, err
}
