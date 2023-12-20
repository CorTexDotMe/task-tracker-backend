package repositories

import (
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
)

type UserRepository struct{}

func (r *UserRepository) SaveFromInput(input model.NewUser) (*model.User, error) {
	//TODO return error when unique constrains violation
	user := model.User{Name: input.Username, Password: input.Password}
	r.Save(user)
	return &user, nil
}

// Returning the inserted data's primary key in id field of user given as parameter
func (r *UserRepository) Save(user model.User) {
	database.DB.Create(user)
}
