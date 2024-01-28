package repository

import "task-tracker-backend/internal/model"

type UserRepository interface {
	SaveFromInput(input model.NewUser) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	Get(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Updates(values *model.User) error
	Remove(id uint) error
	Authenticate(creds model.Credentials) bool
}
