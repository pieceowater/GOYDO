package task

import (
	"GOYDO/src/modules/task/models"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := s.DB.Order("urgent desc, deadline desc").Find(&tasks).Error
	return tasks, err
}

func (s *Service) CreateTask(task models.Task) (models.Task, error) {
	err := s.DB.Create(&task).Error
	return task, err
}

func (s *Service) GetTask(id string) (models.Task, error) {
	var task models.Task
	err := s.DB.First(&task, id).Error
	return task, err
}

func (s *Service) UpdateTask(id string, task models.Task) (models.Task, error) {
	var existingTask models.Task
	if err := s.DB.First(&existingTask, id).Error; err != nil {
		return existingTask, err
	}
	task.ID = existingTask.ID
	err := s.DB.Save(&task).Error
	return task, err
}

func (s *Service) DeleteTask(id string) error {
	return s.DB.Delete(&models.Task{}, id).Error
}
