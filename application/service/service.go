package service

import (
	"sg-task/domain/models"
	myerrs "sg-task/domain/models/errors"

	log "github.com/sirupsen/logrus"
)

const (
	folderName  = "data"
	fileNameRaw = "%s/%s.json"

	validationErr = "validation error"
	internalErr   = "internal error"
)

type Service struct {
	Repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{
		Repo: repo,
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

	if err := validateTaskID(taskID); err != nil {
		log.Infof("%s: %s %v", funcName, validationErr, err)
		return history, err
	}

	history, err := svc.Repo.GetTaskHistory(taskID)
	if err != nil {
		log.Errorf("%s: error while getting tasks from db", funcName)
		return history, myerrs.NewServerError(internalErr, err)
	}

	return history, nil
}
