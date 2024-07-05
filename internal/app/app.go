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

	"github.com/spf13/viper"
)

type App struct {
	router.Router
	repository.Repository
}

func New() *App {
	conf := config.New()
	repo := repository.New(conf.DB)
	service := service.New(repo)
	controller := controller.New(service)
	router := router.New(controller)

	return &App{
		Router: *router,
	}
}

func (a *App) Run() error {
	serverEndpoint := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
	slog.Info("Starting server at " + serverEndpoint)
	s := http.Server{ //nolint:gosec
		Addr:    serverEndpoint,
		Handler: a.Router,
	}

	s.RegisterOnShutdown(func() {
		if err := a.Repository.Close(); err != nil {
			slog.Error(err.Error())
		}
	})

	return s.ListenAndServe()
}
