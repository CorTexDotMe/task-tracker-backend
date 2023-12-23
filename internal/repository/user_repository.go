package repository

import (
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/utils"
)

type UserRepository struct{}

// Creates user from NewUser input
func (r *UserRepository) SaveFromInput(input model.NewUser) (*model.User, error) {
	user := &model.User{Name: input.Username, Password: input.Password}
	return r.Save(user)
}

// Creates user if no id provided. Updates user if id provided
func (r *UserRepository) Save(user *model.User) (*model.User, error) {
	err := database.DB.Save(user).Error
	return user, err
}

func (r *UserRepository) Get(id uint) (*model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	user := model.User{Name: username}
	err := database.DB.First(&user).Error
	return &user, err
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) Updates(values *model.User) error {
	// TODO return user
	// user := &model.User{}
	// database.DB.Model(user).Clauses(clause.Returning{}).Where("id = ?", "").Updates(values)
	return database.DB.Updates(values).Error
}

func (r *UserRepository) Remove(id uint) error {
	return database.DB.Delete(&model.User{}, id).Error
}

func (r *UserRepository) Authenticate(creds model.Credentials) bool {
	user, err := r.GetByUsername(creds.Username)
	utils.HandleError(err)

	//TODO normal encryption
	return user.Password == creds.Password
}
