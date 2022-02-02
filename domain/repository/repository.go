package repository

import "github.com/vsyakunin/sg-task/domain/models"

type Repo struct{}

func NewRepository() *Repo {
	return &Repo{}
}

func (r *Repo) GetAllTasks() ([]models.Task, error) {
	return models.AllTasks, nil
}

func (r *Repo) GetTaskHistory(taskID *int64) ([]models.Message, error) {
	var history []models.Message
	allMessages := models.Messages

	for _, msg := range allMessages {
		if msg.TaskID == *taskID {
			history = append(history, msg)
		}
	}

	return history, nil
}

func (r *Repo) GetMessageByMsgID(msgID *int64) (models.Message, error) {
	var message models.Message

	allMessages := models.Messages

	for _, msg := range allMessages {
		if msg.ID == *msgID {
			return msg, nil
		}
	}

	return message, nil
}
