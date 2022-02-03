package controller

import "github.com/vsyakunin/sg-task/domain/models"

type Service interface {
	GetAllTasks(*string) ([]models.Task, error)
	GetTaskHistory(*int64, *string) ([]models.Message, error)
	DownloadFileFromMessage(*int64, *string) ([]byte, error)
}
