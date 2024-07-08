package service

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/models"
	"github.com/avran02/effective-mobile/internal/repository"
)

type Service interface {
	GetUsers(limit, offset int, filterField []string) ([]models.User, error)
	CreateUser(passportNumber string) (int, error)
	UpdateUserData(user models.User) error
	DeleteUser(userID int) error

	CreateTask(name, description string) (int, error)
	GetUserTasks(userID int, startDate, endDate string) []models.Task
	StartUserTask(userID, taskID int) error
	StopUserTask(userID, taskID int) error
}

type service struct {
	repo            repository.Repository
	externalAPIConf config.ExternalAPI
}

func (s *service) GetUsers(limit, offset int, filterField []string) ([]models.User, error) {
	slog.Info("GetUsers service")
	return s.repo.GetUsers(limit, offset, filterField[0])
}

func (s *service) CreateUser(passportNumber string) (int, error) {
	slog.Info("CreateUser service")
	passportNumberParts := strings.Split(passportNumber, " ")
	user, err := enrichUserData(s.externalAPIConf.EnrichUserDataEndpoint, passportNumberParts[0], passportNumberParts[1])
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}

	user.PassportSerie = passportNumberParts[0]
	user.PassportNumber = passportNumberParts[1]

	slog.Info(fmt.Sprintf("%+v", user))

	return s.repo.CreateUser(user)
}

func (s *service) UpdateUserData(user models.User) error {
	slog.Info("UpdateUserData service")
	return s.repo.UpdateUserData(user)
}

func (s *service) DeleteUser(userID int) error {
	slog.Info("DeleteUser service")
	return s.repo.DeleteUser(models.User{})
}

func (s *service) GetUserTasks(userID int, startDate, endDate string) []models.Task {
	slog.Info("GetUserTasks service")
	return s.repo.GetUserTasks(userID)
}

func (s *service) CreateTask(name, description string) (int, error) {
	slog.Info("CreateTask service")
	return 0, s.repo.CreateTask(models.Task{})
}

func (s *service) StartUserTask(userID, taskID int) error {
	slog.Info("StartUserTask service")
	return s.repo.StartUserTask(models.Task{})
}

func (s *service) StopUserTask(userID, taskID int) error {
	slog.Info("StopUserTask service")
	return s.repo.StopUserTask(models.Task{})
}

func New(repo repository.Repository, conf config.ExternalAPI) Service {
	return &service{
		repo: repo,

		externalAPIConf: conf,
	}
}
