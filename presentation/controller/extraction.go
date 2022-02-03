package controller

import (
	"fmt"
	"net/http"
	"strconv"

	myerrs "github.com/vsyakunin/sg-task/domain/models/errors"

	"github.com/gorilla/mux"
)

const (
	paramNotFoundErr = "url parameter %s not found"
	paramParseErr    = "parameter parsing error"
)

func extractID(r *http.Request, idType string) (*int64, error) {
	idStr, ok := mux.Vars(r)[idType]
	if !ok {
		err := fmt.Errorf(paramNotFoundErr, idType)
		return nil, myerrs.NewBusinessError(paramParseErr, err)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, myerrs.NewBusinessError(paramParseErr, err)
	}

	return &id, nil
}

func extractUserLogin(r *http.Request) *string {
	login, _, _ := r.BasicAuth()
	return &login
}
