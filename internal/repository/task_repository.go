package repository

import "task-tracker-backend/internal/model"

type TaskRepository interface {
	TaskFromNewTask(newTask model.NewTask, user *model.User) *model.Task
	SaveFromInput(input model.NewTask, user *model.User) (*model.Task, error)
	Save(task *model.Task) (*model.Task, error)
	Get(id uint) (*model.Task, error)
	GetByUserId(userId uint) ([]*model.Task, error)
	GetAll() ([]*model.Task, error)
	Updates(values *model.Task) error
	Remove(id uint) error
	LoadUser(task *model.Task)
}
