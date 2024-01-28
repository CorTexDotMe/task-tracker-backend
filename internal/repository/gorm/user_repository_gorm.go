package gorm

import (
	"sync"
	"task-tracker-backend/internal/database/postgres"
	"task-tracker-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	DB *gorm.DB
}

var userRepository *UserRepositoryGorm
var userRepositoryOnce sync.Once

// Get singleton instance of User repository
func GetUserRepository() *UserRepositoryGorm {
	userRepositoryOnce.Do(func() {
		userRepository = &UserRepositoryGorm{postgres.GetService().GetDB()}
	})

	return userRepository
}

// Save User to database from data of NewUser
func (r *UserRepositoryGorm) SaveFromInput(input model.NewUser) (*model.User, error) {
	user := &model.User{Name: input.Username, Password: input.Password}
	return r.Save(user)
}

// Creates user if no id provided. Updates user if id provided
func (r *UserRepositoryGorm) Save(user *model.User) (*model.User, error) {
	err := r.DB.Save(user).Error
	return user, err
}

func (r *UserRepositoryGorm) Get(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryGorm) GetByUsername(username string) (*model.User, error) {
	user := model.User{Name: username}
	err := r.DB.First(&user).Error
	return &user, err
}

func (r *UserRepositoryGorm) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepositoryGorm) Updates(values *model.User) error {
	// TODO return user
	// user := &model.User{}
	// r.DB.Model(user).Clauses(clause.Returning{}).Where("id = ?", "").Updates(values)
	return r.DB.Updates(values).Error
}

func (r *UserRepositoryGorm) Remove(id uint) error {
	return r.DB.Delete(&model.User{}, id).Error
}

// Check if database has User with provided credentials
func (r *UserRepositoryGorm) Authenticate(creds model.Credentials) bool {
	user, err := r.GetByUsername(creds.Username)

	//TODO normal encryption
	return err == nil && user.Password == creds.Password
}
