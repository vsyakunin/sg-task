package controller

import (
	"fmt"
	"net/http"
	"strconv"

	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"

	"github.com/gorilla/mux"
)

const (
	taskIDParam = "taskID"

	paramNotFoundErr = "url parameter %s not found"
	paramParseErr    = "parameter parsing error"
)

func extractID(r *http.Request) (*int64, error) {
	taskIDStr, ok := mux.Vars(r)[taskIDParam]
	if !ok {
		err := fmt.Errorf(paramNotFoundErr, taskIDParam)
		return nil, myerrs.NewBusinessError(paramParseErr, err)
	}

	taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
	if err != nil {
		return nil, myerrs.NewBusinessError(paramParseErr, err)
	}

	return &taskID, nil
}
