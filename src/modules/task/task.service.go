package task

import (
	"GOYDO/src/modules/task/models"
	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := s.DB.Order("urgent desc, deadline desc").Find(&tasks).Error
	return tasks, err
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	err := s.DB.Create(&task).Error
	return task, err
}

func (s *TaskService) GetTask(id string) (models.Task, error) {
	var task models.Task
	err := s.DB.First(&task, id).Error
	return task, err
}

func (s *TaskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	var existingTask models.Task
	if err := s.DB.First(&existingTask, id).Error; err != nil {
		return existingTask, err
	}
	task.ID = existingTask.ID
	err := s.DB.Save(&task).Error
	return task, err
}

func (s *TaskService) DeleteTask(id string) error {
	return s.DB.Delete(&models.Task{}, id).Error
}
