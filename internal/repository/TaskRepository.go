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

func (r *TaskRepository) Get(id uint) (*model.Task, error) {
	var task model.Task
	err := database.DB.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) GetByUserId(userId uint) ([]*model.Task, error) {
	var tasks []*model.Task
	err := database.DB.Find(&tasks, model.Task{UserId: userId}).Error
	return tasks, err
}

// FOR DEBUG
func (r *TaskRepository) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) Updates(values *model.Task) error {
	// TODO return task
	// task := &model.Task{}
	// database.DB.Model(task).Clauses(clause.Returning{}).Where("id = ?", "").Updates(values)
	return database.DB.Updates(values).Error
}

func (r *TaskRepository) Remove(id uint) error {
	return database.DB.Delete(&model.Task{}, id).Error
}

func (r *TaskRepository) LoadUser(task *model.Task) {
	database.DB.Preload("User").First(task)
}
