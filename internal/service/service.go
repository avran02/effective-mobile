package service

import (
	"log/slog"

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
	StartUserTask(userID, taskId int) error
	StopUserTask(userID, taskId int) error
}

type service struct {
	repo repository.Repository
}

func (s *service) GetUsers(limit, offset int, filterField []string) ([]models.User, error) {
	slog.Info("GetUsers service")
	return s.repo.GetUsers(limit, offset, filterField[0])
}

func (s *service) CreateUser(passportNumber string) (int, error) {
	slog.Info("CreateUser service")
	user := models.User{
		PassportNumber: passportNumber,
	}
	return 0, s.repo.CreateUser(user)
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

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
