package repository

import (
	"task-tracker-backend/internal/database"
	"task-tracker-backend/internal/model"
)

type TaskRepository struct{}

func (r *TaskRepository) TaskFromNewTask(newTask model.NewTask, user *model.User) *model.Task {
	return &model.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      *newTask.Status,
		Done:        false,
		DueDate:     nil,
		User:        user,
	}
}

// Creates task from NewTask input
func (r *TaskRepository) SaveFromInput(input model.NewTask, user *model.User) (*model.Task, error) {
	task := r.TaskFromNewTask(input, user)
	return r.Save(task)
}

// Creates task if no id provided. Updates task if id provided
func (r *TaskRepository) Save(task *model.Task) (*model.Task, error) {
	err := database.DB.Save(task).Error
	return task, err
}

func (r *TaskRepository) LoadUser(task *model.Task) {
	database.DB.Preload("User").First(task)
}

func (r *TaskRepository) Remove(id uint) error {
	return database.DB.Delete(&model.Task{}, id).Error
}

func (r *TaskRepository) Get(id uint) error {
	return database.DB.First(&model.Task{}, id).Error
}
