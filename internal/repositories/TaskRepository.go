package repositories

import (
	"strconv"
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/utils"
)

type TaskRepository struct{}

func (r *TaskRepository) NewTaskToTask(newTask model.NewTask) model.Task {
	return model.Task{
		ID:          "",
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      *newTask.Status,
		Done:        false,
		DateCreated: "",
		DueDate:     nil,
	}
}

func (r *TaskRepository) Save(task model.Task) *model.Task {
	statement, err := database.DATABASE_CONNECTION.Prepare(
		"INSERT INTO tasks(title,description,status,done,dateCreated,dueDate,userID) VALUES(?,?,?,?,?,?,?)")
	utils.HandleError(err)

	//TODO add userID
	res, err := statement.Exec(task.Title, task.Description, task.Status, task.Done, task.DateCreated, task.DueDate)
	utils.HandleError(err)

	id, err := res.LastInsertId()
	utils.HandleError(err)

	task.ID = strconv.FormatInt(id, 10)
	return &task
}
