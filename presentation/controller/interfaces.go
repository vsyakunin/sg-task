package controller

import "sg-task/domain/models"

type Service interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskHistory(*int64) ([]models.Message, error)
}
