package repositories

import (
	"strconv"
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/utils"
)

type UserRepository struct{}

func (r *UserRepository) Save(user model.User) model.User {
	statement, err := database.DATABASE_CONNECTION.Prepare("INSERT INTO users(username,password) VALUES(?,?)")
	utils.HandleError(err)

	res, err := statement.Exec(user.Name, user.Password)
	utils.HandleError(err)

	id, err := res.LastInsertId()
	utils.HandleError(err)

	user.ID = strconv.FormatInt(id, 10)
	return user
}
