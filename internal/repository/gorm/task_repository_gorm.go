package gorm

import (
	"sync"
	"task-tracker-backend/internal/database/postgres"
	"task-tracker-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type TaskRepositoryGorm struct {
	DB *gorm.DB
}

var taskRepository *TaskRepositoryGorm
var taskRepositoryOnce sync.Once

// Get singleton instance of Task repository
func GetTaskRepository() *TaskRepositoryGorm {
	taskRepositoryOnce.Do(func() {
		taskRepository = &TaskRepositoryGorm{postgres.GetService().GetDB()}
	})

	return taskRepository
}

// Convert object of type model.NewTask to model.Task
func (r *TaskRepositoryGorm) TaskFromNewTask(newTask model.NewTask, user *model.User) *model.Task {
	var parsedDueDate *time.Time
	if newTask.DueDate != nil {
		parsed, err := time.Parse(time.RFC3339, *newTask.DueDate)
		if err == nil {
			parsedDueDate = &parsed
		}
	}

	return &model.Task{
		Title:       newTask.Title,
		Description: newTask.Description,
		Status:      *newTask.Status,
		Done:        false,
		DueDate:     parsedDueDate,
		User:        user,
	}
}

// Save Task to database from NewTask data
func (r *TaskRepositoryGorm) SaveFromInput(input model.NewTask, user *model.User) (*model.Task, error) {
	task := r.TaskFromNewTask(input, user)
	return r.Save(task)
}

// Creates task if no id provided. Updates task if id provided
func (r *TaskRepositoryGorm) Save(task *model.Task) (*model.Task, error) {
	err := r.DB.Save(task).Error
	return task, err
}

func (r *TaskRepositoryGorm) Get(id uint) (*model.Task, error) {
	var task model.Task
	err := r.DB.First(&task, id).Error
	return &task, err
}

func (r *TaskRepositoryGorm) GetByUserId(userId uint) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.DB.Find(&tasks, model.Task{UserId: userId}).Error
	return tasks, err
}

// FOR DEBUG ONLY
func (r *TaskRepositoryGorm) GetAll() ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryGorm) Updates(values *model.Task) error {
	// TODO return task
	// task := &model.Task{}
	// r.DB.Model(task).Clauses(clause.Returning{}).Where("id = ?", "").Updates(values)
	return r.DB.Updates(values).Error
}

func (r *TaskRepositoryGorm) Remove(id uint) error {
	return r.DB.Delete(&model.Task{}, id).Error
}

// Eagerly load User field of Task object
func (r *TaskRepositoryGorm) LoadUser(task *model.Task) {
	r.DB.Preload("User").First(task)
}
