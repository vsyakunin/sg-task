package main

import (
	"net/http"
	"time"

	"github.com/vsyakunin/sg-task/application/service"
	"github.com/vsyakunin/sg-task/bootstrap"
	"github.com/vsyakunin/sg-task/domain/provider"
	"github.com/vsyakunin/sg-task/domain/repository"
	"github.com/vsyakunin/sg-task/presentation/controller"
	mw "github.com/vsyakunin/sg-task/presentation/middlewares"
	"github.com/vsyakunin/sg-task/presentation/router"

	log "github.com/sirupsen/logrus"
)

const httpTimeout = 30 * time.Second

func main() {
	cfg := bootstrap.InitConfig("config.toml")

	mw.InitBasicAuthMiddleware(cfg)

	repo := repository.NewRepository()
	prov := provider.NewProvider()
	svc := service.NewService(repo, prov)
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
