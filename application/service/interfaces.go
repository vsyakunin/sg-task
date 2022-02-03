package service

import "github.com/vsyakunin/sg-task/domain/models"

type Repo interface {
	GetAllTasks() ([]models.Task, error)
	GetAllTasksForUser(int64) ([]models.Task, error)
	GetTaskByID(*int64) (models.Task, error)
	GetTaskHistory(*int64) ([]models.Message, error)
	GetMessageByMsgID(*int64) (models.Message, error)
	GetUserByLogin(*string) (models.User, error)
}

type Provider interface {
	DownloadFile(string) ([]byte, error)
}
