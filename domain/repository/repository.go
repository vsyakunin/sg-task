package repository

import (
	"strings"

	"github.com/vsyakunin/sg-task/domain/models"
)

type Repo struct{}

func NewRepository() *Repo {
	return &Repo{}
}

func (r *Repo) GetUserByLogin(userLogin *string) (foundUser models.User, err error) {
	for _, user := range models.Users {
		if strings.EqualFold(user.Login, *userLogin) {
			foundUser = user
		}
	}

	return
}

func (r *Repo) GetAllTasks() (tasks []models.Task, err error) {
	return models.AllTasks, nil
}

func (r *Repo) GetTaskByID(taskID *int64) (ret models.Task, err error) {
	for _ ,task := range models.AllTasks {
		if task.ID == *taskID {
			ret = task
		}
	}

	return
}

func (r *Repo) GetAllTasksForUser(userID int64) (tasks []models.Task, err error) {
	for _, task := range models.AllTasks {
		if task.UserID == userID {
			tasks = append(tasks, task)
		}
	}

	return
}

func (r *Repo) GetTaskHistory(taskID *int64) (history []models.Message, err error) {
	for _, msg := range models.Messages {
		if msg.TaskID == *taskID {
			history = append(history, msg)
		}
	}

	return
}

func (r *Repo) GetMessageByMsgID(msgID *int64) (message models.Message, err error) {
	allMessages := models.Messages

	for _, msg := range allMessages {
		if msg.ID == *msgID {
			message = msg
		}
	}

	return
}
