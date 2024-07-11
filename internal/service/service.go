package service

import (
	"fmt"
	"log/slog"
	"sort"
	"strings"
	"time"

	"github.com/avran02/effective-mobile/config"
	"github.com/avran02/effective-mobile/internal/models"
	"github.com/avran02/effective-mobile/internal/repository"
)

type Service interface {
	GetUsers(limit, offset int, filters map[string]string) ([]models.User, error)
	CreateUser(passportNumber string) (int, error)
	UpdateUserData(user models.User) error
	DeleteUser(userID int) error

	CreateTask(task models.Task) (int, error)
	GetUserTasks(userID int, startDate, endDate string) ([]models.TaskTimeSpent, error)
	StartUserTask(taskID int) error
	StopUserTask(taskID int) error
}

type service struct {
	repo            repository.Repository
	externalAPIConf config.ExternalAPI
}

func (s *service) GetUsers(page, pageSize int, filters map[string]string) ([]models.User, error) {
	slog.Info("GetUsers service")
	return s.repo.GetUsers(page, pageSize, filters)
}

func (s *service) CreateUser(passportNumber string) (int, error) {
	slog.Info("CreateUser service")
	passportNumberParts := strings.Split(passportNumber, " ")
	user, err := enrichUserData(s.externalAPIConf.EnrichUserDataEndpoint, passportNumberParts[0], passportNumberParts[1])
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}

	user.PassportNumber = passportNumber

	slog.Info(fmt.Sprintf("%+v", user))

	return s.repo.CreateUser(user)
}

func (s *service) UpdateUserData(user models.User) error {
	slog.Info("UpdateUserData service")
	return s.repo.UpdateUserData(user)
}

func (s *service) DeleteUser(userID int) error {
	slog.Info("DeleteUser service")
	return s.repo.DeleteUser(userID)
}

func (s *service) GetUserTasks(userID int, startDate, endDate string) ([]models.TaskTimeSpent, error) {
	slog.Info("GetUserTasks service")

	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil
	}

	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil
	}

	tasksTimeSpent, err := s.repo.GetUserTasks(userID, startTime, endTime)
	if err != nil {
		slog.Error(err.Error())
		return nil, nil
	}

	sort.Slice(tasksTimeSpent, func(i, j int) bool {
		return tasksTimeSpent[i].TimeSpentSeconds > tasksTimeSpent[j].TimeSpentSeconds
	})

	return tasksTimeSpent, nil
}

func (s *service) CreateTask(task models.Task) (int, error) {
	slog.Info("CreateTask service")
	return s.repo.CreateTask(task)
}

func (s *service) StartUserTask(taskID int) error {
	slog.Info("StartUserTask service")
	return s.repo.StartUserTask(taskID)
}

func (s *service) StopUserTask(taskID int) error {
	slog.Info("StopUserTask service")
	return s.repo.StopUserTask(taskID)
}

func New(repo repository.Repository, conf config.ExternalAPI) Service {
	return &service{
		repo: repo,

		externalAPIConf: conf,
	}
}
