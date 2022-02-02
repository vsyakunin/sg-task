package service

import (
	"github.com/vsyakunin/sg-task/domain/models"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (svc *Service) GetAllTasks() ([]models.Task, error) {
	const funcName = "service.GetAllTasks"

	var tasks []models.Task

	tasks = models.AllTasks

	return tasks, nil
}
