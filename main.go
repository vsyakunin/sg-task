package main

import (
	"net/http"
	"time"

	"sg-task/application/service"
	"sg-task/domain/repository"
	"sg-task/presentation/controller"
	"sg-task/presentation/router"

	log "github.com/sirupsen/logrus"
)

const httpTimeout = 30 * time.Second

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	cont := controller.NewController(svc)

	router := router.NewRouter(cont)
	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  httpTimeout,
		WriteTimeout: httpTimeout,
	}

	log.Info("Server is up and running")
	log.Fatal(server.ListenAndServe().Error())
}
