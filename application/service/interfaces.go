package service

import "github.com/vsyakunin/sg-task/domain/models"

type Repo interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskHistory(*int64) ([]models.Message, error)
	GetMessageByMsgID(msgID *int64) (models.Message, error)
}

type Provider interface {
	DownloadFile(fileKey string) ([]byte, error)
}
