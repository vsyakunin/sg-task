package service

import (
	"errors"

	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"
)

const (
	invalidParameterErr = "invalid parameter"
)

func validateTaskID(taskID *int64) error {
	if *taskID < 1 {
		err := errors.New("task ID parameter must be positive")
		return myerrs.NewBusinessError(invalidParameterErr, err)
	}

	return nil
}
