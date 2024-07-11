package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/controller"
	"github.com/avran02/effective-mobile/internal/repository"
	"github.com/avran02/effective-mobile/internal/router"
	"github.com/avran02/effective-mobile/internal/service"
	"github.com/avran02/effective-mobile/logger"
)

type App struct {
	router.Router
	repository.Repository
	config *config.Config
}

func New() *App {
	conf := config.New()
	logger.Setup(conf.Server)
	slog.Info(fmt.Sprintf("Config: %+v", conf))
	repo := repository.New(conf.DB)
	service := service.New(repo, conf.ExternalAPI)
	controller := controller.New(service)
	router := router.New(controller, &conf.Server)

	return &App{
		Router: *router,
		config: conf,
	}
}

func (a *App) Run() error {
	serverEndpoint := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)
	slog.Info("Starting server at " + serverEndpoint)
	s := http.Server{
		Addr:    serverEndpoint,
		Handler: a.Router,
	}

	s.RegisterOnShutdown(func() {
		if err := a.Repository.Close(); err != nil {
			slog.Error("can't close db conn: " + err.Error())
		}
	})

	return s.ListenAndServe()
}
