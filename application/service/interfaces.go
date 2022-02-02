package service

import "sg-task/domain/models"

type Repo interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskHistory(*int64) ([]models.Message, error)
}
