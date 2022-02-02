package controller

import "github.com/vsyakunin/sg-task/domain/models"

type Service interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskHistory(*int64) ([]models.Message, error)
	DownloadFileFromMessage(*int64) ([]byte, error)
}
