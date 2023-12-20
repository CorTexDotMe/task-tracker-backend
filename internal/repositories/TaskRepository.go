package repositories

import (
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
)

type TaskRepository struct{}

func (r *TaskRepository) TaskFromNewTask(newTask model.NewTask, user *model.User) model.Task {
	return model.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      *newTask.Status,
		Done:        false,
		DueDate:     nil,
		User:        user,
	}
}

func (r *TaskRepository) SaveFromInput(input model.NewTask, user *model.User) (*model.Task, error) {
	//TODO return error when unique constrains violation
	task := r.TaskFromNewTask(input, user)
	r.Save(task)
	return &task, nil
}

// Returning the inserted data's primary key in id field of task given as parameter
func (r *TaskRepository) Save(task model.Task) {
	database.DB.Create(task)
}

func (r *TaskRepository) LoadUser(task *model.Task) {
	database.DB.Preload("User").First(task)
}
