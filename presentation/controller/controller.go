package controller

import (
	"net/http"

	"github.com/vsyakunin/sg-task/application/service"

	log "github.com/sirupsen/logrus"
)

const (
	writerErr = "writer error"
)

type Controller struct {
	Svc Service
}

func NewController(svc Service) *Controller {
	return &Controller{Svc: svc}
}

func (c *Controller) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.GetAllTasks"

	hallLayout, err := c.Svc.GetAllTasks()
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, hallLayout) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) GetTaskHistory(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.GetTaskHistory"

	taskID, err := extractID(r, service.TaskIDParam)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	taskHistory, err := c.Svc.GetTaskHistory(taskID)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeJSONResponse(w, r.URL.Path, taskHistory) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) DownloadFileFromMessage(w http.ResponseWriter, r *http.Request) {
	const funcName = "controller.DownloadFileFromMessage"

	msgID, err := extractID(r, service.MsgIDParam)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	file, err := c.Svc.DownloadFileFromMessage(msgID)
	if err != nil {
		writeErrorResponse(w, err, r.URL.Path)
		return
	}

	if !writeFileResponse(w, r.URL.Path, file) {
		log.Errorf("%s: %s", funcName, writerErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
