package service

import (
	"errors"
	"fmt"

	"github.com/vsyakunin/sg-task/domain/models"
	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"

	log "github.com/sirupsen/logrus"
)

const (
	AccessErr = "access error"

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

func (svc *Service) GetAllTasks(userLogin *string) (tasks []models.Task, err error) {
	const funcName = "service.GetAllTasks"

	user, err := svc.Repo.GetUserByLogin(userLogin)
	if err != nil {
		log.Errorf("%s: error while getting user for login %s from db", funcName, *userLogin)
		return tasks, myerrs.NewServerError(internalErr, err)
	}

	switch user.Role {
	case models.RoleUser:
		tasks, err = svc.Repo.GetAllTasksForUser(user.ID)
	case models.RoleOperator:
		tasks, err = svc.Repo.GetAllTasks()
	default:
		err = fmt.Errorf("unknown user role %s", user.Role)
	}

	if err != nil {
		log.Errorf("%s: error while getting tasks from db", funcName)
		return tasks, myerrs.NewServerError(internalErr, err)
	}

	return tasks, nil
}

func (svc *Service) GetTaskHistory(taskID *int64, userLogin *string) ([]models.Message, error) {
	const funcName = "service.GetTaskHistory"

	var history []models.Message

	if err := validateID(taskID, TaskIDParam); err != nil {
		log.Infof("%s: %s %v", funcName, validationErr, err)
		return history, myerrs.NewBusinessError(validationErr, err)
	}

	user, err := svc.Repo.GetUserByLogin(userLogin)
	if err != nil {
		log.Errorf("%s: error while getting user for login %s from db", funcName, *userLogin)
		return history, myerrs.NewServerError(internalErr, err)
	}

	task, err := svc.Repo.GetTaskByID(taskID)
	if err != nil {
		log.Errorf("%s: error while getting task for task ID = %d from db", funcName, *taskID)
		return history, myerrs.NewServerError(internalErr, err)
	}

	if user.Role != models.RoleOperator && user.ID != task.UserID {
		err = errors.New("you are not authorized to see this task's history")
		return history, myerrs.NewBusinessError(AccessErr, err)
	}

	history, err = svc.Repo.GetTaskHistory(taskID)
	if err != nil {
		log.Errorf("%s: error while getting messages for task ID = %d from db", funcName, *taskID)
		return history, myerrs.NewServerError(internalErr, err)
	}

	return history, nil
}

func (svc *Service) DownloadFileFromMessage(msgID *int64, userLogin *string) ([]byte, error) {
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

	task, err := svc.Repo.GetTaskByID(&msg.TaskID)
	if err != nil {
		log.Errorf("%s: error while getting task for task ID = %d from db", funcName, msg.TaskID)
		return nil, myerrs.NewServerError(internalErr, err)
	}

	user, err := svc.Repo.GetUserByLogin(userLogin)
	if err != nil {
		log.Errorf("%s: error while getting user for login %s from db", funcName, *userLogin)
		return nil, myerrs.NewServerError(internalErr, err)
	}

	if user.Role != models.RoleOperator && user.ID != task.UserID {
		err = errors.New("you are not allowed to download this file since it doesn't relate to any of your tasks")
		return nil, myerrs.NewBusinessError(AccessErr, err)
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
