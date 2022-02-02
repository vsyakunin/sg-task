package service

import (
	"errors"
	"fmt"

	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"
)

const (
	TaskIDParam = "taskID"
	MsgIDParam  = "messageID"

	invalidParameterErr = "invalid parameter"
)

func validateID(id *int64, param string) error {
	if *id < 1 {
		err := errors.New(fmt.Sprintf("%s parameter must be positive", param))
		return myerrs.NewBusinessError(invalidParameterErr, err)
	}

	return nil
}
