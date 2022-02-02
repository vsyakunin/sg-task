package router

import (
	"net/http"

	"sg-task/presentation/controller"

	"github.com/gorilla/mux"
)

const (
	apiVer = "/api/v1"
)

func NewRouter(cont *controller.Controller) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	{
		const routeName = "getAllTasks"
		router.Methods(http.MethodGet).
			Name(routeName).
			PathPrefix(apiVer).
			Path("/tasks").
			Handler(http.HandlerFunc(cont.GetAllTasks))
	}

	{
		const routeName = "getTaskHistory"
		router.Methods(http.MethodGet).
			Name(routeName).
			PathPrefix(apiVer).
			Path("/tasks/{taskID}/history").
			Handler(http.HandlerFunc(cont.GetTaskHistory))
	}

	return router
}
