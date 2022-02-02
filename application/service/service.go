package service

import (
	"github.com/vsyakunin/sg-task/domain/models"
	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"

	log "github.com/sirupsen/logrus"
)

const (
	validationErr = "validation error"
	internalErr   = "internal error"
)

type Service struct {
	Repo     Repo
	Provider Provider
}

func NewService(repo Repo, prov Provider) *Service {
	return &Service{
		Repo:     repo,
		Provider: prov,
	}
}

func (svc *Service) GetAllTasks() ([]models.Task, error) {
	const funcName = "service.GetAllTasks"

	var tasks []models.Task

	tasks, err := svc.Repo.GetAllTasks()
	if err != nil {
		log.Errorf("%s: error while getting tasks from db", funcName)
		return tasks, myerrs.NewServerError(internalErr, err)
	}

	return tasks, nil
}

func (svc *Service) GetTaskHistory(taskID *int64) ([]models.Message, error) {
	const funcName = "service.GetTaskHistory"

	var history []models.Message

	if err := validateID(taskID, TaskIDParam); err != nil {
		log.Infof("%s: %s %v", funcName, validationErr, err)
		return history, err
	}

	history, err := svc.Repo.GetTaskHistory(taskID)
	if err != nil {
		log.Errorf("%s: error while getting messages for task ID = %d from db", funcName, *taskID)
		return history, myerrs.NewServerError(internalErr, err)
	}

	return history, nil
}

func (svc *Service) DownloadFileFromMessage(msgID *int64) ([]byte, error) {
	const funcName = "service.DownloadFileFromMessage"

	if err := validateID(msgID, MsgIDParam); err != nil {
		log.Infof("%s: %s %v", funcName, validationErr, err)
		return nil, err
	}

	msg, err := svc.Repo.GetMessageByMsgID(msgID)
	if err != nil {
		log.Errorf("%s: error while getting message with ID = %d from db", funcName, *msgID)
		return nil, myerrs.NewServerError(internalErr, err)
	}

	if msg.FileKey == "" {
		return nil, nil
	}

	file, err := svc.Provider.DownloadFile(msg.FileKey)
	if err != nil {
		log.Errorf("%s: error while getting file for message with ID = %d and file key = %s from filestorage", funcName, *msgID, msg.FileKey)
		return nil, myerrs.NewServerError(internalErr, err)
	}

	return file, nil
}
